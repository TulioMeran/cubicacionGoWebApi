package utils

import (
	"testing"
)

func TestPasswordEncriptor(t *testing.T) {

	password, err := PasswordEncriptor("123456")

	if err != nil {
		t.Fatal(err)
	}

	if len(password) < 1 {
		t.Errorf("Expected a hash password but got %v", password)
	}

	if password == "123456" {
		t.Error("Expected to hash the password but it is equal to pass before encriptor.")
	}

}
