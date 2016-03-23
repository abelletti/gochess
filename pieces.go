package main

import (
	// "fmt"
	// "strings"
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
	Black = Piece(iota * 128)
	White
)

var piece = [][]rune {
	[]rune("   "),
	[]rune("P♙♟"),
	[]rune("N♘♞"),
	[]rune("B♗♝"),
	[]rune("R♖♜"),
	[]rune("Q♕♛"),
	[]rune("K♔♚"),
}

func Color( color Piece ) string {
    color &= White
    if color == Black {
        return "Black"
    } else {
        return "White"
    }
}

func (p *Piece) Set( color, kind Piece ) {
    *p = color | kind
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
