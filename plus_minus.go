package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	_, _, _ = reader.ReadLine() // number of inputs
	var strData, _, _ = reader.ReadLine()
	var splitData = strings.Split(string(strData), " ")
	var positive_count = 0
	var zero_count = 0
	var negative_count = 0
	for _, item := range splitData {
		var number, _ = strconv.Atoi(item)
		if number > 0 {
			positive_count += 1
		} else if number < 0 {
			negative_count += 1
		} else {
			zero_count += 1
		}
	}
	var total = float64(positive_count + zero_count + negative_count)
	fmt.Println(float64(positive_count) / total)
	fmt.Println(float64(negative_count) / total)
	fmt.Println(float64(zero_count) / total)
}
