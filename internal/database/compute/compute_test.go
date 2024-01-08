package compute

import (
	"context"

	"github.com/Arkosh744/go-buddy-db/pkg/models"
	"github.com/Arkosh744/go-buddy-db/pkg/tests"
)

func (s *Suite) TestCompute_HandleQuery() {
	ctx := context.Background()

	testCases := []tests.TestCase{
		{
			Name: "Error - ParseQuery",
			ActAndAssert: func() {
				s.parser.EXPECT().ParseQuery("GET $").Return(nil, models.ErrUnknownCharacter)

				_, err := s.compute.HandleQuery(ctx, "GET $")
				s.Require().Error(err)
				s.Require().ErrorIs(err, models.ErrUnknownCharacter)
			},
		},
		{
			Name: "Error - AnalyzeQuery SET with 1 arg",
			ActAndAssert: func() {
				s.parser.EXPECT().ParseQuery("SET arg1").Return([]string{"SET", "arg1"}, nil)
				s.analyzer.EXPECT().AnalyzeQuery([]string{"SET", "arg1"}).Return(models.Query{}, errNotEnoughtTokens)

				_, err := s.compute.HandleQuery(ctx, "SET arg1")
				s.Require().Error(err)
				s.Require().ErrorIs(err, errNotEnoughtTokens)
			},
		},
		{
			Name: "Error - AnalyzeQuery SET with 3 args",
			ActAndAssert: func() {
				s.parser.EXPECT().ParseQuery("SET arg1 arg1 arg1").
					Return([]string{"SET", "arg1", "arg1", "arg1"}, nil)
				s.analyzer.EXPECT().AnalyzeQuery([]string{"SET", "arg1", "arg1", "arg1"}).
					Return(models.Query{}, models.ErrInvalidCommandArguments)

				_, err := s.compute.HandleQuery(ctx, "SET arg1 arg1 arg1")
				s.Require().Error(err)
				s.Require().ErrorIs(err, models.ErrInvalidCommandArguments)
			},
		},
		{
			Name: "Success",
			ActAndAssert: func() {
				s.parser.EXPECT().ParseQuery("SET arg1 arg2").Return([]string{"SET", "arg1", "arg2"}, nil)
				s.analyzer.EXPECT().AnalyzeQuery([]string{"SET", "arg1", "arg2"}).Return(models.Query{}, nil)

				_, err := s.compute.HandleQuery(ctx, "SET arg1 arg2")
				s.Require().NoError(err)
			},
		},
	}

	tests.RunTestCases(s, testCases)
}
