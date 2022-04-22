package models

import (
	"time"
)

type Post struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ParentID   *string   `json:"parent_id"`
	PostNumber int       `json:"post_number" gorm:"autoIncrement"`
	FKUser     string    `json:"fk_user"`
	Author     User      `json:"author" gorm:"foreignKey:FKUser"`
	Body       string    `json:"body"`
	Tag        string    `json:"tag"`
	UpVotes    int       `json:"up_votes"`
	DownVotes  int       `json:"down_votes"`
	Reputation int       `json:"reputation"`
}
