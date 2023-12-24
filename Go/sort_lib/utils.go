package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func readNumbersFromFile(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // Set the scanner to split by words

	numbers := []int{}

	for scanner.Scan() {
		word := scanner.Text()
		num, err := strconv.Atoi(word)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

func writeNumbersToFile(titleName string, numbers []int) error {
	var fileName string
	if strings.HasSuffix(titleName, ".log") {
		fileName = titleName
	} else {
		fileName = fmt.Sprintf("%s.log", titleName)
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	var printString string = fmt.Sprintf("%s (%d):\n", titleName, len(numbers))
	file.WriteString(printString)
	for _, num := range numbers {
		_, err := file.WriteString(strconv.Itoa(num) + " ")
		if err != nil {
			return err
		}
	}

	return nil
}

func readNumbersFromStdin() ([]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords) // Set the scanner to split by words

	numbers := []int{}

	for scanner.Scan() {
		word := scanner.Text()
		num, err := strconv.Atoi(word)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

func writeNumbersToStdout(numbers []int, algoName string) error {
	var printString string = fmt.Sprintf("%s (%d):\n", algoName, len(numbers))
	os.Stdout.WriteString(printString)
	for _, num := range numbers {
		_, err := os.Stdout.WriteString(strconv.Itoa(num) + " ")
		if err != nil {
			return err
		}
	}

	return nil
}

func generateRandomNumbers(count int, max int) []int {
	numbers := []int{}
	for i := 0; i < count; i++ {
		numbers = append(numbers, rand.Intn(max)+1)
	}
	return numbers
}
