package handler

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

type svcMock struct {
	out []int
	err error
}

func (s svcMock) AddNumber(ctx context.Context, n int) ([]int, error) {
	return s.out, s.err
}

func TestHandlerJSON(t *testing.T) {
	h := NewHTTPHandler(svcMock{out: []int{1, 2, 3}})
	req := httptest.NewRequest(http.MethodPost, "/numbers", bytes.NewBufferString(`{"value":3}`))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	h.AddNumber(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status %d", rr.Code)
	}
}
