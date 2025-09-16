package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PGRepo struct {
	pool *pgxpool.Pool
}

func NewPGRepo(connStr string) (*PGRepo, error) {
	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return &PGRepo{pool: pool}, nil
}

func (repo *PGRepo) Close() {
	repo.pool.Close()
}

func (repo *PGRepo) GetPool() *pgxpool.Pool {
	return repo.pool
}
