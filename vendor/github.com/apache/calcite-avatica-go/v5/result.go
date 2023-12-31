/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to you under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package avatica

import "errors"

type result struct {
	affectedRows int64
	insertId     int64
}

// LastInsertId returns the database's auto-generated ID
// after, for example, an INSERT into a table with primary
// key.
func (r *result) LastInsertId() (int64, error) {
	return 0, errors.New("use 'SELECT CURRENT VALUE FOR your.sequence' to get the last inserted id. For more information, see: https://phoenix.apache.org/sequences.html")
}

// RowsAffected returns the number of rows affected by the
// query.
func (r *result) RowsAffected() (int64, error) {
	return r.affectedRows, nil
}
