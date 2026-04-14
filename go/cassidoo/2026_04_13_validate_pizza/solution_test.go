package validatePizza

import (
	"cassidoo/utils"
	"testing"
)

func TestValidatePizza(t *testing.T) {
	layers := []string{"dough", "sauce", "cheese", "pepperoni", "basil"}
	t.Run("should work with example 1", func(t *testing.T) {
		rules := [][]string{{"sauce", "cheese"}, {"cheese", "pepperoni"}, {"dough", "basil"}}
		ok, possibleRule := ValidatePizza(layers, rules)
		want := true
		utils.AssertEquals(t, ok, want)
		utils.AssertNil(t, possibleRule)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		rules := [][]string{{"cheese", "pepperoni"}, {"cheese", "sauce"}}
		ok, possibleRule := ValidatePizza(layers, rules)
		want := false
		utils.AssertEquals(t, ok, want)
		utils.AssertDeepEquals(t, possibleRule, []string{"cheese", "sauce"})
	})
}
