package minassemblytime

import "testing"

func TestMinAssemblyTime(t *testing.T) {
	t.Run("should work without arrival days", func(t *testing.T) {
		parts := []AssemblyDetail{
			{
				name:          "keycaps",
				arrivalDays:   0,
				assemblyHours: 2,
			},
			{
				name:          "switches",
				arrivalDays:   0,
				assemblyHours: 2,
			},
		}

		got := MinAssemblyTime(parts)
		want := 4

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("long arrival date should rule over other parts", func(t *testing.T) {
		parts := []AssemblyDetail{
			{
				name:          "keycaps",
				arrivalDays:   0,
				assemblyHours: 2,
			},
			{
				name:          "switches",
				arrivalDays:   5,
				assemblyHours: 2,
			},
			{
				name:          "stabilizers",
				arrivalDays:   0,
				assemblyHours: 2,
			},
		}

		got := MinAssemblyTime(parts)
		want := 122

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("should work with example", func(t *testing.T) {
		parts := []AssemblyDetail{
			{
				name:          "keycaps",
				arrivalDays:   1,
				assemblyHours: 2,
			},
			{
				name:          "switches",
				arrivalDays:   2,
				assemblyHours: 3,
			},
			{
				name:          "stabilizers",
				arrivalDays:   0,
				assemblyHours: 1,
			},
			{
				name:          "PCB",
				arrivalDays:   1,
				assemblyHours: 4,
			},
			{
				name:          "case",
				arrivalDays:   3,
				assemblyHours: 2,
			},
		}

		got := MinAssemblyTime(parts)
		want := 74

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
