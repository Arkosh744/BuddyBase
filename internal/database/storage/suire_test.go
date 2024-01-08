package storage

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type Suite struct {
	suite.Suite

	ctrl    *gomock.Controller
	engine  *MockEngine
	storage *Storage
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
	s.engine = NewMockEngine(s.ctrl)

	storage, err := NewStorage(s.engine, zap.NewNop())
	s.Require().NoError(err)

	s.storage = storage
}
