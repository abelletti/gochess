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
        incheck := board.isCheck(side)
        if incheck {
            fmt.Println("Oh god help me, I'm in CHECK!")
        }
		moves := board.CandidateMoves(side)
		if len(moves) == 0 {
			fmt.Println(side.Color() + " has no moves remaining, STALEMATE?")
			break
		}
		moves.Show("Candidate Moves for " + side.Color())
        moves = moves.PruneForCheck(&board,side)
		moves.Show("Pruned Candidate Moves for " + side.Color())
        if len(moves)==0 && incheck {
            fmt.Println(side.Color() + " has no moves to exit CHECK; CHECKMATE!")
            break
        }

        var chosenmove Move

        if side == White {
		    chosenmove = moves.ChooseRandom(side)
        } else { // Black's move
            chosenmove = moves.ChooseFirst(side)
        }

		fmt.Println("Chosen move: " + chosenmove.Name())
		board.Apply(chosenmove)
		fmt.Printf("\nAfter Turn #%d (%s's move):\n\n", turn, side.Color())
		board.Show()
		fmt.Println()

		side ^= ColorMask
        turn++

		fmt.Scanln(&line)
	}

}
