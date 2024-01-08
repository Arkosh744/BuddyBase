package models

const (
	GetCommand = "GET"
	SetCommand = "SET"
	DelCommand = "DEL"

	GetCommandRequiredArguments = 1
	SetCommandRequiredArguments = 2
	DelCommandRequiredArguments = 1
)

var (
	ValidCommands         = []string{GetCommand, SetCommand, DelCommand}
	ValidCommandArgsCount = map[string]int{
		GetCommand: GetCommandRequiredArguments,
		SetCommand: SetCommandRequiredArguments,
		DelCommand: DelCommandRequiredArguments,
	}
)
