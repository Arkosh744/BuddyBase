package compute

import "github.com/Arkosh744/go-buddy-db/pkg/tests"

func (s *Suite) TestParser_ParseQuery() {
	testCases := []tests.TestCase{
		{
			Name: "Error - Empty query",
			ActAndAssert: func() {
				s.parser.EXPECT().ParseQuery("").Return(nil, errEmptyQuery)

				_, err := s.parser.ParseQuery("")
				s.Require().Error(err)
				s.Require().ErrorIs(err, errEmptyQuery)
			},
		},
		{
			Name: "Error - Unknown character",
			ActAndAssert: func() {
				s.parser.EXPECT().ParseQuery("GET $").Return(nil, errUnknownCharacter)

				_, err := s.parser.ParseQuery("GET $")
				s.Require().Error(err)
				s.Require().ErrorIs(err, errUnknownCharacter)
			},
		},
		{
			Name: "Success - only GET",
			ActAndAssert: func() {
				s.parser.EXPECT().ParseQuery("GET").Return([]string{"GET"}, nil)

				query, err := s.parser.ParseQuery("GET")
				s.Require().NoError(err)
				s.Require().Equal([]string{"GET"}, query)
			},
		},
		{
			Name: "Success - only GET with spaces",
			ActAndAssert: func() {
				s.parser.EXPECT().ParseQuery("  GET  ").Return([]string{"GET"}, nil)

				query, err := s.parser.ParseQuery("  GET  ")
				s.Require().NoError(err)
				s.Require().Equal([]string{"GET"}, query)
			},
		},
		{
			Name: "Success - only GET with spaces and tabs",
			ActAndAssert: func() {
				s.parser.EXPECT().ParseQuery("  GET \t ").Return([]string{"GET"}, nil)

				query, err := s.parser.ParseQuery("  GET \t ")
				s.Require().NoError(err)
				s.Require().Equal([]string{"GET"}, query)
			},
		},
		{
			Name: "Success - only GET with spaces, tabs, newlines",
			ActAndAssert: func() {
				s.parser.EXPECT().ParseQuery("  GET \t \n").Return([]string{"GET"}, nil)

				query, err := s.parser.ParseQuery("  GET \t \n")
				s.Require().NoError(err)
				s.Require().Equal([]string{"GET"}, query)
			},
		},
		{
			Name: "Success - GET with spaces, tabs, newlines and carriage returns",
			ActAndAssert: func() {
				s.parser.EXPECT().ParseQuery("  GET \t \n\r arg1").
					Return([]string{"GET", "arg1"}, nil)

				query, err := s.parser.ParseQuery("  GET \t \n\r arg1")
				s.Require().NoError(err)
				s.Require().Equal([]string{"GET", "arg1"}, query)
			},
		},

		{
			Name: "Success - GET with spaces, tabs, newlines and 2 args",
			ActAndAssert: func() {
				s.parser.EXPECT().ParseQuery("  GET \t \n\r arg1 \t  arg2").
					Return([]string{"GET", "arg1", "arg2"}, nil)

				query, err := s.parser.ParseQuery("  GET \t \n\r arg1 \t  arg2")
				s.Require().NoError(err)
				s.Require().Equal([]string{"GET", "arg1", "arg2"}, query)
			},
		},
	}

	tests.RunTestCases(s, testCases)
}
