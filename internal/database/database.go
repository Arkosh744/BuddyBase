//go:generate mockgen -package=database -destination=./database_mockgen.go -source=${GOFILE}
package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/Arkosh744/go-buddy-db/pkg/models"
	"go.uber.org/zap"
)

var (
	ErrComputeNotProvided = errors.New("compute not provided")
	ErrStorageNotProvided = errors.New("storage not provided")
	ErrLogNotProvided     = errors.New("log not provided")
)

type ComputeLayer interface {
	HandleQuery(ctx context.Context, queryStr string) (models.Query, error)
}

type StorageLayer interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}

type Database struct {
	compute ComputeLayer
	storage StorageLayer
	log     *zap.Logger
}

func NewDatabase(compute ComputeLayer, storage StorageLayer, log *zap.Logger) (*Database, error) {
	if compute == nil {
		return nil, ErrComputeNotProvided
	}

	if storage == nil {
		return nil, ErrStorageNotProvided
	}

	if log == nil {
		return nil, ErrLogNotProvided
	}

	return &Database{
		compute: compute,
		storage: storage,
		log:     log,
	}, nil
}

func (d *Database) HandleQuery(ctx context.Context, queryStr string) (string, error) {
	d.log.Info("handling query", zap.String("query", queryStr))

	query, err := d.compute.HandleQuery(ctx, queryStr)
	if err != nil {
		return "", err
	}

	switch query.Command() {
	case models.GetCommand:
		return d.handleGet(ctx, query)
	case models.SetCommand:
		return d.handleSet(ctx, query)
	case models.DelCommand:
		return d.handleDel(ctx, query)
	default:
		return "", fmt.Errorf("unknown command: %s", query.Command())
	}
}

func (d *Database) handleGet(ctx context.Context, query models.Query) (string, error) {
	value, err := d.storage.Get(ctx, query.Arguments()[0])
	if err != nil {
		return "", fmt.Errorf("handle get: %w", err)
	}

	return value, nil
}

func (d *Database) handleSet(ctx context.Context, query models.Query) (string, error) {
	if err := d.storage.Set(ctx, query.Arguments()[0], query.Arguments()[1]); err != nil {
		return "", fmt.Errorf("handle set: %w", err)
	}

	return "OK", nil
}

func (d *Database) handleDel(ctx context.Context, query models.Query) (string, error) {
	if err := d.storage.Del(ctx, query.Arguments()[0]); err != nil {
		return "", fmt.Errorf("handle del: %w", err)
	}

	return "OK", nil
}
