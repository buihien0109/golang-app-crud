package model

import (
	"time"
)

type Post struct {
	tableName struct{}  `pg:"test.posts"`
	Id        string    `pg:"id,pk" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorId  string    `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
