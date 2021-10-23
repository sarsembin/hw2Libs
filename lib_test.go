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

func Test_squareRoot(t *testing.T) {
	got := SquareRoot(25.0)
	want := 5.0
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Result of SquareRoot was incorrect, got: %v, want: %v\n", got, want)
	}
}

func Test_generateUUID(t *testing.T) {
	got := len(GenerateUUID(10))
	want := 10
	if got != want {
		t.Errorf("Result of GenerateUUID was incorrect, got: %v, want: %v\n", got, want)
	}
}
