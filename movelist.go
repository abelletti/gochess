package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Movelist []Move

func (ml *Movelist) Show(title string) {
	if title != "" {
		fmt.Println(title + "\n" + strings.Repeat("-", len(title)))
	}
	for i := range *ml {
		fmt.Print([]Move(*ml)[i].Name() + " ")
	}
	fmt.Println()
}

func (ml *Movelist) landsOn(pos Position) bool {
    for movenum := range (*ml) {
        if pos == (*ml)[movenum].getTo() {
            return true
        }
    }
    return false
}

func (ml *Movelist) PruneForCheck(b *Board, side Piece) Movelist {
    var newmoves Movelist

    for movenum := range (*ml) {
        move := (*ml)[movenum]
        newb := b.ApplyNew(move)
        if !newb.isCheck(side) {
            newmoves.Add(move)
        } else {
            //fmt.Println("Would be moving into CHECK: "+move.Name())
        }
    }

    return newmoves
}

func (ml *Movelist) Add(m Move) {
	*ml = append(*ml, m)
}

func (ml *Movelist) AddPair(to, from Position) {
	var move Move
	move.Set(to, from)
	*ml = append(*ml, move)
}

func (ml *Movelist) AddList(ml2 *Movelist) {
	if ml2 != nil {
		*ml = append(*ml, *ml2...)
	}
}

func (ml *Movelist) ChooseRandom(side Piece) Move {
	pick := rand.Intn(len(*ml))
	return (*ml)[pick]
}

func (ml *Movelist) ChooseFirst(side Piece) Move {
	pick := 0
	return (*ml)[pick]
}
