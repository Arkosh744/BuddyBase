package compute

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Arkosh744/go-buddy-db/pkg/models"
)

const (
	initialState = iota
	charFoundState
	whiteSpaceFoundState
)

var errNotValidTransition = errors.New("not valid transition")

type transitionAction func(ch rune, stringBuffer *strings.Builder, tokens []string) ([]string, int)

// transition[fromState][ToState].
var transitions = map[int]map[int]transitionAction{
	initialState: {
		charFoundState:       handleCharFound,
		whiteSpaceFoundState: handleWhiteSpaceFound,
	},
	charFoundState: {
		charFoundState:       handleCharFound,
		whiteSpaceFoundState: handleWhiteSpaceFound,
	},
	whiteSpaceFoundState: {
		charFoundState:       handleCharFound,
		whiteSpaceFoundState: handleWhiteSpaceFound,
	},
}

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseQuery(query string) ([]string, error) {
	if len(query) == 0 {
		return nil, models.ErrEmptyQuery
	}

	var state int
	var tokens []string
	stringBuffer := &strings.Builder{}

	for _, ch := range query {
		toState, err := getCharType(ch)
		if err != nil {
			return nil, err
		}

		transition, ok := transitions[state][toState]
		if !ok {
			return nil, errNotValidTransition
		}

		tokens, state = transition(ch, stringBuffer, tokens)
	}

	if action, exists := transitions[state][whiteSpaceFoundState]; exists {
		tokens, _ = action(' ', stringBuffer, tokens)
	}

	return tokens, nil
}

func getCharType(ch rune) (int, error) {
	switch {
	case isValidCharacter(ch):
		return charFoundState, nil
	case isSpace(ch):
		return whiteSpaceFoundState, nil
	}

	return 0, models.ErrUnknownCharacter
}

func handleCharFound(ch rune, stringBuffer *strings.Builder, tokens []string) ([]string, int) {
	stringBuffer.WriteRune(ch)
	return tokens, charFoundState
}

func handleWhiteSpaceFound(_ rune, stringBuffer *strings.Builder, tokens []string) ([]string, int) {
	if stringBuffer.Len() > 0 {
		tokens = append(tokens, stringBuffer.String())
		stringBuffer.Reset()
	}

	return tokens, whiteSpaceFoundState
}

func isSpace(ch rune) bool {
	return unicode.IsSpace(ch)
}

func isValidCharacter(ch rune) bool {
	return unicode.IsLetter(ch) ||
		unicode.IsDigit(ch) ||
		ch == '_' ||
		ch == '/' ||
		ch == '*'
}
