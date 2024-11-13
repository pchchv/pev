package pev_test

import (
	"testing"
	"github.com/pchchv/pev"
)

const (
	sepChars          = `_-., `
	replaceChars      = `!@$&*`
	digitsChars       = `0123456789`
	lowerChars        = `abcdefghijklmnopqrstuvwxyz`
	upperChars        = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	otherSpecialChars = `"#%'()+/:;<=>?[\]^{|}~`
)

func TestGetBase(t *testing.T) {
	actual := pev.getBase("abcd")
	expected := len(pev.lowerChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.getBase("abcdA")
	expected = len(lowerChars) + len(upperChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.getBase("A")
	expected = len(upperChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.getBase("^_")
	expected = len(otherSpecialChars) + len(sepChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.getBase("^")
	expected = len(otherSpecialChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.getBase("!")
	expected = len(replaceChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.getBase("123")
	expected = len(digitsChars)
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.getBase("123ü")
	expected = len(digitsChars) + 1
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}
