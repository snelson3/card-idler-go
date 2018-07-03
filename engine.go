package main

func crackPacks(s *State, i int) {
	for c := 0; c < i; c++ {
		s.stats.packsOpened++
		s.inventory.commons.held += 9
		s.inventory.rares.held++
	}
}

func getInventoryValue(s State) float64 {
	val := 0.0
	val += float64(s.inventory.commons.held) * s.inventory.commons.value
	val += float64(s.inventory.rares.held) * s.inventory.rares.value
	return val
}

func buy(s *State) string {
	// Right now just buys 1 autocracker

	s.shop[0].held++
}
