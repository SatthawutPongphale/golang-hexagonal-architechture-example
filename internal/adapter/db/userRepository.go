package db

import (
	"golang-hexagonal-architechture-example/internal/domain"
	"golang-hexagonal-architechture-example/pkg/db"
)

type UserRepository struct {
	Database *db.PostgresDatabase
}

func NewUserRepository(db *db.PostgresDatabase) *UserRepository {
	return &UserRepository{Database: db}
}

func (r *UserRepository) Save(user *domain.User) (*domain.User, error) {
	// Mock saving to DB
	user.ID = 1
	return user, nil
}
