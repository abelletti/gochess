package main

import "fmt"

type Movelist []Move

func (ml *Movelist) Show() {
    for i := range *ml {
        fmt.Println( []Move(*ml)[i].Name() )
    }
}

func (ml *Movelist) Add( m Move ) {
    *ml = append(*ml,m)
}
