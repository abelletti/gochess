package main

import (
	"fmt"
	"strconv"
)

type Position struct {
	rank int8
	file int8
}

func (p *Position) Sets(pos string) {
	rank, _ := strconv.Atoi(pos[1:])
	p.rank = int8(rank - 1)
	p.file = int8(rune(pos[0]) - 'a')
}

func (p *Position) Set(rank, file int) {
	p.rank = int8(rank)
	p.file = int8(file)
}

func (p *Position) Name() string {
	return fmt.Sprintf("%c%d", int('a')+int(p.file), p.rank+1)
}

func PosName(rank, file int) string {
    var p Position
    p.Set(rank, file)
    return p.Name()
}
