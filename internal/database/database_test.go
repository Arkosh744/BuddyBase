package database

import (
	"context"

	"github.com/Arkosh744/go-buddy-db/pkg/models"
	"github.com/Arkosh744/go-buddy-db/pkg/tests"
)

func (s *Suite) TestDatabase_HandleQuery() {
	ctx := context.Background()

	testCases := []tests.TestCase{
		{
			Name: "Error - ParseQuery ErrUnknownCharacter ",
			ActAndAssert: func() {
				s.compute.EXPECT().HandleQuery(ctx, "GET $").
					Return(models.Query{}, models.ErrUnknownCharacter)

				_, err := s.database.HandleQuery(ctx, "GET $")
				s.Require().Error(err)
				s.Require().ErrorIs(err, models.ErrUnknownCharacter)
			},
		},
		{
			Name: "Error - ParseQuery ErrEmptyQuery",
			ActAndAssert: func() {
				s.compute.EXPECT().HandleQuery(ctx, "GET").
					Return(models.Query{}, models.ErrEmptyQuery)

				_, err := s.database.HandleQuery(ctx, "GET")
				s.Require().Error(err)
				s.Require().ErrorIs(err, models.ErrEmptyQuery)
			},
		},
		{
			Name: "Error - AnalyzeQuery ErrInvalidCommand",
			ActAndAssert: func() {
				s.compute.EXPECT().HandleQuery(ctx, "DELETE").
					Return(models.Query{}, models.ErrInvalidCommand)

				_, err := s.database.HandleQuery(ctx, "DELETE")
				s.Require().Error(err)
				s.Require().ErrorIs(err, models.ErrInvalidCommand)
			},
		},
		{
			Name: "Error - AnalyzeQuery ErrInvalidCommand",
			ActAndAssert: func() {
				s.compute.EXPECT().HandleQuery(ctx, "GET arg1 arg2").
					Return(models.Query{}, models.ErrInvalidCommandArguments)

				_, err := s.database.HandleQuery(ctx, "GET arg1 arg2")
				s.Require().Error(err)
				s.Require().ErrorIs(err, models.ErrInvalidCommandArguments)
			},
		},
		{
			Name: "Success - GET",
			ActAndAssert: func() {
				s.compute.EXPECT().HandleQuery(ctx, "GET arg1").
					Return(models.NewQuery(models.GetCommand, []string{"arg1"}), nil)
				s.storage.EXPECT().Get(ctx, "arg1").Return("value", nil)

				value, err := s.database.HandleQuery(ctx, "GET arg1")
				s.Require().NoError(err)
				s.Require().Equal("value", value)
			},
		},
		{
			Name: "Success - SET",
			ActAndAssert: func() {
				s.compute.EXPECT().HandleQuery(ctx, "SET arg1 arg2").
					Return(models.NewQuery(models.SetCommand, []string{"arg1", "arg2"}), nil)
				s.storage.EXPECT().Set(ctx, "arg1", "arg2").Return(nil)

				value, err := s.database.HandleQuery(ctx, "SET arg1 arg2")
				s.Require().NoError(err)
				s.Require().Equal("OK", value)
			},
		},
		{
			Name: "Success - DEL",
			ActAndAssert: func() {
				s.compute.EXPECT().HandleQuery(ctx, "DEL arg1").
					Return(models.NewQuery(models.DelCommand, []string{"arg1"}), nil)
				s.storage.EXPECT().Del(ctx, "arg1").Return(nil)

				value, err := s.database.HandleQuery(ctx, "DEL arg1")
				s.Require().NoError(err)
				s.Require().Equal("OK", value)
			},
		},
	}

	tests.RunTestCases(s, testCases)
}

func (s *Suite) TestDatabase_handleDel() {
	ctx := context.Background()

	testCases := []tests.TestCase{
		{
			Name: "Success - DEL",
			ActAndAssert: func() {
				s.storage.EXPECT().Del(ctx, "arg1").Return(nil)

				value, err := s.database.handleDel(ctx, models.NewQuery(models.DelCommand, []string{"arg1"}))
				s.Require().NoError(err)
				s.Require().Equal("OK", value)
			},
		},
	}

	tests.RunTestCases(s, testCases)
}

func (s *Suite) TestDatabase_handleGet() {
	ctx := context.Background()

	testCases := []tests.TestCase{
		{
			Name: "Success - GET",
			ActAndAssert: func() {
				s.storage.EXPECT().Get(ctx, "arg1").Return("value", nil)

				value, err := s.database.handleGet(ctx, models.NewQuery(models.GetCommand, []string{"arg1"}))
				s.Require().NoError(err)
				s.Require().Equal("value", value)
			},
		},
	}

	tests.RunTestCases(s, testCases)
}

func (s *Suite) TestDatabase_handleSet() {
	ctx := context.Background()

	testCases := []tests.TestCase{
		{
			Name: "Success - SET",
			ActAndAssert: func() {
				s.storage.EXPECT().Set(ctx, "arg1", "arg2").Return(nil)

				value, err := s.database.handleSet(ctx, models.NewQuery(models.SetCommand, []string{"arg1", "arg2"}))
				s.Require().NoError(err)
				s.Require().Equal("OK", value)
			},
		},
	}

	tests.RunTestCases(s, testCases)
}
