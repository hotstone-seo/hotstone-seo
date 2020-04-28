package repository

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONMap is a representation of JSON in Golang's map
type JSONMap map[string]interface{}

// Value is the representation of JSONMap in database ([]byte)
func (j JSONMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan is how JSONMap value should be populated from database
func (j *JSONMap) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("database value is not []byte")
	}
	return json.Unmarshal(b, &j)
}
