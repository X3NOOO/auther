/*
Copyright © 2022 X3NO <X3NO@disroot.org> [https://github.com/X3NOOO]

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/X3NOOO/auther/values"
)

// read database, decrypt it and return struct with it
func ReadDB(path string, key []byte)([]values.Db_struct, error){
	// read
	db_encrypted, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// decrypt
	db_decrypted, err := Decrypt(db_encrypted, key)
	if(err != nil){
		return nil, err
	}

	// get json
	var db_json []values.Db_struct
	err = json.Unmarshal(db_decrypted, &db_json)
	if err != nil {
		// assuming that db isn't encrypted
		err = json.Unmarshal(db_encrypted, &db_json)
		if(err != nil){
			return nil, err
		}
	}

	return db_json, nil
}
