package repo

import (
	"context"
	"database/sql"
	"errors"
)

type pgRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &pgRepository{db: db}
}

func (p *pgRepository) Insert(ctx context.Context, n int) error {
	if p.db == nil {
		return errors.New("db is nil")
	}
	_, err := p.db.ExecContext(ctx, `INSERT INTO numbers(value, created_at) VALUES($1, now())`, n)
	return err
}

func (p *pgRepository) ListSorted(ctx context.Context) ([]int, error) {
	rows, err := p.db.QueryContext(ctx, `SELECT value FROM numbers ORDER BY value ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []int
	for rows.Next() {
		var v int
		if err := rows.Scan(&v); err != nil {
			return nil, err
		}
		res = append(res, v)
	}
	return res, rows.Err()
}
