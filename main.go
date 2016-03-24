package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var board Board
	turn, side := 1, White

	seed := time.Now().UnixNano()
	fmt.Printf("Random number seed is %d\n", seed)
	rand.Seed(seed)

	board.Init()

	for {
		var line string

		fmt.Printf("\nPrior to Turn #%d (%s's move):\n\n", turn, side.Color())
		board.Show()
		fmt.Println()
		moves := board.CandidateMoves(side)
		if len(moves) == 0 {
			fmt.Println(side.Color() + " has no moves remaining!")
			break
		}
		moves.Show("Candidate Moves for " + side.Color())
        moves = moves.PruneForCheck(&board,side)
		moves.Show("Pruned Candidate Moves for " + side.Color())
		chosenmove := moves.Choose(side)
		fmt.Println("Chosen move: " + chosenmove.Name())
		board.Apply(chosenmove)
		fmt.Printf("\nAfter Turn #%d (%s's move):\n\n", turn, side.Color())
		board.Show()
		fmt.Println()

        if board.isCheck(side) {
            fmt.Println( side.Color() + " is in CHECK!")
        }
        opponent := side ^ ColorMask
        if board.isCheck(opponent) {
            fmt.Println( opponent.Color() + " is in CHECK!")
        }

		side ^= ColorMask
        turn++

		fmt.Scanln(&line)
	}

}
