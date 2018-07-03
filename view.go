package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readLine() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	if scanner.Err() != nil {
		fmt.Println("Something went wrong with that line")
		log.Fatal(scanner.Err())
	}
	return scanner.Text()
}

func displayInfo(s State) {
	fmt.Println("###########################")
	fmt.Printf("You have opened %d packs (worth %d)\n", s.stats.packsOpened, getInventoryValue(s))
	if s.shop[0].held > 0 {
		fmt.Printf("Number of %s: %d\n", s.shop[0].name, s.shop[0].held)
	}
	fmt.Println("###########################")
}

func main() {
	fmt.Println("Welcome to cardsim")
	game := initGame()
	var command string
	for game.running {
		update(&game)
		displayInfo(game)
		fmt.Println("What would you like to do?")
		switch command = strings.ToLower(readLine()); command {
		case "quit":
			stopGame(&game)
			return
		case "crack":
			fmt.Println("You cracked a pack!")
			crackPacks(&game, 1)
		case "buy":
			fmt.Println(buy(&game))
		}
	}
}
