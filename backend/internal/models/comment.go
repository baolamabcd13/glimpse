package models

import (
    "time"

    "github.com/google/uuid"
)

type Comment struct {
    ID        uuid.UUID `json:"id" db:"id"`
    PostID    uuid.UUID `json:"postId" db:"post_id"`
    UserID    uuid.UUID `json:"userId" db:"user_id"`
    Content   string    `json:"content" db:"content"`
    CreatedAt time.Time `json:"createdAt" db:"created_at"`
    UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}