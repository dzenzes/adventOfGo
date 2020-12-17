package testhelper

import (
	"reflect"
	"testing"
)

func AssertEquals(got interface{}, want interface{}, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}
func AssertDeepEquals(got interface{}, want interface{}, t *testing.T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but wanted %v", got, want)
	}
}
func AssertNoError(err error, t *testing.T) {
	t.Helper()
	if err != nil {
		t.Errorf("method returned error '%v'", err)
	}
}
