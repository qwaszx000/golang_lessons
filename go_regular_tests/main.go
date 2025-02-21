package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func add(a, b int) int {
	return a * b
}

func add2(a, b int) int {
	return a + b
}

func reverse_str_bad(s string) (result string) {
	bytes_str := []byte(s)

	for first_el, last_el := 0, len(bytes_str)-1; first_el < len(bytes_str)/2; first_el, last_el = first_el+1, last_el-1 {
		bytes_str[first_el], bytes_str[last_el] = bytes_str[last_el], bytes_str[first_el]
	}

	return string(bytes_str)
}

// still bad
// Probably i can add support for broken utf8 encoding if i do it with [][]byte instead of []rune
/*func reverse_str_good(s string) (result string) {

	const (
		EMPTY_STR = 0
		BAD_RUNE  = 1
	)

	utf8_runes_count := utf8.RuneCountInString(s)
	runes_slice := make([]rune, 0, utf8_runes_count)

	//Translate str to utf8 runes
	var offset int = 0
	for {
		r, size := utf8.DecodeRuneInString(s[offset:])

		//Stop if we've reached end of str
		if r == utf8.RuneError && size == EMPTY_STR {
			break
		}

		if r == utf8.RuneError {
			//Add raw 1 byte rune to slice because it is not utf8 char
			runes_slice = append(runes_slice, rune(s[offset]))
		} else {
			runes_slice = append(runes_slice, r)
		}

		//size == 1 when BAD_RUNE, so we can use it as 1 byte offset
		offset += size
	}

	//Reverse utf8 runes order
	for first_el, last_el := 0, len(runes_slice)-1; first_el < len(runes_slice)/2; first_el, last_el = first_el+1, last_el-1 {
		runes_slice[first_el], runes_slice[last_el] = runes_slice[last_el], runes_slice[first_el]
	}

	return string(runes_slice)
}*/

// And there is simple solution
func reverse_str_good(s string) (result string, err error) {
	if !utf8.ValidString(s) {
		return "", errors.New("invalid utf8 string")
	}
	runes_slice := []rune(s)

	for first_el, last_el := 0, len(runes_slice)-1; first_el < len(runes_slice)/2; first_el, last_el = first_el+1, last_el-1 {
		runes_slice[first_el], runes_slice[last_el] = runes_slice[last_el], runes_slice[first_el]
	}

	return string(runes_slice), nil
}

func main() {
	fmt.Println("Start")
}
