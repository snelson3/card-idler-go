package main

import (
	"time"
)

// values of different rarities
const COMMONVAL = 1 
const RAREVAL = 3 // value of rare rarity

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

func getPacksPerSecond(s *State) int{
	// abstract into function
	packs_to_crack := 0
	for r := range s.shop {
		selcracker := s.shop[r]
		packs_to_crack += selcracker.held * selcracker.rate
	}
	return packs_to_crack
}

func update(s *State) {
	tick := time.Now().UTC().Unix()
	seconds_passed := int(tick - s.last_time)
	packs_to_crack := getPacksPerSecond(s) * seconds_passed

	crackPacks(s, packs_to_crack)
	s.last_time = tick
}


func stopGame(s *State) {
	// Eventually save the state to a file here
	s.running = false
}
