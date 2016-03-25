package main

import (
	"fmt"
	"strconv"
)

var _ = fmt.Printf

type Piece uint8

const (
	Empty = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

const (
	KindMask  = 0x07
	ColorMask = 0x80
)

const (
	Black = Piece(iota * ColorMask)
	White
)

var piece = [][]rune{
	[]rune("   "),
	[]rune("P♙♟"),
	[]rune("N♘♞"),
	[]rune("B♗♝"),
	[]rune("R♖♜"),
	[]rune("Q♕♛"),
	[]rune("K♔♚"),
}

var pieceval = []int{
	0,
	1,
	3,
	3,
	5,
	9,
	1000000,
}

type GetMovesFunc func(*Board, Piece, int, int) *Movelist

var movefunc = map[Piece]GetMovesFunc{
	Pawn:   GetPawnMoves,
	Knight: GetKnightMoves,
	Bishop: GetBishopMoves,
	Rook:   GetRookMoves,
	Queen:  GetQueenMoves,
	King:   GetKingMoves,
}

type MoveDesc struct {
	rankstep int
	filestep int
	longmove bool
}

var knightmoves = []MoveDesc{
	{2, -1, false},
	{2, 1, false},
	{1, 2, false},
	{-1, 2, false},
	{-2, 1, false},
	{-2, -1, false},
	{-1, -2, false},
	{1, -2, false},
}

var bishopmoves = []MoveDesc{
	{1, 1, true},
	{-1, 1, true},
	{-1, -1, true},
	{1, -1, true},
}

var rookmoves = []MoveDesc{
	{1, 0, true},
	{0, 1, true},
	{-1, 0, true},
	{0, -1, true},
}

var queenmoves = append(bishopmoves, rookmoves...)

var kingmoves = []MoveDesc{
	{1, 0, false},
	{1, 1, false},
	{0, 1, false},
	{-1, 1, false},
	{-1, 0, false},
	{-1, -1, false},
	{0, -1, false},
	{1, -1, false},
}

func (p *Piece) Color() string {
	color := *p & ColorMask
	if color == Black {
		return "Black"
	} else {
		return "White"
	}
}

func (p *Piece) Val() int {
    kind := *p & KindMask
    return pieceval[kind]
}

func (p *Piece) Set(color, kind Piece) {
	*p = (color & ColorMask) | (kind & KindMask)
}

func (p *Piece) Name() string {
	color := *p & White
	kind := *p & ^White

	if color == White {
		return string(piece[kind][2:3])
	} else {
		return string(piece[kind][1:2])
	}
}

func (p *Piece) isEmpty() bool {
	return (*p == 0)
}

func (p *Piece) isColor(col Piece) bool {
	return ((*p & White) == (col & White)) && !p.isEmpty()
}

func (p *Piece) GetMoves(b *Board, rank, file int) *Movelist {
	kind := *p & KindMask
	side := *p & ColorMask

	switch kind {
	case Empty:
		panic("Listing moves for an empty piece doesn't even make sense!")
	case Pawn:
		return GetPawnMoves(b, side, rank, file)
	case Knight:
		return GetKnightMoves(b, side, rank, file)
	case Bishop:
		return GetBishopMoves(b, side, rank, file)
	case Rook:
		return GetRookMoves(b, side, rank, file)
	case Queen:
		return GetQueenMoves(b, side, rank, file)
	case King:
		return GetKingMoves(b, side, rank, file)
	default:
		err := "Encountered unknown piece, value = " + strconv.Itoa(int(kind))
		panic(err)
	}
}

func GetPawnMoves(b *Board, side Piece, rank, file int) *Movelist {
	//    fmt.Println("GetPawnMoves")
	moves := make(Movelist, 0)
	var to, from Position
	var direction int
	var lookr, lookf int

	from.Set(rank, file)

	if side == White {
		direction = 1
	} else {
		direction = -1
	}

	// single step forward
	lookr = rank + direction
	if isValid(lookr, file) && b.isEmpty(lookr, file) {
		to.Set(lookr, file)
		moves.AddPair(from, to)
	}

	// two steps forward, only from starting position
	lookr = rank + 2*direction
	if (side == White && rank == 1) || (side == Black && rank == 6) {
		if isValid(lookr, file) && b.isEmpty(lookr, file) && b.isEmpty(lookr-direction, file) {
			to.Set(lookr, file)
			moves.AddPair(from, to)
		}
	}

	// capture to the left
	lookr = rank + direction
	lookf = file - direction
	if isValid(lookr, lookf) && b.isEnemy(lookr, lookf, side) {
		to.Set(lookr, lookf)
		moves.AddPair(from, to)
	}

	// capture to the right
	lookr = rank + direction
	lookf = file + direction
	if isValid(lookr, lookf) && b.isEnemy(lookr, lookf, side) {
		to.Set(lookr, lookf)
		moves.AddPair(from, to)
	}

	// en passant to the left
	// en passant to the right

	return &moves
}

func GetKnightMoves(b *Board, side Piece, rank, file int) *Movelist {
	//    fmt.Println("GetKnightMoves")
	moves := make(Movelist, 0)
	var to, from Position
	var lookr, lookf int

	from.Set(rank, file)

	for _, movedesc := range knightmoves {
		lookr = rank + movedesc.rankstep
		lookf = file + movedesc.filestep
		if isValid(lookr, lookf) && (b.isEnemy(lookr, lookf, side) || b.isEmpty(lookr, lookf)) {
			to.Set(lookr, lookf)
			moves.AddPair(from, to)
		}
	}

	return &moves
}

func GetBishopMoves(b *Board, side Piece, rank, file int) *Movelist {
	return GetLongMoves(b, side, rank, file, bishopmoves)
}

func GetRookMoves(b *Board, side Piece, rank, file int) *Movelist {
	return GetLongMoves(b, side, rank, file, rookmoves)
}

func GetQueenMoves(b *Board, side Piece, rank, file int) *Movelist {
	return GetLongMoves(b, side, rank, file, queenmoves)
}

func GetKingMoves(b *Board, side Piece, rank, file int) *Movelist {
	return GetLongMoves(b, side, rank, file, kingmoves)
}

func GetLongMoves(b *Board, side Piece, rank, file int, movestyle []MoveDesc) *Movelist {
	moves := make(Movelist, 0)
	var to, from Position
	var lookr, lookf int

	from.Set(rank, file)

	for _, movedesc := range movestyle {
		for distance := 1; distance <= 7; distance++ {
			lookr = rank + movedesc.rankstep*distance
			lookf = file + movedesc.filestep*distance
			if isValid(lookr, lookf) && (b.isEnemy(lookr, lookf, side) || b.isEmpty(lookr, lookf)) {
				to.Set(lookr, lookf)
				moves.AddPair(from, to)
			}
			if !isValid(lookr, lookf) || !b.isEmpty(lookr, lookf) || !movedesc.longmove {
				break
			}
		}
	}

	return &moves
}
