package repository

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Attrs is map with string key-value
type Attrs map[string]string

// Value of attrs in database perspective
func (a Attrs) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan to set attrrs when sql scan
func (a *Attrs) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
