package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func readLine(outline string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(outline)
	scanner.Scan()
	if scanner.Err() != nil {
		// handle error
	}
	return scanner.Text()
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 3; i++ {
		x, _ := strconv.Atoi(readLine("How many cards are in your hand?"))
		fmt.Println("Discard: ")
		for j := 0; j < 3; j++ {
			fmt.Println(rand.Intn(x - j))
		}
	}
}
