package compute

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type Suite struct {
	suite.Suite

	ctrl     *gomock.Controller
	compute  *Compute
	parser   *Mockparser
	analyzer *Mockanalyzer
}

func TestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(Suite))
}

func (s *Suite) TearDownSuite() {
	s.ctrl.Finish()
}

func (s *Suite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	s.analyzer = NewMockanalyzer(s.ctrl)
	s.parser = NewMockparser(s.ctrl)

	compute, err := NewCompute(s.parser, s.analyzer, zap.NewNop())
	s.Require().NoError(err)

	s.compute = compute
}
