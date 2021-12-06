package model

import (
	"time"
)

type User struct {
	tableName struct{}  `pg:"test.users"`
	Id        string    `pg:"id,pk" json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
