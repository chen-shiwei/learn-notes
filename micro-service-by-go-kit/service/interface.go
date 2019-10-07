package service

import (
	"context"
)

type StringServiceInterface interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}
