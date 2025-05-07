package models

import (
    "time"

    "github.com/google/uuid"
)

type User struct {
    ID        uuid.UUID `json:"id" db:"id"`
    Username  string    `json:"username" db:"username"`
    Email     string    `json:"email" db:"email"`
    Password  string    `json:"-" db:"password_hash"`
    FullName  string    `json:"fullName" db:"full_name"`
    Bio       string    `json:"bio" db:"bio"`
    AvatarURL string    `json:"avatarUrl" db:"avatar_url"`
    CreatedAt time.Time `json:"createdAt" db:"created_at"`
    UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}