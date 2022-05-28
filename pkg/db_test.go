package hermes

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMonthsSince(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected int
	}{
		{
			name:     "365 days",
			t1:       time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			t2:       time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 12,
		},
		{
			name:     "366 days",
			t1:       time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			t2:       time.Date(2011, 1, 2, 0, 0, 0, 0, time.UTC),
			expected: 13,
		},
		{
			name:     "negative time",
			t1:       time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			t2:       time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 0,
		},
		{
			name:     "one month",
			t1:       time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			t2:       time.Date(2010, 2, 1, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "round up",
			t1:       time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			t2:       time.Date(2010, 5, 2, 23, 0, 0, 0, time.UTC),
			expected: 5,
		},
		{
			name:     "15 days",
			t1:       time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			t2:       time.Date(2010, 1, 15, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "31 days",
			t1:       time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			t2:       time.Date(2010, 1, 31, 23, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "last day to first day",
			t1:       time.Date(2010, 1, 31, 0, 0, 0, 0, time.UTC),
			t2:       time.Date(2010, 2, 1, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "last day to second day",
			t1:       time.Date(2010, 1, 31, 0, 0, 0, 0, time.UTC),
			t2:       time.Date(2010, 2, 2, 0, 0, 0, 0, time.UTC),
			expected: 2,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, monthsSince(testCase.t1, testCase.t2))
		})
	}
}

func TestYearsSince(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected int
	}{
		{
			name:     "same date",
			t1:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			t2:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 0,
		},
		{
			name:     "next day",
			t1:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			t2:       time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
			expected: 0,
		},
		{
			name:     "next month",
			t1:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			t2:       time.Date(2020, 2, 1, 0, 0, 0, 0, time.UTC),
			expected: 0,
		},
		{
			name:     "next year",
			t1:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			t2:       time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "previous year",
			t1:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			t2:       time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, yearsSince(testCase.t1, testCase.t2))
		})
	}
}
