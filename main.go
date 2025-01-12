package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	startup()

	exitChannel := make(chan bool)

	go getInput(exitChannel)

	<-exitChannel

	close(exitChannel)

	fmt.Println("Exiting")
}

func getInput(exitChannel chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		result, done := getNumber(scanner)

		if done {
			exitChannel <- true
			return
		}

		fmt.Println(result)
		prompt()
	}
}

func getNumber(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	number, err := strconv.Atoi(scanner.Text())

	if err != nil {
		return "Enter a whole number", false
	}

	_, outcome := checkPrime(number)

	return outcome, false
}

func startup() {
	fmt.Println("Is number prime?")
	fmt.Println("Enter a whole number; q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("> ")
}

func checkPrime(data int) (bool, string) {
	if data <= 1 {
		return false, fmt.Sprintf("%d is not prime", data)
	}

	for i := 2; i*i <= data; i++ {
		if data%i == 0 {
			return false, fmt.Sprintf("%d is not prime; divisible by %d", data, i)
		}
	}

	return true, fmt.Sprintf("%d is prime", data)
}
