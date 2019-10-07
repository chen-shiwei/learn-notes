package service

import (
	"context"
	"errors"
	"strings"
)

type StringService struct {
}

var ErrEmpty = errors.New("empty string")

func (StringService) Uppercase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}
func (StringService) Count(_ context.Context, s string) int {
	return len(s)
}
