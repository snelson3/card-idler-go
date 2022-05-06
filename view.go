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
	fmt.Printf("You have cracked %d packs\n", s.stats.packs_opened)
	fmt.Printf("You are cracking %d packs per second\n", getPacksPerSecond(&s))
	fmt.Printf("Card Collection:\n")
	for r := range s.inventory.card_types {
		selrarity := s.inventory.card_types[r]
		val := float64(selrarity.held) * selrarity.value
		fmt.Printf("\t%d %s (Worth %f)\n", selrarity.held, selrarity.name, val)
	}
	fmt.Printf("You have %f moneys\n", s.inventory.cash)
	fmt.Printf("Shop\n")
	for a := range s.shop {
		selcracker := s.shop[a]
		fmt.Printf("\tNumber of %s (Price %f): %d\n", selcracker.name, selcracker.price, selcracker.held)
	}
	fmt.Println("###########################")
}

func commandLineAction(game *State) {
	fmt.Println(`What would you like to do?\n
	crack // manually open a pack 
	sell // sell all of your cards for money
	buy x // buy the xth autocracker if you have enough money
	quit // quit the game (no saving right now)
	`)
	line := strings.Fields(strings.ToLower(readLine()))
	var command string
	switch command = line[0]; command {
	case "quit":
		stopGame(game)
		return
	case "crack":
		fmt.Println("You cracked a pack!")
		crackPacks(game, 1)
	case "sell":
		fmt.Println("You are selling all your cards")
		sell(game)
	case "buy":
		tobuy, _ := strconv.Atoi(line[1])
		buy(game, tobuy)
	}
}

// todo 
// 10 minutes to gather as much money as you can (or crack as many packs as you can)
// high score chart at the end that saves to file
// save/load/offline progress
// sell as a percentage
// basic dueling feature not sure what I want to do (dials to alter about how much of each rarity to use per duel etc)
// seasons (mvp autocracker?) + some bonuses + (pack odds change each season, something about dueling, and stats/order/etc of autocrackers)
// upgrades / synergies / random chance
// achievements
// then make it an actual game with a visual display and sounds and stuff
// can you make it play through a website using WASM?
// basic working prototype for my friends to get feedback on