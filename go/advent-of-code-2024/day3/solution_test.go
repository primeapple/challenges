package main

import (
	"reflect"
	"testing"
)

func TestParseNumbers(t *testing.T) {
	casesSolution1 := []struct {
		Name   string
		Text   string
		Result []Pair
	}{
		{
			"single mul",
			"mul(1,2)",
			[]Pair{{1, 2}},
		},
		{
			"multiple mul",
			"mul(1,2)mul(1,2),mul(1,2)",
			[]Pair{{1, 2}, {1, 2}, {1, 2}},
		},
	}

	for _, test := range casesSolution1 {
		t.Run(test.Name, func(t *testing.T) {
			got := ParseNumbers(test.Text, false)
			want := test.Result

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v expected %v given text %q", got, want, test.Text)
			}
		})
	}

	casesSolution2 := []struct {
		Name   string
		Text   string
		Result []Pair
	}{
		{
			"single mul with dodont",
			"mul(1,2)",
			[]Pair{{1, 2}},
		},
		{
			"multiple mul with dodont",
			"mul(1,2)mul(1,2),mul(1,2)",
			[]Pair{{1, 2}, {1, 2}, {1, 2}},
		},
		{
			"single dont with dodont",
			"mul(1,2)don't()mul(1,2),mul(1,2)",
			[]Pair{{1, 2}},
		},
		{
			"multiple do/dont with dodont",
			"mul(1,2)don't()mul(1,3)do(),mul(1,4)",
			[]Pair{{1, 2}, {1, 4}},
		},
	}

	for _, test := range casesSolution2 {
		t.Run(test.Name, func(t *testing.T) {
			got := ParseNumbers(test.Text, true)
			want := test.Result

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v expected %v given text %q", got, want, test.Text)
			}
		})
	}
}
