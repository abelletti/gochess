package main

import (
	"fmt"
)

func main() {
    var board Board

    board.Init()
    board.Show()

    fmt.Println( board.isEmpty(0, 0))
    fmt.Println( board.isEmpty(4, 0))
    fmt.Println( board.isColor(0,0,White))
    fmt.Println( board.isColor(0,0,Black))
    fmt.Println( board.isColor(4,4,White))
    fmt.Println( board.isColor(4,4,Black))
}
