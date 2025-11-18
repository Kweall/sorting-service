package repo

import "context"

type Repository interface {
	Insert(ctx context.Context, n int) error
	ListSorted(ctx context.Context) ([]int, error)
}
