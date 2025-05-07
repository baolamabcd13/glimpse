package repository

import (
    "context"
    "errors"
    "fmt"

    "github.com/baolamabcd13/glimpse/internal/database"
    "github.com/baolamabcd13/glimpse/internal/models"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v5"
)

// UserRepository handles database operations for users
type UserRepository struct {
    db *database.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *database.DB) *UserRepository {
    return &UserRepository{db: db}
}

// Create inserts a new user into the database
func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
    query := `
        INSERT INTO users (id, username, email, password_hash, full_name, bio, avatar_url, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    `

    _, err := r.db.Pool.Exec(ctx, query,
        user.ID, user.Username, user.Email, user.Password, user.FullName,
        user.Bio, user.AvatarURL, user.CreatedAt, user.UpdatedAt,
    )
    if err != nil {
        return fmt.Errorf("failed to create user: %w", err)
    }

    return nil
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
    query := `
        SELECT id, username, email, password_hash, full_name, bio, avatar_url, created_at, updated_at
        FROM users
        WHERE id = $1
    `

    var user models.User
    err := r.db.Pool.QueryRow(ctx, query, id).Scan(
        &user.ID, &user.Username, &user.Email, &user.Password, &user.FullName,
        &user.Bio, &user.AvatarURL, &user.CreatedAt, &user.UpdatedAt,
    )
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return nil, fmt.Errorf("user not found")
        }
        return nil, fmt.Errorf("failed to get user: %w", err)
    }

    return &user, nil
}

// GetByEmail retrieves a user by email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
    query := `
        SELECT id, username, email, password_hash, full_name, bio, avatar_url, created_at, updated_at
        FROM users
        WHERE email = $1
    `

    var user models.User
    err := r.db.Pool.QueryRow(ctx, query, email).Scan(
        &user.ID, &user.Username, &user.Email, &user.Password, &user.FullName,
        &user.Bio, &user.AvatarURL, &user.CreatedAt, &user.UpdatedAt,
    )
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return nil, fmt.Errorf("user not found")
        }
        return nil, fmt.Errorf("failed to get user: %w", err)
    }

    return &user, nil
}

// GetByUsername retrieves a user by username
func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
    query := `
        SELECT id, username, email, password_hash, full_name, bio, avatar_url, created_at, updated_at
        FROM users
        WHERE username = $1
    `

    var user models.User
    err := r.db.Pool.QueryRow(ctx, query, username).Scan(
        &user.ID, &user.Username, &user.Email, &user.Password, &user.FullName,
        &user.Bio, &user.AvatarURL, &user.CreatedAt, &user.UpdatedAt,
    )
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return nil, fmt.Errorf("user not found")
        }
        return nil, fmt.Errorf("failed to get user: %w", err)
    }

    return &user, nil
}

// Update updates a user in the database
func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
    query := `
        UPDATE users
        SET username = $2, email = $3, full_name = $4, bio = $5, avatar_url = $6, updated_at = $7
        WHERE id = $1
    `

    _, err := r.db.Pool.Exec(ctx, query,
        user.ID, user.Username, user.Email, user.FullName,
        user.Bio, user.AvatarURL, user.UpdatedAt,
    )
    if err != nil {
        return fmt.Errorf("failed to update user: %w", err)
    }

    return nil
}