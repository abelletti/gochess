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

	/*
	   fmt.Println(board.isEmpty(0, 0))
	   fmt.Println(board.isEmpty(4, 0))
	   fmt.Println(board.isColor(0, 0, White))
	   fmt.Println(board.isColor(0, 0, Black))
	   fmt.Println(board.isColor(4, 4, White))
	   fmt.Println(board.isColor(4, 4, Black))

	   var pos1, pos2 Position
	   pos1.Set(2, 4)
	   pos2.Sets("f2")
	   fmt.Println(pos1.Name(), pos2.Name())

	   var move1, move2 Move
	   move1.Sets("b1", "a3")
	   move2.Sets("a2", "a3")
	   fmt.Println(move1.Name())
	   fmt.Println(move2.Name())

	   movelist := make(Movelist, 0)
	   movelist.Add(move1)
	   movelist.Add(move2)
	   movelist.Show("")
	   movelist.AddList(&movelist)
	   movelist.Show("Doubled")
	*/

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
