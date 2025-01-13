package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(input io.Reader, output io.Writer) {
	startup(output)

	exitChannel := make(chan bool)

	go getInput(input, output, exitChannel)

	<-exitChannel

	close(exitChannel)

	fmt.Fprintln(output, "Exiting")
}

func getInput(input io.Reader, output io.Writer, exitChannel chan bool) {
	scanner := bufio.NewScanner(input)

	for {
		result, done := getNumber(scanner)

		if done {
			exitChannel <- true
			return
		}

		fmt.Fprintln(output, result)
		prompt(os.Stdout)
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

func startup(out io.Writer) {
	fmt.Fprintln(out, "Is number prime?")
	fmt.Fprintln(out, "Enter a whole number; q to quit.")
	prompt(out)
}

func prompt(out io.Writer) {
	fmt.Fprint(out, "> ")
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
