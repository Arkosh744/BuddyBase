package models

import (
	"errors"
)

var (
	ErrInvalidCommand          = errors.New("invalid command")
	ErrInvalidCommandArguments = errors.New("invalid command arguments")
	ErrEmptyQuery              = errors.New("empty query")
	ErrUnknownCharacter        = errors.New("unknown character")
)
