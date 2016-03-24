package main

import (
    "fmt"
    "strconv"
)

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
    KindMask = 0x07
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

type GetMovesFunc func(*Board, Piece, int, int) *Movelist

var movefunc = map[Piece]GetMovesFunc {
    Pawn : GetPawnMoves,
    Knight : GetKnightMoves,
    Bishop : GetBishopMoves,
    Rook : GetRookMoves,
    Queen : GetQueenMoves,
    King : GetKingMoves,
}

func Color(color Piece) string {
	color &= White
	if color == Black {
		return "Black"
	} else {
		return "White"
	}
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
		//return strings.ToLower(piece[kind])
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
        panic( "Listing moves for an empty piece doesn't even make sense!" )
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
    fmt.Println("GetPawnMoves")
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
    lookr = rank+direction
    if isValid(lookr, file) && b.isEmpty(lookr, file) {
        to.Set(lookr, file)
        moves.AddPair(from, to)
    }

    // two steps forward, only from starting position
    lookr = rank+2*direction
    if (side == White && rank == 1) || (side == Black && rank == 6) {
        if isValid(lookr, file) && b.isEmpty(lookr, file) {
            to.Set(lookr, file)
            moves.AddPair(from, to)
        }
    }

    // capture to the left
    lookr = rank+direction
    lookf = rank-direction
    if isValid(lookr, lookf) && b.isEnemy(lookr, lookf, side) {
        to.Set(lookr, lookf)
        moves.AddPair(from, to)
    }

    // capture to the right
    lookr = rank+direction
    lookf = rank+direction
    if isValid(lookr, lookf) && b.isEnemy(lookr, lookf, side) {
        to.Set(lookr, lookf)
        moves.AddPair(from, to)
    }

    // en passant to the left
    // en passant to the right

    return &moves
}

func GetKnightMoves(b *Board, side Piece, rank, file int) *Movelist {
    fmt.Println("GetKnightMoves")
    return nil
}

func GetBishopMoves(b *Board, side Piece, rank, file int) *Movelist {
    return nil
}

func GetRookMoves(b *Board, side Piece, rank, file int) *Movelist {
    return nil
}

func GetQueenMoves(b *Board, side Piece, rank, file int) *Movelist {
    return nil
}

func GetKingMoves(b *Board, side Piece, rank, file int) *Movelist {
    return nil
}
