package konamiCode

import (
	"cassidoo/utils"
	"testing"
)

func TestKonamiCode(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		result, err := KonamiCode("xx2233454590yy11110")
		want := map[string]string{"0": "A", "2": "U", "3": "D", "4": "L", "5": "R", "9": "B"}
		utils.AssertNil(t, err)
		utils.AssertDeepEquals(t, result, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		result, err := KonamiCode("sduwahoda22ii0d0dbn")
		want := map[string]string{"0": "L", "2": "U", "i": "D", "d": "R", "b": "B", "n": "A"}
		utils.AssertNil(t, err)
		utils.AssertDeepEquals(t, result, want)
	})

	t.Run("should return err if no mapping found", func(t *testing.T) {
		_, err := KonamiCode("abcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabc")
		utils.AssertError(t, err, ErrNoMappingFound)
	})
}
