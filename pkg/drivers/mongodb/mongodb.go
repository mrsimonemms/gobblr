package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client        *mongo.Client
	connectionUri string
}

func (db *MongoDB) Auth() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(db.connectionUri).
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	db.client = client

	return nil
}

func (db *MongoDB) Close() error {
	return db.client.Disconnect(context.TODO())
}

func (db *MongoDB) InsertBulk(table string, data []map[string]interface{}) (int, error) {
	return len(data), nil
}

func (db *MongoDB) Truncate(table string) error {
	return nil
}

func New(connectionUri string) *MongoDB {
	return &MongoDB{
		connectionUri: connectionUri,
	}
}
