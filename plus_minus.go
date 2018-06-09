// 6
// -4 3 -9 0 4 1
// print fraction of positive, negative, and zeros
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	reader.ReadLine()
	fmt.Println("Hello!")
}
