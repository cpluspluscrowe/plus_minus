package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func transformInputToIntegers(input []string) []int {
	var transformed []int
	for _, item := range input {
		var number, _ = strconv.Atoi(item)
		transformed = append(transformed, number)
	}
	return transformed
}

func getSignCounts(numbers []int) (int, int, int) {
	var positive_count = 0
	var zero_count = 0
	var negative_count = 0
	for _, number := range numbers {
		if number > 0 {
			positive_count += 1
		} else if number < 0 {
			negative_count += 1
		} else {
			zero_count += 1
		}
	}
	return positive_count, zero_count, negative_count
}

func printRatios(positive_count int, zero_count int, negative_count int, total float64) {
	fmt.Println(float64(positive_count) / total)
	fmt.Println(float64(negative_count) / total)
	fmt.Println(float64(zero_count) / total)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	_, _, _ = reader.ReadLine() // number of inputs
	var strData, _, _ = reader.ReadLine()
	var splitData = strings.Split(string(strData), " ")
	var numbers = transformInputToIntegers(splitData)
	var positive_count, zero_count, negative_count int = getSignCounts(numbers)
	var total = float64(positive_count + zero_count + negative_count)
	printRatios(positive_count, zero_count, negative_count, total)
}
