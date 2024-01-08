package compute

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
)

const (
	GetCommand = "GET"
	SetCommand = "SET"
	DelCommand = "DEL"

	GetCommandRequiredArguments = 1
	SetCommandRequiredArguments = 2
	DelCommandRequiredArguments = 1
)

var (
	errNotEnoughtTokens = errors.New("not enough tokens provided")
)

var (
	validCommands         = []string{GetCommand, SetCommand, DelCommand}
	validCommandArgsCount = map[string]int{
		GetCommand: GetCommandRequiredArguments,
		SetCommand: SetCommandRequiredArguments,
		DelCommand: DelCommandRequiredArguments,
	}
)

type Analyzer struct{}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) AnalyzeQuery(tokens []string) (Query, error) {
	if len(tokens) < 2 {
		return Query{}, errNotEnoughtTokens
	}

	command := tokens[0]
	args := tokens[1:]
	query := NewQuery(command, args)

	if err := analyzeCommand(query); err != nil {
		return Query{}, err
	}

	return query, nil
}

func analyzeCommand(query Query) error {
	if !lo.Contains(validCommands, query.Command()) {
		return fmt.Errorf("%w: got %s", errInvalidCommand, query.Command())
	}

	if len(query.Arguments()) != validCommandArgsCount[query.Command()] {
		return fmt.Errorf("%w: got %d", errInvalidCommandArguments, len(query.Arguments()))
	}

	return nil
}
