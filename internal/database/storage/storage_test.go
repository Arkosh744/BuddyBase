package storage

import (
	"context"

	"github.com/Arkosh744/go-buddy-db/internal/database/storage/mem"
	"github.com/Arkosh744/go-buddy-db/pkg/tests"
	"go.uber.org/zap"
)

func (s *Suite) TestNewStorage() {
	s.Run("err - engine not provided", func() {
		_, err := NewStorage(nil, zap.NewNop())
		s.Require().Error(err)
	})

	s.Run("err - log not provided", func() {
		_, err := NewStorage(mem.NewEngine(), nil)
		s.Require().Error(err)
	})

	s.Run("success", func() {
		_, err := NewStorage(mem.NewEngine(), zap.NewNop())
		s.Require().NoError(err)
	})
}

func (s *Suite) TestStorage_Del() {
	var (
		ctx   = context.Background()
		key   = "test-del-key"
		value = "test-del-value"
	)

	testCases := []tests.TestCase{
		{
			Name: "success",
			Arrange: func() {
				s.engine.EXPECT().Set(ctx, key, value)
				err := s.storage.Set(ctx, key, value)
				s.Require().NoError(err)
			},
			ActAndAssert: func() {
				s.engine.EXPECT().Del(ctx, key)
				err := s.storage.Del(ctx, key)
				s.Require().NoError(err)
			},
		},
	}

	tests.RunTestCases(s, testCases)
}

func (s *Suite) TestStorage_Get() {
	var (
		ctx   = context.Background()
		key   = "test-get-key"
		value = "test-get-value"
	)

	testCases := []tests.TestCase{
		{
			Name: "success",
			Arrange: func() {
				s.engine.EXPECT().Set(ctx, key, value)
				err := s.storage.Set(ctx, key, value)
				s.Require().NoError(err)
			},
			ActAndAssert: func() {
				s.engine.EXPECT().Get(ctx, key).Return(value, true)
				val, err := s.storage.Get(ctx, key)
				s.Require().NoError(err)
				s.Require().Equal(value, val)
			},
		},
		{
			Name: "err - not found",
			ActAndAssert: func() {
				s.engine.EXPECT().Get(ctx, key).Return("", false)
				_, err := s.storage.Get(ctx, key)
				s.Require().Error(err)
			},
		},
	}

	tests.RunTestCases(s, testCases)
}

func (s *Suite) TestStorage_Set() {
	var (
		ctx   = context.Background()
		key   = "test-set-key"
		value = "test-set-value"
	)

	testCases := []tests.TestCase{
		{
			Name: "success",
			ActAndAssert: func() {
				s.engine.EXPECT().Set(ctx, key, value)
				err := s.storage.Set(ctx, key, value)
				s.Require().NoError(err)
			},
		},
	}

	tests.RunTestCases(s, testCases)
}
