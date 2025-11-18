package service

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
)

type mockRepo struct {
	inserted []int
	list     []int
	err      error
}

func (m *mockRepo) Insert(ctx context.Context, n int) error {
	if m.err != nil {
		return m.err
	}
	m.inserted = append(m.inserted, n)
	return nil
}

func (m *mockRepo) ListSorted(ctx context.Context) ([]int, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.list, nil
}

func TestAddNumber_Success(t *testing.T) {
	m := &mockRepo{
		list: []int{1, 2, 3},
	}
	s := NewNumberService(m)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	out, err := s.AddNumber(ctx, 4)
	if err != nil {
		t.Fatalf("expected nil err, got %v", err)
	}
	if !reflect.DeepEqual(out, m.list) {
		t.Fatalf("expected %v got %v", m.list, out)
	}
}
func TestAddNumber_RepoError(t *testing.T) {
	m := &mockRepo{err: errors.New("boom")}
	s := NewNumberService(m)
	ctx := context.Background()
	_, err := s.AddNumber(ctx, 5)
	if err == nil {
		t.Fatalf("expected error")
	}
}
