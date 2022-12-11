// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Posts struct {
	ID        uuid.UUID 	`json:"id"`
	UserID    uuid.UUID     `json:"user_id"`
	Content   string        `json:"content"`
	IsActive  bool          `json:"is_active"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
}
