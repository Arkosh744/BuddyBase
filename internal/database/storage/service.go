//go:generate mockgen -package=storage -destination=./service_mockgen.go -source=${GOFILE}
package storage

import (
	"context"
	"errors"

	"go.uber.org/zap"
)

var errNotFound = errors.New("nothing found")

type Engine interface {
	Set(ctx context.Context, key string, value string)
	Get(context.Context, string) (string, bool)
	Del(context.Context, string)
}

type Storage struct {
	engine Engine
	log    *zap.Logger
}

func NewStorage(engine Engine, l *zap.Logger) (*Storage, error) {
	if engine == nil {
		return nil, errors.New("engine not provided")
	}

	if l == nil {
		return nil, errors.New("log not provided")
	}

	return &Storage{
		engine: engine,
		log:    l,
	}, nil
}
