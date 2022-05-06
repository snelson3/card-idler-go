package main

import (
	"fmt"
	"math/rand"
)

const CARDSPERPACK = 5

const BUYPRICERATE = 0.5

func crackPacks(s *State, i int) {
	for c := 0; c < i; c++ {
		s.stats.packs_opened++
		for ci := 0; ci < CARDSPERPACK; ci++ {
			rarity := &s.inventory.card_types[0]
			failed := false
			for r := range s.inventory.card_types {
				selrarity := &s.inventory.card_types[r]
				
				// I'm eventually going to need to abstract this randomness it's far too expensive an operation
				if !failed {
					randnum := rand.Float64() // rand between 0 and 1
					if randnum < selrarity.chance {
						rarity = selrarity
					} else {
						failed = true
					}
				}
			}
			rarity.held++
		}
	}
}

func getInventoryValue(s State) float64 {
	val := 0.0
	for r := range s.inventory.card_types {
		selrarity := s.inventory.card_types[r]
		val += float64(selrarity.held) * selrarity.value
	}
	return val
}

func buy(s *State, i int) {
	// Right now just buys 1 autocracker
	tobuy := &s.shop[i]
	if tobuy.price <= s.inventory.cash {
		tobuy.held++
		s.inventory.cash -= tobuy.price
		tobuy.price += tobuy.price * BUYPRICERATE
		fmt.Println("You bought a %s", tobuy.name) 
	}
}

func sell(s *State) {
	// Right now sells everything
	// eventually sell percentage
	s.inventory.cash += getInventoryValue(*s)
	for r := range s.inventory.card_types {
		s.inventory.card_types[r].held = 0
	}
}