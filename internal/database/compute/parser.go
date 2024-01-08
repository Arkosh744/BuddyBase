package compute

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseQuery(query string) ([]string, error) {
	tokens, err := newStateMachine().parse(query)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
