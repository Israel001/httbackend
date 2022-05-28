package models

import (
	"database/sql"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type NullString struct {
	sql.NullString
}

type NullInt64 struct {
	sql.NullInt64
}

func (nullString *NullString) MarshalJSON() ([]byte, error) {
	if !nullString.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nullString.String)
}

func (nullString *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nullString.String)
	nullString.Valid = (err == nil && nullString.String != "")
	return err
}

func (nullInt64 *NullInt64) MarshalJSON() ([]byte, error) {
	if !nullInt64.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nullInt64.Int64)
}

func (nullInt64 *NullInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nullInt64.Int64)
	nullInt64.Valid = (err == nil && nullInt64.Int64 > 0)
	return err
}

type BaseModel struct {
	ID        NullInt64      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
