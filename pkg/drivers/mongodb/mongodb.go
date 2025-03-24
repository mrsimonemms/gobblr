/*
 * Copyright 2022 Simon Emms <simon@simonemms.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type mongoConnection struct {
	client *mongo.Client
	db     *mongo.Database
}

type MongoDB struct {
	activeConnection mongoConnection
	connectionURI    string
	database         string
}

func (db *MongoDB) Auth() error {
	opts := options.Client().
		ApplyURI(db.connectionURI).
		SetTimeout(time.Second).
		SetConnectTimeout(time.Second) // Connection should have a short timeout to fail fast

	client, err := mongo.Connect(opts)
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

func (db *MongoDB) DriverName() string {
	return "mongodb"
}

func (db *MongoDB) InsertBulk(collection string, raw []map[string]any) (int, error) {
	var err error
	data := make([]any, 0)

	// Conver the raw data into MongoDB format
	for _, row := range raw {
		item := bson.D{}

		for k, v := range row {
			data, err := parseData(k, v)
			if err != nil {
				return 0, err
			}

			item = append(item, data)
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

func New(connectionURI, database string) *MongoDB {
	return &MongoDB{
		connectionURI: connectionURI,
		database:      database,
	}
}

func parseData(k string, v any) (bson.E, error) {
	if k == "_id" {
		id, err := bson.ObjectIDFromHex(v.(string))
		if err != nil {
			return bson.E{}, fmt.Errorf("error converting string to _id: %w", err)
		}
		v = id
	}

	return bson.E{
		Key:   k,
		Value: v,
	}, nil
}
