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

// The response of the "data" function is what's put in the database
function data() {
  return {
    meta: {
      createdKey: 'created_at',
      updatedKey: 'updated_at'
    },
    data: [
      {
        item: 1,
        some_date: new Date(),
      },
      {
        item: 2,
        some_date: new Date(),
      },
    ]
  }
}
