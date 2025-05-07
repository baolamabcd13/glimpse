package models

import (
    "time"

    "github.com/google/uuid"
)

type Post struct {
    ID        uuid.UUID `json:"id" db:"id"`
    UserID    uuid.UUID `json:"userId" db:"user_id"`
    Caption   string    `json:"caption" db:"caption"`
    ImageURL  string    `json:"imageUrl" db:"image_url"`
    LikesCount int       `json:"likesCount" db:"likes_count"`
    CreatedAt time.Time `json:"createdAt" db:"created_at"`
    UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}