package resolvePath

import (
	"cassidoo/utils"
	"testing"
)

func sp(s string) *string {
	return &s
}

func TestResolvePath(t *testing.T) {
	fs := map[string]*string{
		"/a":     sp("/b"),
		"/b":     sp("/c"),
		"/c":     nil,
		"/loop1": sp("/loop2"),
		"/loop2": sp("/loop1"),
		"/real":  nil,
		"/alias": sp("/real"),
	}

	t.Run("should work non existing paths", func(t *testing.T) {
		got, err := ResolvePath(fs, "/abcdefghijklmnopqrstuvwxyz")
		utils.AssertError(t, err, ErrNoPathFound)
		utils.AssertEquals(t, got, "")
	})

	t.Run("should work with example 1", func(t *testing.T) {
		got, err := ResolvePath(fs, "/a")
		utils.AssertNil(t, err)
		utils.AssertEquals(t, got, "/c")
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got, err := ResolvePath(fs, "/alias")
		utils.AssertNil(t, err)
		utils.AssertEquals(t, got, "/real")
	})

	t.Run("should work with example 3", func(t *testing.T) {
		got, err := ResolvePath(fs, "/loop1")
		utils.AssertError(t, err, ErrCycleFound)
		utils.AssertEquals(t, got, "")
	})

	t.Run("should work with example 4", func(t *testing.T) {
		got, err := ResolvePath(fs, "/real")
		utils.AssertNil(t, err)
		utils.AssertEquals(t, got, "/real")
	})
}
