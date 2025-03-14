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

	"gorm.io/driver/sqlserver"
)

func SQLServer(database string, host string, password string, port int, user string) *SQL {
	dsn := "sqlserver://"

	if user != "" && password != "" {
		dsn += fmt.Sprintf("%s:%s@", user, password)
	}
	if host != "" {
		dsn += fmt.Sprintf("%s:%d", host, port)
	}
	if database != "" {
		dsn += fmt.Sprintf("?database=%s", database)
	}

	return &SQL{
		driver: sqlserver.Open(dsn),
	}
}
