package main

import "testing"

func Test_checkPrime(t *testing.T) {
	primeTests := []struct {
		condition string
		data      int
		result    bool
		outcome   string
	}{
		{"0 is not prime", 0, false, "0 is not prime"},
		{"4 is not prime", 4, false, "4 is not prime; divisible by 2"},
		{"7 is prime", 7, true, "7 is prime"},
	}

	for _, test := range primeTests {
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
	}
}
