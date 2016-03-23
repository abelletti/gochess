package main

import (
	"fmt"
)

func main() {
    var board Board
    turn, turnc := 1, White

    board.Init()
    fmt.Printf( "\nPrior to Turn #%d (%s's move):\n\n", turn, Color(turnc) )
    board.Show()
    fmt.Println()

    fmt.Println( board.isEmpty(0, 0))
    fmt.Println( board.isEmpty(4, 0))
    fmt.Println( board.isColor(0,0,White))
    fmt.Println( board.isColor(0,0,Black))
    fmt.Println( board.isColor(4,4,White))
    fmt.Println( board.isColor(4,4,Black))

    var pos1, pos2 Position
    pos1.Set(2,4)
    pos2.Sets("f2")
    fmt.Println( pos1.Name(), pos2.Name() )

    var move1 Move
    move1.Sets( "b1", "a3" )
    fmt.Println( move1.Name() )
}
