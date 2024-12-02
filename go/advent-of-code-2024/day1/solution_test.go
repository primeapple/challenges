package main

import "testing"

func TestDistanceSumBetweenLists(t *testing.T) {
	cases := []struct {
		Name        string
		List1       []int
		List2       []int
		DistanceSum int
	}{
		{
			"identical lists",
			[]int{1, 2},
			[]int{1, 2},
			0,
		},
		{
			"identical unordered lists",
			[]int{3, 1, 2},
			[]int{1, 2, 3},
			0,
		},
		{
			"different lists",
			[]int{1, 2},
			[]int{1, 3},
			1,
		},
		{
			"different unorderedlists",
			[]int{2, 1},
			[]int{1, 3},
			1,
		},
		{
			"aoc example",
			[]int{3, 4, 2, 1, 3, 3},
			[]int{4, 3, 5, 3, 9, 3},
			11,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got, err := DistanceSumBetweenLists(test.List1, test.List2)
			want := test.DistanceSum

			if err != nil {
				t.Errorf("got unexpected error %q", err.Error())
			}

			if got != want {
				t.Errorf("got %d expected %d given list1 %v list2 %v", got, want, test.List1, test.List2)
			}
		})
	}
}

func TestSimilarityScore(t *testing.T) {
	cases := []struct {
		Name            string
		List1           []int
		List2           []int
		SimilarityScore int
	}{
		{
			"identical lists",
			[]int{1, 2},
			[]int{1, 2},
			3,
		},
		{
			"non intersecting lists",
			[]int{2},
			[]int{1},
			0,
		},
		{
			"intersecting lists",
			[]int{3, 2, 1},
			[]int{3, 3, 3},
			9,
		},
		{
			"aoc example",
			[]int{3, 4, 2, 1, 3, 3},
			[]int{4, 3, 5, 3, 9, 3},
			31,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := SimilarityScore(test.List1, test.List2)
			want := test.SimilarityScore

			if got != want {
				t.Errorf("got %d expected %d given list1 %v list2 %v", got, want, test.List1, test.List2)
			}
		})
	}
}
