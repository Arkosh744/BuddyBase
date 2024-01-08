package compute

import (
	"context"

	"github.com/Arkosh744/go-buddy-db/pkg/tests"
)

func (s *Suite) TestCompute_HandleQuery() {
	ctx := context.Background()

	testCases := []tests.TestCase{
		{
			Name: "Error - ParseQuery",
			ActAndAssert: func() {
				s.parser.EXPECT().ParseQuery("GET $").Return(nil, errUnknownCharacter)

				_, err := s.compute.HandleQuery(ctx, "GET $")
				s.Require().Error(err)
				s.Require().ErrorIs(err, errUnknownCharacter)
			},
		},
		{
			Name: "Error - AnalyzeQuery SET with 1 arg",
			ActAndAssert: func() {
				s.parser.EXPECT().ParseQuery("SET arg1").Return([]string{"SET", "arg1"}, nil)
				s.analyzer.EXPECT().AnalyzeQuery([]string{"SET", "arg1"}).Return(Query{}, errNotEnoughtTokens)

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
					Return(Query{}, errInvalidCommandArguments)

				_, err := s.compute.HandleQuery(ctx, "SET arg1 arg1 arg1")
				s.Require().Error(err)
				s.Require().ErrorIs(err, errInvalidCommandArguments)
			},
		},
		{
			Name: "Success",
			ActAndAssert: func() {
				s.parser.EXPECT().ParseQuery("SET arg1 arg2").Return([]string{"SET", "arg1", "arg2"}, nil)
				s.analyzer.EXPECT().AnalyzeQuery([]string{"SET", "arg1", "arg2"}).Return(Query{}, nil)

				_, err := s.compute.HandleQuery(ctx, "SET arg1 arg2")
				s.Require().NoError(err)
			},
		},
	}

	tests.RunTestCases(s, testCases)
}
