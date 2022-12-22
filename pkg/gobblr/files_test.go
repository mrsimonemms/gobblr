//go:build !e2e
// +build !e2e

package gobblr

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseDates(t *testing.T) {
	testCases := []struct {
		Name     string
		Input    interface{}
		Expected interface{}
	}{
		{
			Name:     "boolean true",
			Input:    true,
			Expected: true,
		},
		{
			Name:     "boolean false",
			Input:    false,
			Expected: false,
		},
		{
			Name:     "rfc3339 date",
			Input:    "2022-12-23T10:20:30Z",
			Expected: time.Date(2022, 12, 23, 10, 20, 30, 0, time.UTC),
		},
		{
			Name:     "YYYY-MM-DD date",
			Input:    "2022-12-23",
			Expected: time.Date(2022, 12, 23, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:     "string",
			Input:    "hello world",
			Expected: "hello world",
		},
		{
			Name:     "numeric string",
			Input:    "2022",
			Expected: "2022",
		},
		{
			Name:     "number",
			Input:    2022,
			Expected: 2022,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			require.Equal(t, parseDates(testCase.Input), testCase.Expected)
		})
	}
}
