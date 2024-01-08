//go:generate mockgen -package=compute -destination=./compute_mockgen.go -source=${GOFILE}
package compute

import (
	"context"
	"errors"
	"fmt"

	"github.com/Arkosh744/go-buddy-db/pkg/models"

	"go.uber.org/zap"
)

type parser interface {
	ParseQuery(query string) ([]string, error)
}

type analyzer interface {
	AnalyzeQuery(tokens []string) (models.Query, error)
}

type Compute struct {
	parser   parser
	analyzer analyzer
	logger   *zap.Logger
}

func NewCompute(parser parser, analyzer analyzer, logger *zap.Logger) (*Compute, error) {
	if parser == nil {
		return nil, errors.New("parser not provided")
	}

	if analyzer == nil {
		return nil, errors.New("analyzer not provided")
	}

	if logger == nil {
		return nil, errors.New("logger not provided")
	}

	return &Compute{
		parser:   parser,
		analyzer: analyzer,
		logger:   logger,
	}, nil
}

func (c *Compute) HandleQuery(_ context.Context, queryStr string) (models.Query, error) {
	tokens, err := c.parser.ParseQuery(queryStr)
	if err != nil {
		return models.Query{}, fmt.Errorf("parse: %w", err)
	}

	query, err := c.analyzer.AnalyzeQuery(tokens)
	if err != nil {
		return models.Query{}, fmt.Errorf("analyze: %w", err)
	}

	return query, nil
}
