// repository/user_repository.go

package repository

import (
	"context"
	db "njajal-go/db/sqlc"
)

type UserRepository interface {
	CreateUser(ctx context.Context, email string, password string) error
	FindUserByEmail(ctx context.Context, email string) (db.User, error)
	ResetPassword(ctx context.Context, email string, password string) error
}

type UserRepositoryImpl struct{
	// Add the db instance here
	db *db.Queries
}

func NewUserRepository(conn *db.Queries) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: conn,
	}
}

func (ur *UserRepositoryImpl) CreateUser(ctx context.Context, email string, password string) error {
	// Use the queries instance to interact with the database
	err := ur.db.CreateUsers(ctx, db.CreateUsersParams{
		Email:    email,
		Password: password,
	})

	return err
}

func (ur *UserRepositoryImpl) FindUserByEmail(ctx context.Context, email string) (db.User, error) {
	user, err := ur.db.FindUserByEmail(ctx, email)
	return user, err
}

func (ur *UserRepositoryImpl) ResetPassword(ctx context.Context, email string, password string) error {
	err := ur.db.ResetPassword(ctx, db.ResetPasswordParams{
		Email:    email,
		Password: password,
	})
	return err
}