package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func Test_checkPrime(t *testing.T) {
	primeTests := []struct {
		condition string
		data      int
		result    bool
		outcome   string
	}{
		{"not prime", -1, false, "-1 is not prime"},
		{"not prime", 0, false, "0 is not prime"},
		{"not prime", 1, false, "1 is not prime"},
		{"not prime", 4, false, "4 is not prime; divisible by 2"},
		{"not prime", 9, false, "9 is not prime; divisible by 3"},
		{"is prime", 7, true, "7 is prime"},
	}

	for _, test := range primeTests {
		testName := fmt.Sprintf("%s_%d", test.condition, test.data)

		t.Run(testName, func(t *testing.T) {
			result, outcome := checkPrime(test.data)

			if test.result && !result {
				t.Errorf("FAIL: %s - checkPrime(%d) was %t; expected %t",
					test.condition, test.data, result, test.result)
			}

			if !test.result && result {
				t.Errorf("FAIL: %s - checkPrime(%d) was %t; expected %t",
					test.condition, test.data, result, test.result)
			}

			if test.outcome != outcome {
				t.Errorf("FAIL: %s - checkPrime(%d) message was %s; expected %s",
					test.condition, test.data, outcome, test.outcome)
			}
		})
	}
}

func Test_prompt(t *testing.T) {
	var buf bytes.Buffer

	prompt(&buf)

	if buf.String() != "> " {
		t.Errorf("incorrect prompt; expected '>' | got '%s'", buf.String())
	}
}

func Test_startup(t *testing.T) {
	var buf bytes.Buffer

	startup(&buf)

	output := buf.String()

	expectedLines := []string{
		"Is number prime?",
		"Enter a whole number; q to quit.",
	}

	for _, line := range expectedLines {
		if !strings.Contains(output, line) {
			t.Errorf("missing expected output: %q; got '%s'", line, output)
		}
	}
}
