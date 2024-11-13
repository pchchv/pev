package pev_test

import (
	"math"
	"testing"

	"github.com/pchchv/pev"
)

func TestLogPow(t *testing.T) {
	expected := math.Round(math.Log2(math.Pow(7, 8)))
	actual := math.Round(pev.logPow(7, 8, 2))
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	expected = math.Round(math.Log2(math.Pow(10, 11)))
	actual = math.Round(pev.logPow(10, 11, 2))
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	expected = math.Round(math.Log2(math.Pow(11, 17)))
	actual = math.Round(pev.logPow(11, 17, 2))
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	expected = math.Round(math.Log10(math.Pow(13, 21)))
	actual = math.Round(pev.logPow(13, 21, 10))
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
