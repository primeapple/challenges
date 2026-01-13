package hungryBears

import (
	"cassidoo/utils"
	"testing"
)

func TestHungryBearsAI(t *testing.T) {
	t.Run("should work with example", func(t *testing.T) {
		bears := []Bear{
			{name: "Baloo", hunger: 6},
			{name: "Yogi", hunger: 9},
			{name: "Paddington", hunger: 4},
			{name: "Winnie", hunger: 10},
			{name: "Chicago", hunger: 20},
		}
		got := HungryBearsAI(bears)
		want := []string{"Chicago", "Winnie"}
		utils.AssertDeepEquals(t, got, want)
	})
}
