package main

import (
	"math/rand"
	"time"
)

// COMMONVAL value of common rarity
const COMMONVAL = 1 // value of common rarity
// RAREVAL value of rare rarity
const RAREVAL = 5 // value of rare rarity

// Rarity Keeps track of how much of one rarity the player owns
type Rarity struct {
	name  string
	held  int
	value float64
}

// Inventory I believe this should get replaced with a Map so I can use generated rarities
// But I don't know how to use maps yet
type Inventory struct {
	commons Rarity
	rares   Rarity
}

// Stats keeps track of various stats
type Stats struct {
	packsOpened int
}

// State holds all the relevant game state information
type State struct {
	stats     Stats
	inventory Inventory // Should be a map
	running   bool
	shop      []AutoCracker
	lastTime  int64
}

func update(s *State) {
	tick := time.Now().UTC().Unix()
	secondsPassed := int(tick - s.lastTime)
	crackPacks(s, secondsPassed*s.shop[0].held)
	s.lastTime = tick
}

func initGame() State {
	// eventually load the state from a file here
	rand.Seed(time.Now().UTC().UnixNano())
	s := State{
		lastTime: time.Now().UTC().Unix(),
		running:  true,
		stats:    Stats{},
		inventory: Inventory{commons: Rarity{name: "Commons", value: COMMONVAL},
			rares: Rarity{name: "Rares", value: RAREVAL}},
		shop: []AutoCracker{AutoCracker{name: "Shop Lackey", rate: 1}},
	}
	return s
}

func stopGame(s *State) {
	// Eventually save the state to a file here
	s.running = false
}
