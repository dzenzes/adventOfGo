package filehandler_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/dmies/adventOfGo/filehandler"
)

func Test_ImportNumberList(t *testing.T) {
	t.Run("Test something", func(t *testing.T) {
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

func Test_ImportNumberPerLineList(t *testing.T) {
	t.Run("Test something", func(t *testing.T) {
		want := []int{1, 2, 3}
		got, err := filehandler.ImportNumberPerLineList(strings.NewReader("1\n2\n3"))

		if err != nil {
			t.Errorf("ImportNumberPerLineList() error = %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ImportNumberPerLineList: got %v want %v", got, want)
		}
	})
}
