package main

import "testing"

func TestGetCharAfterVimCommands(t *testing.T) {
	text := `Hello, world!
how are ya?`

	t.Run("should work with h command", func(t *testing.T) {
		got := GetCharAfterVimCommands(text, "h")
		want := "H"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should work with j command", func(t *testing.T) {
		got := GetCharAfterVimCommands(text, "j")
		want := "h"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should work with k command", func(t *testing.T) {
		got := GetCharAfterVimCommands(text, "k")
		want := "H"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should work with l command", func(t *testing.T) {
		got := GetCharAfterVimCommands(text, "l")
		want := "e"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should work with example", func(t *testing.T) {
		text := `Hello, world!
how are ya?`
		commands := "jlhll"

		got := GetCharAfterVimCommands(text, commands)
		want := "w"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should handle moving down past last line", func(t *testing.T) {
		text := "line1\nline2"
		got := GetCharAfterVimCommands(text, "jjj") // Try to go down 3 times from first line
		want := "l"                                 // Should stay at start of last line
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should handle moving right past end of line", func(t *testing.T) {
		text := "Hi\nLonger"
		got := GetCharAfterVimCommands(text, "lllll") // Try to go right 5 times
		want := "i"                                   // Should stay at last char of first line
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should handle moving to end of shorter line after longer line", func(t *testing.T) {
		text := "LongerLine\nShort"
		got := GetCharAfterVimCommands(text, "lllllllllj") // Move to end of long line, down
		want := "t"                                        // Should stay at last char of short line
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
