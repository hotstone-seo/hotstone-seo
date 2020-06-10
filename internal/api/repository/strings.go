package repository

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Strings is slice of string type in json
type Strings []string

// Value is the representation of Strings in database ([]byte)
func (j Strings) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan is how Strings value should be populated from database
func (j *Strings) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("database value is not []byte")
	}
	return json.Unmarshal(b, &j)
}
