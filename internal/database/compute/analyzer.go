package compute

import (
	"errors"
	"fmt"

	"github.com/Arkosh744/go-buddy-db/pkg/models"
	"github.com/samber/lo"
)

var errNotEnoughtTokens = errors.New("not enough tokens provided")

type Analyzer struct{}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) AnalyzeQuery(tokens []string) (models.Query, error) {
	if len(tokens) < 2 {
		return models.Query{}, errNotEnoughtTokens
	}

	command := tokens[0]
	args := tokens[1:]
	query := models.NewQuery(command, args)

	if err := analyzeCommand(query); err != nil {
		return models.Query{}, err
	}

	return query, nil
}

func analyzeCommand(query models.Query) error {
	if !lo.Contains(models.ValidCommands, query.Command()) {
		return fmt.Errorf("%w: got %s", models.ErrInvalidCommand, query.Command())
	}

	if len(query.Arguments()) != models.ValidCommandArgsCount[query.Command()] {
		return fmt.Errorf("%w: got %d", models.ErrInvalidCommandArguments, len(query.Arguments()))
	}

	return nil
}
