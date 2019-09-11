package main

import (
	"testing"
)

func TestStrUnpackBasic(t *testing.T) {
	teststring := "a4bc2d5e"
	expected := "aaaabccddddde"
	result, err := StrUnpack(teststring)

	if err != nil {
		t.Error(err)
	}
	if result != expected {
		t.Errorf("Unpacked string %s does not match the expected string %s", result, expected)
	}
}
func TestStrUnpackNothingToUnpack(t *testing.T) {
	teststring := "abcd"
	expected := "abcd"
	result, err := StrUnpack(teststring)

	if err != nil {
		t.Error(err)
	}
	if result != expected {
		t.Errorf("Unpacked string %s does not match the expected string %s", result, expected)
	}
}

func TestStrUnpackIncorrectInputString(t *testing.T) {
	teststring := "45"
	expectedString := "45"
	expectedError := ErrUnableToParseString
	result, err := StrUnpack(teststring)

	if err != expectedError {
		t.Error(err)
	}
	if result != expectedString {
		t.Errorf("Returned string %s does not match the expected string %s", result, expectedString)
	}

}

func TestStrUnpackWithEscapeChar1(t *testing.T) {
	teststring := `qwe\4\5`
	expectedString := "qwe45"
	result, err := StrUnpack(teststring)

	if err != nil {
		t.Error(err)
	}
	if result != expectedString {
		t.Errorf("Returned string %s does not match the expected string %s", result, expectedString)
	}

}
func TestStrUnpackWithEscapeChar2(t *testing.T) {
	teststring := `qwe\45`
	expectedString := "qwe44444"
	result, err := StrUnpack(teststring)

	if err != nil {
		t.Error(err)
	}
	if result != expectedString {
		t.Errorf("Returned string %s does not match the expected string %s", result, expectedString)
	}

}

func TestStrUnpackWithEscapeChar3(t *testing.T) {
	teststring := `qwe\\5`
	expectedString := `qwe\\\\\`
	result, err := StrUnpack(teststring)

	if err != nil {
		t.Error(err)
	}
	if result != expectedString {
		t.Errorf("Returned string %s does not match the expected string %s", result, expectedString)
	}

}
