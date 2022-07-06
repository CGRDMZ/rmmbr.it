package services

import (
	"context"
	"fmt"
	"github.com/CGRDMZ/rmmbrit-api/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserService struct {
	Db *pgxpool.Pool
}


type CreateUserParams struct {
	Username string
	Email string
} 

func (us *UserService) CreateUser(ctx context.Context, params CreateUserParams) (*models.User, error) {
	var err error
	tx, err := us.Db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("something wrong happened while beginning transaction: %w", err)
	}


	_, err = tx.Exec(context.Background(), "INSERT INTO users (username, email) VALUES ($1, $2)", params.Username, params.Email)
	if err != nil {
		return nil, fmt.Errorf("something happened while creating a new 'user': %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not commit the transaction: %w", err)
	}

	user, err := us.FindByUsername(ctx, params.Username)
	if err != nil {
		return nil, fmt.Errorf("something wrong happened while finding user: %w", err)
	}

	return user, nil
	
}

func (us *UserService) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var err error
	tx, err := us.Db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("something wrong happened while beginning transaction: %w", err)
	}

	var user models.User
	err = tx.QueryRow(context.Background(), "SELECT id, username, email, password, created_at, updated_at FROM users WHERE username = $1", username).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("something happened while finding a user: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not commit the transaction: %w", err)
	}

	return &user, nil
}