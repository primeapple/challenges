package maxMealPrepTasks

import (
	"testing"
)

func TestMaxMealPrepTasks(t *testing.T) {
	// EXAMPLE
	t.Run("should work with example", func(t *testing.T) {
		tasks := []Task{
			{name: "Make Gravy", start: 10, end: 11},
			{name: "Mash Potatoes", start: 11, end: 12},
			{name: "Bake Rolls", start: 11, end: 13},
			{name: "Prep Salad", start: 12, end: 13},
		}

		got := MaxMealPrepTasks(tasks)
		want := TasksOrder{
			count:  3,
			chosen: []string{"Make Gravy", "Mash Potatoes", "Prep Salad"},
		}

		assertDeepEquals(t, got, want)
	})
}
