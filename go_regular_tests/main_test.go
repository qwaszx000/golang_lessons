package main

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestAdd(t *testing.T) {
	a, b := 1, 5

	result := add(a, b)
	expected := a + b
	if result != expected {
		t.Fatalf("Add1(%d, %d) resulted in %d instead of expected %d", a, b, result, expected)
	}
}

func TestAdd2(t *testing.T) {
	a, b := 1, 5

	result := add2(a, b)
	expected := a + b
	if result != expected {
		t.Fatalf("Add1(%d, %d) resulted in %d instead of expected %d", a, b, result, expected)
	}

}

func TestReverse_str_bad(t *testing.T) {
	var corpus [][2]string = [][2]string{
		{"12345", "54321"},
		{"567", "7651"}, //Intentionally wrong expected
		{"12", "21"},
		{"1", "1"},
	}

	for _, test_data := range corpus {
		test_name := test_data[0]

		t.Run(test_name, func(t2 *testing.T) {
			result := reverse_str_bad(test_data[0])
			if result != test_data[1] {
				t2.Fatalf("reverse_str_bad(\"%s\") resulted in \"%s\" instead of expected \"%s\"", test_data[0], result, test_data[1])
			}
		})
	}
}

// execute `go test -v` to run tests and see each test result
// `go test` to see only failed tests

func BenchmarkReverse_str_bad(b *testing.B) {
	var (
		long_test_str string
		bytes_str     []byte = make([]byte, 1000)
	)
	const STR_LEN int = 1000

	for i := 0; i < STR_LEN; i++ {
		bytes_str[i] = 'A'
	}

	long_test_str = string(bytes_str)

	for b.Loop() {
		reverse_str_bad(long_test_str)
	}
}

// `go test -bench=. -run=BenchmarkReverse_str_bad` to run benchmarks
// -bench=. -- means run all benchmarks
// -run=BenchmarkReverse_str_bad -- is used to run only BenchmarkReverse_str_bad test, because 2 tests are failing intentionally and it prevents benchmarks from running

func Example() {
	fmt.Println("Hello")
	//Output: Hello1
}

//examples are run with simple `go test`

func FuzzReverse_str_bad(f *testing.F) {
	input_seeds := []string{
		"Hello",
		"12345",
		"A",
	}

	for _, input_str := range input_seeds {
		f.Add(input_str)
	}

	//func internal_fuzz() {}
	internal_fuzz := func(t *testing.T, inp string) {
		var reversed string = reverse_str_bad(inp)
		var reversed_back string = reverse_str_bad(reversed)

		t.Logf("Input str: %q == %v\n", inp, []byte(inp))
		t.Logf("Reversed str: %q == %v\n", reversed, []byte(reversed))

		if reversed_back != inp {
			t.Fatal("Double reverse returned different string\n")
		}

		if utf8.ValidString(inp) && !utf8.ValidString(reversed) {
			t.Fatal("Reversed string is invalid utf8 string, but input string was utf8 valid\n")
		}
	}
	f.Fuzz(internal_fuzz)
}

func FuzzReverse_str_good(f *testing.F) {
	input_seeds := []string{
		"Hello",
		"12345",
		"A",
	}

	for _, input_str := range input_seeds {
		f.Add(input_str)
	}

	internal_fuzz := func(t *testing.T, inp string) {
		reversed, err := reverse_str_good(inp)
		if err != nil {
			return
		}

		reversed_back, err := reverse_str_good(reversed)
		if err != nil {
			return
		}

		t.Logf("Input str: %q == %v\n", inp, []byte(inp))
		t.Logf("Reversed str: %q == %v\n", reversed, []byte(reversed))
		t.Logf("Doubl reversed str: %q == %v\n", reversed_back, []byte(reversed_back))

		if reversed_back != inp {
			t.Fatal("Double reverse returned different string\n")
		}

		if utf8.ValidString(inp) && !utf8.ValidString(reversed) {
			t.Fatal("Reversed string is invalid utf8 string, but input string was utf8 valid\n")
		}
	}
	f.Fuzz(internal_fuzz)
}

// `go test -fuzz=FuzzReverse_str_good -v -fuzztime 10s -run=FuzzReverse_str_good` to run fuzzing test
// `go test -fuzz=FuzzReverse_str_bad -v -fuzztime 10s -run=FuzzReverse_str_bad`
