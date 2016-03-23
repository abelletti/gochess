package main

import "fmt"

/*
 * Board array is row, col with 0,0 at bottom-left
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

	for col := 0; col < 8; col++ {
		b[1][col].Set(White, Pawn)
		b[6][col].Set(Black, Pawn)
	}
}

func (b *Board) Show() {
	for row := 7; row >= 0; row-- {
		letter := make([]byte, 1)
		letter[0] = byte('a') + byte(row)
		line := string(letter[0:1]) + "|"
		for col := 0; col < 8; col++ {
			line += " " + b[row][col].Name() + " "
		}
		fmt.Println(line)
	}
	fmt.Println("  -----------------------")
	fmt.Println("   1  2  3  4  5  6  7  8")
}

func (b *Board) isEmpty(row, col int) bool {
    return b[row][col].isEmpty()
}

func (b *Board) isColor(row int, col int, color Piece) bool {
    return b[row][col].isColor(color)
}


