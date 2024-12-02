package main

import "testing"

func TestReportSafetyChecking1(t *testing.T) {
	cases1 := []struct {
		Name           string
		Reports        [][]int
		NumSafeReports int
	}{
		{
			"all safe ascending",
			[][]int{[]int{1, 2}},
			1,
		},
		{
			"all safe descending",
			[][]int{[]int{7, 4, 3, 1}},
			1,
		},
		{
			"unsafe because of big difference ascending",
			[][]int{[]int{1, 2, 6}},
			0,
		},
		{
			"unsafe because of big difference descending",
			[][]int{[]int{7, 1}},
			0,
		},
		{
			"unsafe because of equal elements",
			[][]int{[]int{1, 1}},
			0,
		},
		{
			"unsafe because of changing order",
			[][]int{[]int{1, 2, 1}},
			0,
		},
		{
			"aoc example",
			[][]int{
				[]int{7, 6, 4, 2, 1},
				[]int{1, 2, 7, 8, 9},
				[]int{9, 7, 6, 2, 1},
				[]int{1, 3, 2, 4, 5},
				[]int{8, 6, 4, 4, 1},
				[]int{1, 3, 6, 7, 9},
			},
			2,
		},
	}

	for _, test := range cases1 {
		t.Run(test.Name, func(t *testing.T) {
			got := ReportSafetyChecking(test.Reports, false)
			want := test.NumSafeReports

			if got != want {
				t.Errorf("got %d expected %d given reports %v", got, want, test.Reports)
			}
		})
	}
}

func TestReportSafetyChecking2(t *testing.T) {
	cases2 := []struct {
		Name           string
		Reports        [][]int
		NumSafeReports int
	}{
		{
			"safe with first wrong",
			[][]int{[]int{3, 2, 3, 4}},
			1,
		},
		{
			"safe with second wrong",
			[][]int{[]int{1, 0, 2, 3}},
			1,
		},
		{
			"unsafe because two wrong",
			[][]int{[]int{1, 2, 1, 2}},
			0,
		},
		{
			"aoc example",
			[][]int{
				[]int{7, 6, 4, 2, 1},
				[]int{1, 2, 7, 8, 9},
				[]int{9, 7, 6, 2, 1},
				[]int{1, 3, 2, 4, 5},
				[]int{8, 6, 4, 4, 1},
				[]int{1, 3, 6, 7, 9},
			},
			4,
		},
	}

	for _, test := range cases2 {
		t.Run(test.Name, func(t *testing.T) {
			got := ReportSafetyChecking(test.Reports, true)
			want := test.NumSafeReports

			if got != want {
				t.Errorf("got %d expected %d given reports %v", got, want, test.Reports)
			}
		})
	}
}
