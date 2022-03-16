package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/X3NOOO/auther/values"
)

// read database, decrypt it and return struct with it
func ReadDB(path string) ([]values.Db_struct, error){
	// read
	db_encrypted, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// decrypt
	// TODO add encryption
	db_decrypted := db_encrypted

	// get json
	var db_json []values.Db_struct
	err = json.Unmarshal(db_decrypted, &db_json)
	if err != nil {
		return nil, err
	}

	return db_json, nil
}