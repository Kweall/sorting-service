package service

import (
	"context"
	"log"

	"sorting-service/internal/repo"
)

type numberService struct {
	repo repo.Repository
}

func NewNumberService(r repo.Repository) NumberService {
	return &numberService{repo: r}
}

func (s *numberService) AddNumber(ctx context.Context, n int) ([]int, error) {
	if err := s.repo.Insert(ctx, n); err != nil {
		return nil, err
	}
	list, err := s.repo.ListSorted(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("Добавили число %d => исходный массив: %v", n, list)

	return list, nil
}
