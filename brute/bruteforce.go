package brute

import (
	"bufio"
	"fmt"
	"os"
)

func Run(w string, d string) {
	fmt.Println("Starting dictionary attack")
	f, err := os.Open(w)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}
