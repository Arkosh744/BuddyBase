package database

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type Suite struct {
	suite.Suite

	ctrl     *gomock.Controller
	compute  *MockComputeLayer
	storage  *MockStorageLayer
	database *Database
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
	s.compute = NewMockComputeLayer(s.ctrl)
	s.storage = NewMockStorageLayer(s.ctrl)

	database, err := NewDatabase(s.compute, s.storage, zap.NewNop())
	s.Require().NoError(err)

	s.database = database
}
