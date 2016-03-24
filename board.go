package main

import "fmt"
import "strconv"

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

func (b *Board) Show() {
	for rank := 7; rank >= 0; rank-- {
		line := strconv.Itoa(rank+1) + "|"
		for file := 0; file < 8; file++ {
			line += " " + b[rank][file].Name() + " "
		}
		fmt.Println(line)
	}
	fmt.Println("  -----------------------")
	fmt.Println("   a  b  c  d  e  f  g  h")
}

func (b *Board) isEmpty(rank, file int) bool {
	return b[rank][file].isEmpty()
}

func isValid( rank, file int) bool {
    return !(rank < 0 || rank > 7 || file < 0 || file > 7)
}

func (b *Board) isColor(rank int, file int, color Piece) bool {
	return b[rank][file].isColor(color)
}

func (b *Board) isEnemy(rank int, file int, color Piece) bool {
	return b[rank][file].isColor(color ^ ColorMask)
}

func (b *Board) Apply(m Move) {
    from := m.getFrom()
    to := m.getTo()

    // add logic to record captured pieces
    piece := b.GetPos(from)
    b.SetPos(to, *piece)
    *piece = Empty
}

func (b *Board) CandidateMoves(side Piece) Movelist {
    movelist := make(Movelist, 0)

    for rank := 0; rank < 8; rank++ {
        for file := 0; file <8; file++ {
            if b.Get(rank, file).isColor(side) {
                piece := *b.Get(rank, file)
                //fmt.Println( "found: " + piece.Name() + " at " + PosName(rank, file) )
                movelist.AddList(piece.GetMoves(b, rank, file))
            }
        }
    }

    return movelist
}
