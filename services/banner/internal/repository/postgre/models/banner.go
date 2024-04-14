package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type BankAccountData struct {
}

func (a *Banner) Value() (driver.Value, error) {
	return json.Marshal(a.Content)
}

func (a *Banner) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a.Content)
}

type Banner struct {
	//UUID      uuid.UUID
	ID      int                     `db:"id"`
	Content *map[string]interface{} `db:"content"  sql:"type:jsonb"`

	//Content  []byte `db:"content"`
	IsActive bool `db:"is_active"`

	Feature int   `db:"feature"`
	Tags    []int `db:"tags"`

	CreatedAt time.Time `db:"created_at"`
	//UpdatedAt time.Time
}
