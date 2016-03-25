package main

import (
	"fmt"
	"math/rand"
)

type Movelist []Move

func (ml *Movelist) Show(title string) {
	fmt.Print(title)
	for i := range *ml {
		fmt.Print([]Move(*ml)[i].NameVal() + " ")
	}
	fmt.Println()
}

func (ml *Movelist) landsOn(pos Position) bool {
	for movenum := range *ml {
		if pos == (*ml)[movenum].getTo() {
			return true
		}
	}
	return false
}

func (ml *Movelist) PruneForSelfCheck(b *Board, side Piece) Movelist {
	var newmoves Movelist

	for movenum := range *ml {
		move := (*ml)[movenum]
		newb, capvalue := b.ApplyNew(move)
		if !newb.isCheck(side) {
			move.setScore(capvalue)
			newmoves.Add(move)
		} else {
			//fmt.Println("Would be moving into CHECK: "+move.Name())
		}
	}

	return newmoves
}

func (ml *Movelist) PruneForCheck(b *Board, side Piece) Movelist {
	var newmoves Movelist

	for movenum := range *ml {
		move := (*ml)[movenum]
		newb, capvalue := b.ApplyNew(move)
		if !newb.isCheck(side) {
			if newb.isCheckMate(*(side.Other())) {
				capvalue = 1000000
				move.setScore(capvalue)
				var matemoves Movelist
				matemoves.Add(move)
				return matemoves // not much point in alteratives, this is mate
			}
			move.setScore(capvalue)
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

func (ml *Movelist) ChooseNoDepth(side Piece) Move {
	var topscore int

	// identify score of top move(s)
	for movenum := range *ml {
		movescore := (*ml)[movenum].getScore()
		if movescore > topscore {
			topscore = movescore
		}
	}

	// and choose randomly from amongst identical scores
	var choices []int
	for movenum := range *ml {
		if (*ml)[movenum].getScore() == topscore {
			choices = append(choices, movenum)
		}
	}
	pick := choices[rand.Intn(len(choices))]

	fmt.Printf("Top scoring move worth %d (%d identical)\n", topscore, len(choices))
	return (*ml)[pick]
}
