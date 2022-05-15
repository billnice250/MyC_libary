package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"billnice.it/libary/hailstone" // import the util package
)

var verboseON = false

func main() {
	// get the verbose flag
	flag.BoolVar(&verboseON, "v", false, "verbose output")

	flag.Parse()
	s := flag.Arg(0)
	// string to int
	tmp, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Usage: stone <number>\n //for just counts or\n  stone -v <number> //to get the sequences printed")
		fmt.Println(err)
		os.Exit(2)
	}

	Max := int64(tmp)

	fmt.Println("Hello, I'm a program that prints the hailstone sequence for a given number")
	fmt.Println("I will print the sequence for the first", Max, "numbers")

	running := make(chan int64, Max)
	var i int64

	for i = 0; i <= int64(Max); i++ {
		go findSequence(i, running)
	}

	received := int64(0)
	for range running {
		received++
		if received >= Max {
			break

		}

	}
	println("Received", received, "values")

}

func findSequence(n int64, running chan int64) {

	output := hailstone.Generate(n)
	if verboseON {
		fmt.Println("The sequence for #", n, "has (", len(output), " elements):\n", output)
	} else {
		fmt.Println("The sequence for #", n, "has (", len(output), " elements).")

	}

	running <- n

}
