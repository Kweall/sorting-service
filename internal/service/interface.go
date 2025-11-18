package service

import "context"

type NumberService interface {
	AddNumber(ctx context.Context, n int) ([]int, error)
}
