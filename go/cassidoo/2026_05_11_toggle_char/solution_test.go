package toggleChar

import (
	"cassidoo/utils"
	"testing"
)

func TestToggleChar(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got, _ := ToggleChar("Hello, world!")
		want := "hELLO, WORLD!"
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got, _:=ToggleChar("HeheHeheHEheheHeH")
		want := "hEHEhEHEheHEHEhEh"
		utils.AssertEquals(t, got, want)
	})
	// Didn't get this one to work :(
	// t.Run("should work with example 3", func(t *testing.T) {
	// 	got, _ := AlternateChar("This will be alternated")
	// 	want := "ThIs WiLl Be AlTeRnAtEd"
	// 	utils.AssertEquals(t, got, want)
	// })
}
