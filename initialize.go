package main

import (
	"math/rand"
	"time"
	"fmt"
)

func initGame() State {
	// eventually load the state from a file here
	startTime := time.Now().UTC().Unix()
	rand.Seed(startTime)
	s := State{
		last_time: startTime,
		running:  true,
		stats:    Stats{},
		inventory: Inventory{
			card_types: []Rarity{
				{name: "Commons", value: COMMONVAL, chance: COMMONCHANCE},
				{name: "Rare", value: RAREVAL, chance: RARECHANCE},
			},
		},
		shop: []AutoCracker{
			AutoCracker{name: "Sweaty Nerd", rate: 1, price: 50.0},
			AutoCracker{name: "Shop Lackey", rate: 2, price: 500.0},
			AutoCracker{name: "Auto Sorter", rate: 5, price: 10000.0},
		},
	}
	return s
}

func main() {
	fmt.Println("Welcome to cardsim")
	game := initGame()
	for game.running {
		update(&game)
		displayInfo(game)
		commandLineAction(&game)
	}
	fmt.Printf("Game done")
}