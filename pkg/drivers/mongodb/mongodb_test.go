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
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestParseDate(t *testing.T) {
	tests := []struct {
		Name    string
		Key     string
		Value   any
		Output  bson.E
		ErrorIs error
	}{
		{
			Name:  "valid id string",
			Key:   "_id",
			Value: "507f1f77bcf86cd799439011",
			Output: func() bson.E {
				val, err := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
				if err != nil {
					panic(err)
				}

				return bson.E{
					Key:   "_id",
					Value: val,
				}
			}(),
		},
		{
			Name:    "invalid id string",
			Key:     "_id",
			Value:   "invalid",
			ErrorIs: bson.ErrInvalidHex,
		},
		{
			Name:  "something else",
			Key:   "id",
			Value: "hello",
			Output: bson.E{
				Key:   "id",
				Value: "hello",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			val, err := parseData(test.Key, test.Value)
			if test.ErrorIs != nil {
				assert.ErrorIs(t, err, test.ErrorIs)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.Output, val)
		})
	}
}
