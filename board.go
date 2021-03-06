package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

/*
 * Board array is rank, file with 0,0 at bottom-left
 * White at bottom, black at top
 */

type Board [8][8]Piece

func (b *Board) Init() {
	b[0][0].Set(White, Rook)
	b[0][1].Set(White, Knight)
	b[0][2].Set(White, Bishop)
	b[0][3].Set(White, Queen)
	b[0][4].Set(White, King)
	b[0][5].Set(White, Bishop)
	b[0][6].Set(White, Knight)
	b[0][7].Set(White, Rook)
	b[7][0].Set(Black, Rook)
	b[7][1].Set(Black, Knight)
	b[7][2].Set(Black, Bishop)
	b[7][3].Set(Black, Queen)
	b[7][4].Set(Black, King)
	b[7][5].Set(Black, Bishop)
	b[7][6].Set(Black, Knight)
	b[7][7].Set(Black, Rook)

	for file := 0; file < 8; file++ {
		b[1][file].Set(White, Pawn)
		b[6][file].Set(Black, Pawn)
	}
}

func (b *Board) Set(rank, file int, p Piece) {
	(*b)[rank][file] = p
}

func (b *Board) SetPos(pos Position, p Piece) {
	rank, file := pos.getRankFile()
	(*b)[rank][file] = p
}

func (b *Board) Get(rank, file int) *Piece {
	return &b[rank][file]
}

func (b *Board) GetPos(pos Position) *Piece {
	rank, file := pos.getRankFile()
	return &b[rank][file]
}

func (b *Board) isEmpty(rank, file int) bool {
	return b.Get(rank, file).isEmpty()
}

func (b *Board) isEmptyPos(pos Position) bool {
	return b.GetPos(pos).isEmpty()
}

func (b *Board) Show() {
	b.ShowMarkedbyRankFile(-1, -1)
}

func (b *Board) ShowMarked(pos Position) {
	b.ShowMarkedbyRankFile(pos.getRankFile())
}

func (b *Board) ShowMarkedbyRankFile(rankmark, filemark int) {
	for rank := 7; rank >= 0; rank-- {
		line := strconv.Itoa(rank+1) + "|"
		for file := 0; file < 8; file++ {
			if rank == rankmark && file == filemark {
				line += ">"
			} else {
				line += " "
			}
			line += b[rank][file].Name() + " "
		}
		fmt.Println(line)
	}
	fmt.Println("  -----------------------")
	fmt.Println("   a  b  c  d  e  f  g  h")
}

func isValid(rank, file int) bool {
	return !(rank < 0 || rank > 7 || file < 0 || file > 7)
}

func (b *Board) isColor(rank int, file int, color Piece) bool {
	return b[rank][file].isColor(color)
}

func (b *Board) isEnemy(rank int, file int, color Piece) bool {
	return b[rank][file].isColor(color ^ ColorMask)
}

func (b *Board) ApplyNew(m Move) (*Board, int) {
	from := m.getFrom()
	to := m.getTo()
	newb := *b
	var capvalue int

	piece := newb.GetPos(from)
	if !newb.isEmptyPos(to) {
		capture := newb.GetPos(to)
		capvalue = capture.Val()
		//fmt.Println(piece.Color() + " would capture: " + capture.Name())
	}
	newb.SetPos(to, *piece)
	*piece = Empty

	return &newb, capvalue
}

func (b *Board) Apply(m Move) {
	from := m.getFrom()
	to := m.getTo()

	piece := b.GetPos(from)
	var capture Piece = Empty
	if !b.isEmptyPos(to) {
		// add logic to record captured pieces
		capture = *(b.GetPos(to))
		fmt.Println(piece.Color() + " has captured: " + (&capture).Name())
	}
	b.SetPos(to, *piece)
	*piece = Empty
}

