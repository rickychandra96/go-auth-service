package repository

import (
	"auth-service/internal/domain"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByID(ctx context.Context, id string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, id string) error
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	user.BeforeCreate()

	query := `INSERT INTO users (id, email, password, name, is_active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.db.Exec(ctx, query, user.ID, user.Email, user.Password, user.Name, user.IsActive, user.CreatedAt, user.UpdatedAt)

	return err
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, email, password, name, is_active, created_at, updated_at, deleted_at FROM users WHERE email = $1 AND deleted_at IS NULL`

	var user domain.User
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	query := `SELECT id, email, password, name, is_active, created_at, updated_at, deleted_at FROM users WHERE id = $1 AND deleted_at IS NULL`

	var user domain.User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	user.UpdatedAt = time.Now()

	query := `UPDATE users SET email = $1, password = $2, name = $3, is_active = $4, updated_at = $5 WHERE id = $6 AND deleted_at IS NULL`

	cmdTag, err := r.db.Exec(ctx, query, user.Email, user.Password, user.Name, user.IsActive, user.UpdatedAt, user.ID)

	if err != nil {
		return nil, err
	}

	if cmdTag.RowsAffected() == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE users
      SET deleted_at = $1
      WHERE id = $2 AND deleted_at IS NULL`

	cmdTag, err := r.db.Exec(ctx, query, time.Now(), id)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}
