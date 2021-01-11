package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	fmt.Println("")

	for line, num := range counts {
		if num != 1 {
			fmt.Printf("%s\n", line)
		}
	}
}
