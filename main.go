package main

import "fmt"

func main() {
	data := 2

	_, outcome := checkPrime(data)

	fmt.Println(outcome)
}

func checkPrime(data int) (bool, string) {
	if data == 0 || data == 1 {
		return false, fmt.Sprintf("%d is not prime", data)
	}

	if data < 0 {
		return false, "negative numbers are not prime"
	}

	for i := 2; i <= data/2; i++ {
		if data%i == 0 {
			return false, fmt.Sprintf("%d is not prime; divisible by %d", data, i)
		}
	}

	return true, fmt.Sprintf("%d is prime", data)
}
