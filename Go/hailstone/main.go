package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
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
	tmp, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		if string(s) != "" {
			fmt.Println("Error the number provided wasn't accepted:", err)
		}
		os.Exit(2)
	}

	Max := int64(tmp)

	now := time.Now()
	numThreads := runtime.NumCPU()
	bufferSize := numThreads * 100
	running := make(chan int64, bufferSize)
	fmt.Printf("Created a buffered channel with capacity: %d\n", bufferSize)

	var wg sync.WaitGroup

	if allseq {
		fmt.Println("I will print the sequence for the first", Max, "numbers")
		for i := int64(0); i <= Max; i++ {
			wg.Add(1)
			go findSequence(i, running, &wg)
		}
	} else {
		fmt.Println("I will print the sequence for", Max, "only")
		wg.Add(1)
		go findSequence(Max, running, &wg)
	}

	go func() {
		wg.Wait()
		close(running)
	}()

	received := int64(0)
	for range running {
		received++
		if received >= Max {
			break
		}
	}
	println("Received", received, "values and done in", time.Since(now).Milliseconds(), "ms")
}

func findSequence(n int64, running chan int64, wg *sync.WaitGroup) {
	defer wg.Done()

	output := Generate(n)
	if verboseON {
		fmt.Println("The sequence for #", n, "has (", len(output), " elements):\n", output)
	} else {
		fmt.Println("The sequence for #", n, "has (", len(output), " elements).")
	}
	running <- n
}
