package pev_test

import (
	"testing"

	"github.com/pchchv/pev"
)

func TestRemoveMoreThanTwoFromSequence(t *testing.T) {
	actual := pev.removeMoreThanTwoFromSequence("12345678", "0123456789")
	expected := "12"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.removeMoreThanTwoFromSequence("abcqwertyabc", "qwertyuiop")
	expected = "abcqwabc"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.removeMoreThanTwoFromSequence("", "")
	expected = ""
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.removeMoreThanTwoFromSequence("", "12345")
	expected = ""
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}

func TestGetReversedString(t *testing.T) {
	actual := pev.getReversedString("abcd")
	expected := "dcba"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.getReversedString("1234")
	expected = "4321"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}

func TestRemoveRepeatingChars(t *testing.T) {
	actual := pev.removeMoreThanTwoRepeatingChars("aaaa")
	expected := "aa"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.removeMoreThanTwoRepeatingChars("bbbbbbbaaaaaaaaa")
	expected = "bbaa"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.removeMoreThanTwoRepeatingChars("ab")
	expected = "ab"
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.removeMoreThanTwoRepeatingChars("")
	expected = ""
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}

func TestGetLength(t *testing.T) {
	actual := pev.getLength("aaaa")
	expected := 2
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.getLength("11112222")
	expected = 4
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.getLength("aa123456")
	expected = 4
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.getLength("876543")
	expected = 2
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}

	actual = pev.getLength("qwerty123456z")
	expected = 5
	if actual != expected {
		t.Errorf("Wanted %v, got %v", expected, actual)
	}
}
