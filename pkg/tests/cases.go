package tests

type TestCase struct {
	Name         string
	Arrange      func()
	ActAndAssert func()
	AfterTest    func()

	SubTests []TestCase
}

type Suite interface {
	Run(name string, subtest func()) bool
}

func RunTestCases(s Suite, testCases []TestCase) {
	for _, tc := range testCases {
		RunTestCase(s, tc)
	}
}

func RunTestCase(s Suite, tc TestCase) {
	s.Run(tc.Name, func() {
		if tc.Arrange != nil {
			tc.Arrange()
		}

		if tc.AfterTest != nil {
			defer tc.AfterTest()
		}

		if tc.ActAndAssert != nil {
			tc.ActAndAssert()
		}

		RunTestCases(s, tc.SubTests)
	})
}
