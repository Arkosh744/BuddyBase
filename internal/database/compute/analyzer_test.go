package compute

import (
	"github.com/Arkosh744/go-buddy-db/pkg/models"
	"github.com/Arkosh744/go-buddy-db/pkg/tests"
)

func (s *Suite) TestAnalyzer_AnalyzeQuery() {
	testCases := []tests.TestCase{
		{
			Name: "Error - GET w/o args",
			ActAndAssert: func() {
				s.analyzer.EXPECT().AnalyzeQuery([]string{"GET"}).Return(models.Query{}, errNotEnoughtTokens)

				_, err := s.analyzer.AnalyzeQuery([]string{"GET"})
				s.Require().Error(err)
				s.Require().ErrorIs(err, errNotEnoughtTokens)
			},
		},
		{
			Name: "Error - GET with 2 args",
			ActAndAssert: func() {
				s.analyzer.EXPECT().AnalyzeQuery([]string{"GET", "arg1", "arg2"}).
					Return(models.Query{}, models.ErrInvalidCommandArguments)

				_, err := s.analyzer.AnalyzeQuery([]string{"GET", "arg1", "arg2"})
				s.Require().Error(err)
				s.Require().ErrorIs(err, models.ErrInvalidCommandArguments)
			},
		},
		{
			Name: "Error - SET w/o args",
			ActAndAssert: func() {
				s.analyzer.EXPECT().AnalyzeQuery([]string{"SET"}).Return(models.Query{}, errNotEnoughtTokens)

				_, err := s.analyzer.AnalyzeQuery([]string{"SET"})
				s.Require().Error(err)
				s.Require().ErrorIs(err, errNotEnoughtTokens)
			},
		},
		{
			Name: "Error - SET with 1 arg",
			ActAndAssert: func() {
				s.analyzer.EXPECT().AnalyzeQuery([]string{"SET", "arg1"}).Return(models.Query{}, models.ErrInvalidCommandArguments)

				_, err := s.analyzer.AnalyzeQuery([]string{"SET", "arg1"})
				s.Require().Error(err)
				s.Require().ErrorIs(err, models.ErrInvalidCommandArguments)
			},
		},
		{
			Name: "Error - DEL w/o args",
			ActAndAssert: func() {
				s.analyzer.EXPECT().AnalyzeQuery([]string{"DEL"}).Return(models.Query{}, errNotEnoughtTokens)

				_, err := s.analyzer.AnalyzeQuery([]string{"DEL"})
				s.Require().Error(err)
				s.Require().ErrorIs(err, errNotEnoughtTokens)
			},
		},
		{
			Name: "Error - DEL with 2 args",
			ActAndAssert: func() {
				s.analyzer.EXPECT().AnalyzeQuery([]string{"DEL", "arg1", "arg2"}).
					Return(models.Query{}, models.ErrInvalidCommandArguments)

				_, err := s.analyzer.AnalyzeQuery([]string{"DEL", "arg1", "arg2"})
				s.Require().Error(err)
				s.Require().ErrorIs(err, models.ErrInvalidCommandArguments)
			},
		},
		{
			Name: "Success - GET",
			ActAndAssert: func() {
				s.analyzer.EXPECT().AnalyzeQuery([]string{"GET", "arg1"}).
					Return(models.NewQuery(models.GetCommand, []string{"arg1"}), nil)

				query, err := s.analyzer.AnalyzeQuery([]string{"GET", "arg1"})
				s.Require().NoError(err)
				s.Require().Equal("GET", query.Command())
				s.Require().Equal([]string{"arg1"}, query.Arguments())
			},
		},
		{
			Name: "Success - SET",
			ActAndAssert: func() {
				s.analyzer.EXPECT().AnalyzeQuery([]string{"SET", "arg1", "arg2"}).
					Return(models.NewQuery(models.SetCommand, []string{"arg1", "arg2"}), nil)

				query, err := s.analyzer.AnalyzeQuery([]string{"SET", "arg1", "arg2"})
				s.Require().NoError(err)
				s.Require().Equal("SET", query.Command())
				s.Require().Equal([]string{"arg1", "arg2"}, query.Arguments())
			},
		},
		{
			Name: "Success - DEL",
			ActAndAssert: func() {
				s.analyzer.EXPECT().AnalyzeQuery([]string{"DEL", "arg1"}).
					Return(models.NewQuery(models.DelCommand, []string{"arg1"}), nil)

				query, err := s.analyzer.AnalyzeQuery([]string{"DEL", "arg1"})
				s.Require().NoError(err)
				s.Require().Equal("DEL", query.Command())
				s.Require().Equal([]string{"arg1"}, query.Arguments())
			},
		},
	}

	tests.RunTestCases(s, testCases)
}
