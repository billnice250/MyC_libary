package main

import (
	"flag"
	"fmt"
	"sort_lib/algorithms"
	"sync"
	"time"
)

var (
	verbose       bool
	quick         bool
	merge         bool
	sleep         bool
	all           bool
	help          bool
	file          string
	input         string
	save          bool
	randomize     int
	randomize_max int
)

const RANDOM_MAX = 1000

type AlgoOutput struct {
	Duration time.Duration
	AlgoName string
	Numbers  []int
}

func main() {
	// read the flags
	// if -quick then run quicksort
	// if -merge then run mergesort
	// if -sleep then run sleepsort
	// if -all then run all
	// if -help then print help
	// if no flags then print help
	// if --file then read the file and sort the numbers in the specified file
	// if --input then read the input and sort the numbers
	// if --output then write the output to the file, if no file then print to stdout
	helpString := `Usage: sort_lib [options] [file] [input] [output]`
	flag.BoolVar(&help, "help", false, helpString)
	if help {
		flag.PrintDefaults()
		return
	}
	flag.BoolVar(&quick, "quick", false, "run quicksort")
	flag.BoolVar(&merge, "merge", false, "run mergesort")
	flag.BoolVar(&sleep, "sleep", false, "run sleepsort")
	flag.BoolVar(&all, "all", false, "run all")
	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.StringVar(&file, "file", "", "read the file and sort the numbers in the specified file")
	flag.StringVar(&input, "input", "", "read the input and sort the numbers")
	flag.IntVar(&randomize, "random", 0, "generate N random numbers ")
	flag.IntVar(&randomize_max, "random_celing", RANDOM_MAX, "randomize the numbers")
	flag.BoolVar(&save, "save", false, "write the output to the file, if no file then print to stdout")
	flag.Parse()
	var numbers []int
	if file != "" {
		numbers, _ = readNumbersFromFile(file)
	} else if input != "" {
		numbers, _ = readNumbersFromStdin()
	} else if randomize > 0 {
		numbers = generateRandomNumbers(randomize, randomize_max)
	}
	if len(numbers) == 0 {
		fmt.Println("No numbers to sort")
		flag.PrintDefaults()
		return
	}

	var algoOutputs chan AlgoOutput = make(chan AlgoOutput, 3)
	wg := sync.WaitGroup{}

	if quick || all {
		wg.Add(1)
		go func() {
			start := time.Now()
			newNumbers := make([]int, len(numbers))
			copy(newNumbers, numbers)
			algorithms.QuickSort(newNumbers)
			duration := time.Since(start)
			algoOutputs <- AlgoOutput{duration, "QuickSort", newNumbers}
			wg.Done()
		}()
	}
	if merge || all {
		wg.Add(1)
		go func() {
			start := time.Now()
			newNumbers := make([]int, len(numbers))
			copy(newNumbers, numbers)
			algorithms.MergeSort(newNumbers)
			duration := time.Since(start)
			algoOutputs <- AlgoOutput{duration, "MergeSort", newNumbers}
			wg.Done()
		}()
	}
	if sleep || all {
		wg.Add(1)
		go func() {
			start := time.Now()
			newNumbers := make([]int, len(numbers))
			copy(newNumbers, numbers)
			algorithms.SleepSort(&newNumbers)
			duration := time.Since(start)
			algoOutputs <- AlgoOutput{duration, "SleepSort", newNumbers}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(algoOutputs)
	}()
	fmt.Println("Sorting ", len(numbers), " numbers using ", len(algoOutputs), " algorithms")

	var results []AlgoOutput = make([]AlgoOutput, 3)
	for algoOutput := range algoOutputs {
		fmt.Println(algoOutput.AlgoName, " took ", algoOutput.Duration)
		results = append(results, algoOutput)
		if save {
			writeNumbersToFile(algoOutput.AlgoName, algoOutput.Numbers)
		} else if verbose {
			writeNumbersToStdout(algoOutput.Numbers, algoOutput.AlgoName)
		}
	}
	for _, result := range results {
		fmt.Println(result.AlgoName, " took ", result.Duration, " and sorted ", len(result.Numbers), " numbers")
	}

}
