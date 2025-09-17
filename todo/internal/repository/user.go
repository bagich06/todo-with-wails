package repository

import (
	"context"
	"todo/internal/models"
)

type UserRepositoryInterface interface {
	CreateUser(user models.User) (int, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(id int) (models.User, error)
	UserExists(email string) (bool, error)
}

func (repo *PGRepo) CreateUser(user models.User) (int, error) {
	err := repo.pool.QueryRow(context.Background(), `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`, user.Username, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (repo *PGRepo) GetUserByEmail(email string) (user models.User, err error) {
	err = repo.pool.QueryRow(context.Background(), `SELECT id, username, email, password FROM users WHERE email=$1`, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	return user, err
}

func (repo *PGRepo) GetUserByID(id int) (user models.User, err error) {
	err = repo.pool.QueryRow(context.Background(), `SELECT id, username, email, password FROM users WHERE id=$1`, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	return user, err
}

func (repo *PGRepo) UserExists(email string) (bool, error) {
	var count int
	err := repo.pool.QueryRow(context.Background(), `SELECT COUNT(*) FROM users WHERE email=$1`, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
