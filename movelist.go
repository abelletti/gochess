package main

import (
    "fmt"
    "strings"
)

type Movelist []Move

func (ml *Movelist) Show(title string) {
    if title != "" {
        fmt.Println(title+"\n"+strings.Repeat("-", len(title)))
    }
	for i := range *ml {
		fmt.Println([]Move(*ml)[i].Name())
	}
}

func (ml *Movelist) Add(m Move) {
	*ml = append(*ml, m)
}


func (ml *Movelist) AddList(ml2 *Movelist) {
	*ml = append(*ml, *ml2...)
}
