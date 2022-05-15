package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"billnice.it/libary/hailstone" // import the util package
)

var verboseON bool
var allseq bool
var help bool

func main() {
	fmt.Println("Hello, I'm a program that prints the hailstone sequence for a given number ")
	if len(os.Args[1:]) < 1 {
		fmt.Println("Usage: stone <number>  //default print the count of sequences for the number given")
		//usage options
		fmt.Println("\n\t-v			//verbose mode,\n\t-all		//print all the sequences for numbers below <number>\n\t-h		//print this help")
		os.Exit(1)
	}

	// get the verbose flag
	flag.BoolVar(&verboseON, "v", false, "verbose output")
	flag.BoolVar(&allseq, "all", false, "only print the sequence of that one number")
	flag.BoolVar(&help, "h", false, "help")
	flag.Parse()
	if help {
		fmt.Println("Usage: stone <number>  //default print the sequence for the number given.")
		//usage options
		fmt.Println("\n\t-v			//verbose mode,\n\t-only		//only print the sequence for the number given\n\t-help		//print this help")
		os.Exit(1)

	}

	s := flag.Arg(0)
	// string to int
	tmp, err := strconv.Atoi(s)
	if err != nil {
		if string(s) != "" {
			fmt.Println("Error the number provided wasn't accepted:", err)
		}
		os.Exit(2)
	}

	Max := int64(tmp)

	running := make(chan int64, Max)
	var i int64

	if allseq {
		fmt.Println("I will print the sequence for the first", Max, "numbers")

		for i = 0; i <= int64(Max); i++ {
			go findSequence(i, running)
		}

	} else {
		fmt.Println("I will print the sequence for", Max, "only")

		go findSequence(Max, running)
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

	if !allseq {
		close(running)
	}

}
