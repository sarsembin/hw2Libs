package hw2Libs

import (
	"reflect"
	"testing"
)

func Test_registerFlip(t *testing.T) {
	got := RegisterFlip("AbObUs")
	want := "aBoBuS"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Result of registerFlip was incorrect, got: %v, want: %v\n", got, want)
	}
}