func (b *Board) CandidateMoves(side Piece) Movelist {
	movelist := make(Movelist, 0)

	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			if b.Get(rank, file).isColor(side) {
				piece := *b.Get(rank, file)
				//fmt.Println( "found: " + piece.Name() + " at " + PosName(rank, file) )
				movelist.AddList(piece.GetMoves(b, rank, file))
			}
		}
	}

	return movelist
}

func (b *Board) isCheckMate(side Piece) bool {
	if !b.isCheck(side) {
		return false
	}
	// we're in check; do we have a move?
	moves := b.CandidateMoves(side)
	moves.Show("Testing check: Candidate Moves for " + side.Color() + ": ")
	moves = moves.PruneForCheck(b, side)
	moves.Show("Testing check: Pruned Candidate Moves for " + side.Color() + ": ")

	return len(moves) == 0
}

func (b *Board) isCheck(side Piece) bool {
	king := b.findKing(side)
	opponent := side ^ ColorMask
	//fmt.Println( "Found " + side.Color() + " king at " + king.Name())
	moves := b.CandidateMoves(opponent)
	return moves.landsOn(king)
}

func (b *Board) findKing(side Piece) Position {
	var pos Position
	var p Piece
	p.Set(side, King)

	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			if (*b)[rank][file] == p {
				pos.Set(rank, file)
				return pos
			}
		}
	}

	err := "Failed to find " + side.Color() + " king!"
	panic(err)
}

func (b *Board) ChooseBestWithDepth(side Piece, depth int, incheck bool) (Move, bool) {
	return b.ChooseBestWithDepthInner(side, depth, incheck, true)
}

func (b *Board) ChooseBestWithDepthInner(side Piece, depth int, incheck, myself bool) (Move, bool) {
	var bestmove Move

	moves := b.CandidateMoves(side) // what legal moves can I make?
	//  moves.Show("Candidate Moves for " + side.Color() + ": ")
	moves = moves.PruneForSelfCheck(b, side) // cannot put myself into check, remove those
	//  at this point, each "moves" entry is marked with a point value
	//  moves.Show("Pruned Candidate Moves for " + side.Color() + ": ")

	if len(moves) == 0 {
		if incheck {
			bestmove.setScore(-1000000)
			return bestmove, true // checkmate
		} else {
			bestmove.setScore(5)
			return bestmove, true // this way leads stalemate
		}
	}

	//  fmt.Printf( "BEFORE: depth = %d, len(moves) = %d\n", depth, len(moves))

	if depth > 1 {
		// evaluate opponent options for each of our moves
		for movenum := range moves {
			newb, _ := b.ApplyNew(moves[movenum])
			check := newb.isCheck(*(side.Other()))
			best, _ := newb.ChooseBestWithDepthInner(*(side.Other()), depth-1, check, !myself)
			//          fmt.Println( "Got back: " + best.NameVal())
			moves[movenum].addScore(best.getScore())
			//            if myself {
			//                // "good" opponent moves are bad for me
			//                moves[movenum].addScore(-1 * best.getScore())
			//            } else {
			//                // but good moves for me are lovely
			//                moves[movenum].addScore(best.getScore())
			//            }
			//          fmt.Printf( "Move score = %d\n", moves[movenum].getScore())
		}
	}
	//  fmt.Printf( "AFTER: depth = %d, len(moves) = %d\n", depth, len(moves))

	// otherwise choose the best of what we've got
	topscore := -1000000000

	// identify score of top move(s)
	for movenum := range moves {
		movescore := moves[movenum].getScore()
		if movescore > topscore {
			topscore = movescore
		}
	}

	// and choose randomly from amongst identical scores
	var choices []int
	for movenum := range moves {
		if moves[movenum].getScore() == topscore {
			choices = append(choices, movenum)
		}
	}
	//  fmt.Printf("len(choices) = %d\n", len(choices))
	pick := choices[rand.Intn(len(choices))]

	//  fmt.Printf("Top scoring move worth %d (%d identical)\n", topscore, len(choices))
	return moves[pick], false
}
