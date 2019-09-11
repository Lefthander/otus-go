// Sergey Olisov (c)
// Lesson 2 String Unpacker

// TODO: Fix all linter complains
// TODO: Add parsing escaped chars
package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	ErrUnableToParseString         = errors.New("Unable to parse the inmput string. Incorrect chars sequence - two trailing digits")
	ErrEmptyString                 = errors.New("Input string is empty. Nothing to parse")
	WrnNothingToParse              = errors.New("Input string has no packed chars. Nothing to parse")
	ErrUnableToConvertRuneToNumber = errors.New("Can not convert Rune to number")
	ErrZeroNumberDetected          = errors.New("Warning Zero number detected. Unable to unpack. Skipping...")
)

// isNumberRune check does the rune in range of 0..9
// if yes return the numeric representation of the rune and true
// if no return the 0 and false
func isNumberRune(r rune) (int, bool) {
	number := 0
	flag := false
	var err error

	if string(r) >= "0" && string(r) <= "9" {
		if number, err = strconv.Atoi(string(r)); err != nil {
			log.Println(ErrUnableToConvertRuneToNumber, r)
			flag = false
			return number, flag
		}
		if number == 0 {
			flag = false
			log.Println(ErrZeroNumberDetected)
			return number, flag
		}
		flag = true
	}
	return number, flag
}

// StrUnpack does unpacking the string with following pattern - {chart}{number of repeats}...
// In case of empty string the error should be returned.
// In case of only numbers string the error should be returned.
// 2Do Add parsing escape characters.
func StrUnpack(s string) (string, error) {
	var r rune
	result := ""
	if len(s) <= 1 {
		return s, ErrEmptyString // Empty string detected.
	}
	for _, v := range s {
		fmt.Println("r=", string(r), "v=", string(v))
		if number, flag := isNumberRune(v); flag == true {
			if _, rflag := isNumberRune(r); rflag == true { // Check the previous charcter in order to detect trailing digits.
				return s, ErrUnableToParseString
			}
			result = result + strings.Repeat(string(r), number-1)
			r = v
			continue
		}
		r = v
		result = result + string(v)
	}
	return result, nil
}
func main() {
	teststring := `qwe\45`
	fmt.Println(StrUnpack(teststring))
}
