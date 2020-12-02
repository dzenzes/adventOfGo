package filehandler_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/dmies/adventOfGo/filehandler"
)

func Test_ImportNumberList(t *testing.T) {
	t.Run("Test that int list is transformed as expected", func(t *testing.T) {
		want := []int{1, 2, 3}
		got, err := filehandler.ToNumberList(strings.NewReader("1,2,3"))

		if err != nil {
			t.Errorf("ToNumberList() error = %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ToNumberList: got %v want %v", got, want)
		}
	})
}

func Test_ToNumberPerLineList(t *testing.T) {
	t.Run("Test that list of numbers is transformed as expected", func(t *testing.T) {
		want := []int{1, 2, 3}
		got, err := filehandler.ToNumberPerLineList(strings.NewReader("1\n2\n3"))

		if err != nil {
			t.Errorf("ToNumberPerLineList() error = %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ToNumberPerLineList: got %v want %v", got, want)
		}
	})
}

func Test_ToStringList(t *testing.T) {
	t.Run("Test that contens are transformed to string list", func(t *testing.T) {
		want := []string{"a", "b", "c"}
		got, err := filehandler.ToSringList(strings.NewReader("a\nb\nc"))

		if err != nil {
			t.Errorf("ToSringList() error = %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ToSringList: got %v want %v", got, want)
		}
	})
}
