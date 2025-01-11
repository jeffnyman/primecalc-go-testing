package main

import "testing"

func Test_checkPrime(t *testing.T) {
	// Test 1
	result, outcome := checkPrime(0)

	if result {
		t.Error("condition: 0 | expected false, got true")
	}

	if outcome != "0 is not prime" {
		t.Error("incorrect outcome")
		t.Errorf("expected: 0 is not prime | got: %s", outcome)
	}

	// Test 2
	result, outcome = checkPrime(4)

	if result {
		t.Error("condition: 4 | expected false, got true")
	}

	if outcome != "4 is not prime; divisible by 2" {
		t.Error("incorrect outcome")
		t.Errorf("expected: 4 is not prime; divisible by 2 | got: %s", outcome)
	}

	// Test 3
	result, outcome = checkPrime(7)

	if !result {
		t.Error("condition: 7 | expected true, got false")
	}

	if outcome != "7 is prime" {
		t.Error("incorrect outcome")
		t.Errorf("expected: 7 is prime | got: %s", outcome)
	}
}
