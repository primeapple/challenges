package main

import (
	"reflect"
	"testing"
)

func TestRulesetCompare(t *testing.T) {
	ruleset := Ruleset{
		1: &Comparator{[]int{}, []int{2}},
		2: &Comparator{[]int{1}, []int{3}},
		3: &Comparator{[]int{2}, []int{}},
	}

	cases := []struct {
		Name   string
		Num1   int
		Num2   int
		Result int
	}{
		{
			"non existing",
			4,
			5,
			0,
		},
		{
			"same",
			3,
			3,
			0,
		},
		{
			"smaller",
			1,
			2,
			-1,
		},
		{
			"bigger",
			2,
			1,
			1,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := ruleset.Compare(test.Num1, test.Num2)
			want := test.Result

			if got != want {
				t.Errorf("got %d expected %d", got, want)
			}
		})
	}
}

func TestCreateRuleset(t *testing.T) {
	t.Run("adding multiple bigger to single comparator", func(t *testing.T) {
		got := CreateRuleset([]string{
			"1|2",
			"1|3",
		})
		want := Ruleset{
			1: &Comparator{[]int{}, []int{2, 3}},
			2: &Comparator{[]int{1}, []int{}},
			3: &Comparator{[]int{1}, []int{}},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v expected %v", got, want)
		}
	})

	t.Run("basic", func(t *testing.T) {
		got := CreateRuleset([]string{
			"1|2",
			"2|3",
		})
		want := Ruleset{
			1: &Comparator{[]int{}, []int{2}},
			2: &Comparator{[]int{1}, []int{3}},
			3: &Comparator{[]int{2}, []int{}},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v expected %v", got, want)
		}
	})
}

func TestSumMiddlePageOfOrdered(t *testing.T) {
	ruleset := Ruleset{
		1: &Comparator{[]int{}, []int{2}},
		2: &Comparator{[]int{1}, []int{3}},
		3: &Comparator{[]int{2}, []int{}},
	}

	cases := []struct {
		Name   string
		Lines  []string
		Result int
	}{
		{
			"correct lines",
			[]string{
				"1,2,3",
				"1,3,4",
				"1",
				"0",
			},
			6,
		},
		{
			"wrong lines",
			[]string{
				"3,1,2",
				"2,1,5,2,3",
				"0,1,2,3,1",
			},
			0,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := SumMiddlePageOfOrdered(test.Lines, ruleset)
			want := test.Result

			if got != want {
				t.Errorf("got %d expected %d given lines %v", got, want, test.Lines)
			}
		})
	}
}

func TestSumMiddlePageOfUnordered(t *testing.T) {
	ruleset := Ruleset{
		1: &Comparator{[]int{}, []int{2}},
		2: &Comparator{[]int{1}, []int{3}},
		3: &Comparator{[]int{2}, []int{}},
	}

	cases := []struct {
		Name   string
		Lines  []string
		Result int
	}{
		{
			"ordered lines",
			[]string{
				"1,2,3",
				"1,3,4",
				"1",
				"0",
			},
			0,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := SumMiddlePageOfUnordered(test.Lines, ruleset)
			want := test.Result

			if got != want {
				t.Errorf("got %d expected %d given lines %v", got, want, test.Lines)
			}
		})
	}
}
