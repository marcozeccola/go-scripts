package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileToPrint := os.Args[1]
	file, _ := os.Open(fileToPrint)
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	file.Close()
}
