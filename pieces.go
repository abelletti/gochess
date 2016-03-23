package main

import (
	// "fmt"
	"strings"
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

var piece = []string {
	" ",
	"P",
	"N",
	"B",
	"R",
	"Q",
	"K",
}

func (p *Piece) Set( color, kind Piece ) {
    *p = color | kind
}

func (p *Piece) Name() string {
	color := *p & White
	kind := *p & ^White

	if color == White {
		return piece[kind]
	} else {
		return strings.ToLower(piece[kind])
	}
}

func (p *Piece) isEmpty() bool {
    return (*p == 0)
}

func (p *Piece) isColor(col Piece) bool {
    return ((*p & White) == (col & White)) && !p.isEmpty()
}
