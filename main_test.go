package main

import (
	"bufio"
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

func Test_getNumber(t *testing.T) {
	tests := []struct {
		condition string
		data      string
		expected  string
	}{
		{condition: "quit", data: "q", expected: ""},
		{condition: "QUIT", data: "Q", expected: ""},
		{condition: "empty", data: "", expected: "Enter a whole number"},
		{condition: "text", data: "three", expected: "Enter a whole number"},
		{condition: "decimal", data: "1.3", expected: "Enter a whole number"},
		{condition: "negative", data: "-1", expected: "-1 is not prime"},
		{condition: "zero", data: "0", expected: "0 is not prime"},
		{condition: "one", data: "1", expected: "1 is not prime"},
		{condition: "four", data: "4", expected: "4 is not prime; divisible by 2"},
		{condition: "nine", data: "9", expected: "9 is not prime; divisible by 3"},
		{condition: "seven", data: "7", expected: "7 is prime"},
	}

	for _, e := range tests {
		input := strings.NewReader(e.data)
		reader := bufio.NewScanner(input)

		res, _ := getNumber(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s; got %s", e.condition, e.expected, res)
		}
	}
}

func Test_getInput(t *testing.T) {
	exitChannel := make(chan bool)

	var stdin bytes.Buffer
	var stdout bytes.Buffer

	stdin.Write([]byte("7\nq\n"))

	go getInput(&stdin, &stdout, exitChannel)

	<-exitChannel

	close(exitChannel)
}

func Test_run(t *testing.T) {
	var stdin bytes.Buffer
	var stdout bytes.Buffer

	stdin.WriteString("7\nq\n")

	run(&stdin, &stdout)

	output := stdout.String()

	expectedLines := []string{
		"Is number prime?",
		"Enter a whole number; q to quit.",
		"7 is prime",
		"Exiting",
	}

	for _, line := range expectedLines {
		if !strings.Contains(output, line) {
			t.Errorf("missing expected output: %q; got '%s'", line, output)
		}
	}
}
