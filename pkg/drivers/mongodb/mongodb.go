package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoConnection struct {
	client *mongo.Client
	db     *mongo.Database
}

type MongoDB struct {
	activeConnection mongoConnection
	connectionUri    string
	database         string
}

func (db *MongoDB) Auth() error {
	opts := options.Client().
		ApplyURI(db.connectionUri).
		SetConnectTimeout(1 * time.Second) // Connection should have a short timeout to fail fast

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	// Store the Mongo connection
	db.activeConnection = mongoConnection{
		client: client,
		db:     client.Database(db.database),
	}

	return nil
}

func (db *MongoDB) Close() error {
	return db.activeConnection.client.Disconnect(context.TODO())
}

func (db *MongoDB) InsertBulk(collection string, raw []map[string]interface{}) (int, error) {
	var err error
	data := make([]interface{}, 0)

	// Conver the raw data into MongoDB format
	for _, row := range raw {
		item := bson.D{}

		for k, v := range row {
			item = append(item, bson.E{
				Key:   k,
				Value: v,
			})
		}

		data = append(data, item)
	}

	insertCount := len(data)

	if insertCount > 0 {
		_, err = db.activeConnection.db.
			Collection(collection).
			InsertMany(context.TODO(), data)
	}

	return insertCount, err
}

func (db *MongoDB) Truncate(collection string) error {
	return db.activeConnection.db.
		Collection(collection).
		Drop(context.TODO())
}

func New(connectionUri string, database string) *MongoDB {
	return &MongoDB{
		connectionUri: connectionUri,
		database:      database,
	}
}
