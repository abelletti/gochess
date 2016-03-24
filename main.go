package main

import (
	"fmt"
)

func main() {
	var board Board
	turn, side := 1, White

	board.Init()
	fmt.Printf("\nPrior to Turn #%d (%s's move):\n\n", turn, Color(side))
	board.Show()
	fmt.Println()

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

    moves := board.CandidateMoves(side)
    moves.Show("Candidate Moves for "+Color(side) )
}
