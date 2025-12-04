package utils

import (
	"reflect"
	"testing"
)

func AssertDeepEquals(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v expected %v", got, want)
	}
}

func AssertIsPositive(t *testing.T, got int) {
	t.Helper()
	if got <= 0 {
		t.Errorf("got %d but wanted positive value", got)
	}
}

func AssertIsNegative(t *testing.T, got int) {
	t.Helper()
	if got >= 0 {
		t.Errorf("got %d but wanted negative value", got)
	}
}
