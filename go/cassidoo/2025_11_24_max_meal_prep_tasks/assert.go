package maxMealPrepTasks

import (
	"reflect"
	"testing"
)

func assertDeepEquals(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v expected %v", got, want)
	}
}
