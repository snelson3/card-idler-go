package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
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
	fmt.Printf("You have opened %d packs (worth %d)\n", s.stats.packs_opened, getInventoryValue(s))
	fmt.Printf("You have %d moneys\n", )
	for a := range s.shop {
		selcracker := s.shop[a]
		if selcracker.held > 0 {
			fmt.Printf("Number of %s (Price %f): %d\n", selcracker.name, selcracker.price, selcracker.held)
		}
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
		line := strings.Fields(strings.ToLower(readLine()))
		switch command = line[0]; command {
		case "quit":
			stopGame(&game)
			return
		case "crack":
			fmt.Println("You cracked a pack!")
			crackPacks(&game, 1)
		case "sell":
			fmt.Println("You are selling all your cards")
			sell(&game)
		case "buy":
			tobuy, _ := strconv.Atoi(line[1])
			buy(&game, tobuy)
		}
	}
}

// todo 
// fix runtime errors
// 10 minutes to gather as much money as you can (or crack as many packs as you can)
// high score chart at the end that saves to file
// basic dueling feature not sure what I want to do (dials to alter about how much of each rarity to use per duel etc)
// maybe a few other things to make it more interesting, but not making it a huge task to transfer over
// then make it an actual game with a visual display and sounds and stuff
// can you make it play through a website using WASM?
// "seasons", achievements, upgrades, etc