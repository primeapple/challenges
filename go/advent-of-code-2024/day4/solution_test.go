package main

import (
	"testing"
)

func TestXmasCount(t *testing.T) {
	casesSolution1 := []struct {
		Name   string
		Text   [][]byte
		Result int
	}{
		{
			"in line",
			[][]byte{
				[]byte("abcXMASdefSAMX"),
			},
			2,
		},
		{
			"in column",
			[][]byte{
				[]byte("a"),
				[]byte("b"),
				[]byte("c"),
				[]byte("X"),
				[]byte("M"),
				[]byte("A"),
				[]byte("S"),
				[]byte("d"),
				[]byte("e"),
				[]byte("f"),
				[]byte("S"),
				[]byte("A"),
				[]byte("M"),
				[]byte("X"),
				[]byte("h"),
				[]byte("i"),
				[]byte("j"),
			},
			2,
		},
		{
			"diagonal",
			[][]byte{
				[]byte("XS.....SX"),
				[]byte(".MA...AM."),
				[]byte("..AM.MA.."),
				[]byte("...SXS..."),
			},
			4,
		},
		{
			"aoc example1",
			[][]byte{
				[]byte("..X..."),
				[]byte(".SAMX."),
				[]byte(".A..A."),
				[]byte("XMAS.S"),
				[]byte(".X...."),
			},
			4,
		},
	}

	for _, test := range casesSolution1 {
		t.Run(test.Name, func(t *testing.T) {
			got := XmasCount(test.Text)
			want := test.Result

			if got != want {
				t.Errorf("got %d expected %d given text %q", got, want, test.Text)
			}
		})
	}
}

func TestXMasCount(t *testing.T) {
	casesSolution2 := []struct {
		Name   string
		Text   [][]byte
		Result int
	}{
		{
			"aoc example2",
			[][]byte{
				[]byte("M.S"),
				[]byte(".A."),
				[]byte("M.S"),
			},
			1,
		},
		{
			"bigger example",
			[][]byte{
				[]byte(".M.S......"),
				[]byte("..A..MSMS."),
				[]byte(".M.S.MAA.."),
				[]byte("..A.ASMSM."),
				[]byte(".M.S.M...."),
				[]byte(".........."),
				[]byte("S.S.S.S.S."),
				[]byte(".A.A.A.A.."),
				[]byte("M.M.M.M.M."),
				[]byte(".........."),
			},
			9,
		},
	}

	for _, test := range casesSolution2 {
		t.Run(test.Name, func(t *testing.T) {
			got := XMasCount(test.Text)
			want := test.Result

			if got != want {
				t.Errorf("got %d expected %d given text %q", got, want, test.Text)
			}
		})
	}
}
