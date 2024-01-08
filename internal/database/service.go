package database

import (
	"context"

	"github.com/Arkosh744/go-buddy-db/internal/database/compute"
	"go.uber.org/zap"
)

type ComputerLayer interface {
	HandleQuery(ctx context.Context, queryStr string) (query compute.Query, err error)
}

type StorageLayer interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}

type Database struct {
	computer ComputerLayer
	storage  StorageLayer
	logger   *zap.Logger
}
