package main

import (
	"fmt"
	"math/rand"
)

func main() {
	winner := DiceGame(3, 4)
	fmt.Printf("Player %d the Winner\n", winner)
}

func DiceGame(players, dice int) int {
	playerDice := make(map[int][]int) // [player][]dices
	playerTotalPoint := make(map[int]int)

	for i := 0; i < players; i++ {
		playerDice[i] = make([]int, dice, 1000)
	}

	totalPlayer := players

	// Lopping game until remain one player
	for totalPlayer != 1 {

		// Shuffle
		for i := 0; i < len(playerDice); i++ {
			totalDice := len(playerDice[i])
			for j := 0; j < totalDice; j++ {
				diceNumber := shuffleDice()
				playerDice[i][j] = diceNumber
			}
		}

		// Evaluate
		for i := 0; i < players; i++ {

			totalDice := len(playerDice[i])

			for j := 0; j < totalDice; j++ {
				diceNumber := playerDice[i][j]

				if diceNumber == -1 {
					playerDice[i][j] = 1
				}

				if diceNumber == 1 {
					playerDice[i] = removeDice(playerDice[i], j)
					// if last player
					if i == len(playerDice)-1 {
						// will give it to first player
						if isPlayerInGame(len(playerDice[0])) {
							playerDice[0] = append(playerDice[0], 1)
						}
					} else {
						// if not will give to next player
						if isPlayerInGame(len(playerDice[i+1])) {
							playerDice[i+1] = append(playerDice[i+1], -1)
						}
					}

					if len(playerDice[i]) == 0 {
						totalPlayer = totalPlayer - 1
					}

					totalDice--
					j--

				} else if diceNumber == 6 {
					playerDice[i] = removeDice(playerDice[i], j)
					playerTotalPoint[i] += 1

					if len(playerDice[i]) == 0 {
						totalPlayer = totalPlayer - 1
					}
					totalDice--
					j--
				}
			}
		}
	}

	max, winner := 0, 0
	for key, val := range playerTotalPoint {
		if val > max {
			winner = key
		}
	}

	return winner
}

func removeDice(playerDice []int, n int) []int {
	newDice := playerDice[:n]
	newDice = append(newDice, playerDice[n+1:]...)
	return newDice
}

func shuffleDice() int {
	min, max := 1, 6
	return rand.Intn(max-min+1) + min
}

func isPlayerInGame(totalDice int) bool {
	return totalDice != 0
}

// this function is used for debug purpose
// func debugDice(playerDice map[int][]int, point map[int]int) {
// 	keys := make([]int, 0, len(playerDice))
// 	for key := range playerDice {
// 		keys = append(keys, key)
// 	}

// 	sort.Ints(keys)
// 	for _, key := range keys {
// 		val := playerDice[key]
// 		fmt.Printf("%d (%d) %d \n", key, point[key], val)
// 	}
// }
