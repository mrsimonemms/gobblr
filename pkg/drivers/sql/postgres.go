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

package sql

import (
	"fmt"
	"strings"

	"gorm.io/driver/postgres"
)

func PostgreSQL(database string, host string, password string, port int, user string) *SQL {
	var dsn []string

	if user != "" {
		dsn = append(dsn, fmt.Sprintf("user=%s", user))
	}
	if password != "" {
		dsn = append(dsn, fmt.Sprintf("password=%s", password))
	}
	if host != "" {
		dsn = append(dsn, fmt.Sprintf("host=%s", host))
	}
	dsn = append(dsn, fmt.Sprintf("port=%d", port))
	dsn = append(dsn, fmt.Sprintf("dbname=%s", database))

	return &SQL{
		driver: postgres.Open(strings.Join(dsn, " ")),
	}
}
