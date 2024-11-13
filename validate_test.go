package pev_test

import (
	"testing"

	"github.com/pchchv/pev"
)

func TestValidate(t *testing.T) {
	expectedError := "insecure password, try including more special characters, using uppercase letters, using numbers or using a longer password"
	err := pev.Validate("mypass", 50)
	if err.Error() != expectedError {
		t.Errorf("Wanted %v, got %v", expectedError, err)
	}

	expectedError = "insecure password, try including more special characters, using lowercase letters, using numbers or using a longer password"
	if err = pev.Validate("MYPASS", 50); err.Error() != expectedError {
		t.Errorf("Wanted %v, got %v", expectedError, err)
	}

	if err = pev.Validate("mypassword", 4); err != nil {
		t.Errorf("Err should be nil")
	}

	if err = pev.Validate("aGoo0dMi#oFChaR2", 80); err != nil {
		t.Errorf("Err should be nil")
	}

	expectedError = "insecure password, try including more special characters, using lowercase letters, using uppercase letters or using a longer password"
	if err = pev.Validate("123", 60); err.Error() != expectedError {
		t.Errorf("Wanted %v, got %v", expectedError, err)
	}
}
