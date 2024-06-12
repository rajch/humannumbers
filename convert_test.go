package humannumbers_test

import (
	"testing"

	"github.com/rajch/humannumbers"
)

func TestOutOfRange(t *testing.T) {
	_, err := humannumbers.NumToWords(-1)
	if err == nil {
		t.Fatal("out of range value did not return error")
	}
}

func numbertest(t *testing.T, numberToTest int, expectedresult string) {
	result, err := humannumbers.NumToWords(numberToTest)
	if err != nil {
		t.Fatalf("valid value %v resulted in an error: %v", numberToTest, err)
	}
	if result != expectedresult {
		t.Fatalf("expected %v, got %v", expectedresult, result)
	}
}

func TestZero(t *testing.T) {
	numbertest(t, 0, "zero")
}

func TestUnit(t *testing.T) {
	numbertest(t, 1, "one")
	numbertest(t, 13, "thirteen")
}

func TestTens(t *testing.T) {
	numbertest(t, 20, "twenty")
	numbertest(t, 33, "thirty three")
	numbertest(t, 56, "fifty six")
}

func TestHundred(t *testing.T) {
	numbertest(t, 100, "hundred")
}
