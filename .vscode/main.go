package main

import "fmt"

func rollDices(dices int, faces int, target int, mapa [string]int) int {
	if target == 0 {
		return 1
	}

	if target < 0 {
		return 0
	}
	v, ok := mapa[fmt.Sprint(dices, target)]
	if ok {
		return v
	}
	var sum int

	for i := 1; i < faces+1; i++ {
		sum += rollDices(dices-1, faces, target-i, mapa)
	}
	mapa[fmt.Sprint(dices, target)] = sum
	return sum
}
func numRollsToTarget(d int, f int, target int) int {
	mapa := make(map[string]int)
	return rollDices(d, f, target, mapa)
}

func main() {
	numRollsToTarget(3, 6, 10)
}
