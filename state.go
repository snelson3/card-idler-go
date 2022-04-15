package main

import (
	"math/rand"
	"time"
)

// values of different rarities
const COMMONVAL = 1 
const RAREVAL = 5 // value of rare rarity

const COMMONCHANCE = 1.0
const RARECHANCE = 0.3

// Rarity Keeps track of how much of one rarity the player owns
type Rarity struct {
	name  string
	held  int
	value float64
	chance float64
}


// Each successive rarity is more rare than the last
type Inventory struct {
	card_types []Rarity
	cash 		float64
}

// Stats keeps track of various stats
type Stats struct {
	packs_opened int
}

// State holds all the relevant game state information
type State struct {
	stats     Stats
	inventory Inventory
	running   bool
	shop      []AutoCracker
	last_time  int64
}

func update(s *State) {
	tick := time.Now().UTC().Unix()
	seconds_passed := int(tick - s.last_time)
	packs_to_crack := 0
	for r := range s.shop {
		selcracker := s.shop[r]
		packs_to_crack += selcracker.held * selcracker.rate* seconds_passed
	}
	crackPacks(s, packs_to_crack)
	s.last_time = tick
}

func initGame() State {
	// eventually load the state from a file here
	rand.Seed(time.Now().UTC().UnixNano())
	s := State{
		last_time: time.Now().UTC().Unix(),
		running:  true,
		stats:    Stats{},
		inventory: Inventory{
			card_types: make([]Rarity, 10)},
		shop: []AutoCracker{
			AutoCracker{name: "Sweaty Nerd", rate: 1, price: 5.0},
			AutoCracker{name: "Shop Lackey", rate: 3, price: 20.0},
			AutoCracker{name: "Auto Sorter", rate: 10, price: 100.0},
		},
	}
	s.inventory.card_types[0] = Rarity{name: "Commons", value: COMMONVAL, chance: COMMONCHANCE}
	s.inventory.card_types[1] = Rarity{name: "Rare", value: RAREVAL, chance: RARECHANCE}
	return s
}

func stopGame(s *State) {
	// Eventually save the state to a file here
	s.running = false
}
