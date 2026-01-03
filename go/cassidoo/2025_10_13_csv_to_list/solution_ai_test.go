package csvToList

import (
	"cassidoo/utils"
	"testing"
)

func TestCsvToListAI(t *testing.T) {
	t.Run("should work with empty csv", func(t *testing.T) {
		got := CsvToListAI("")
		want := ""
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with csv that only contains header", func(t *testing.T) {
		got := CsvToListAI("name,age,city")
		want := ""
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example", func(t *testing.T) {
		got := CsvToListAI("name,age,city\n\"Ryu, Mi-yeong\",30,\"Seoul\"\nZoey,24,\"Burbank\"")
		want := `
- Ryu, Mi-yeong, age 30, from Seoul
- Zoey, age 24, from Burbank
`
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example with different column order", func(t *testing.T) {
		got := CsvToListAI("age,city,name\n30,\"Seoul\",\"Ryu, Mi-yeong\"\n24,\"Burbank\",Zoey")
		want := `
- Ryu, Mi-yeong, age 30, from Seoul
- Zoey, age 24, from Burbank
`
		utils.AssertEquals(t, got, want)
	})
}
