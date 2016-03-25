package main

import "fmt"

type Move struct {
	from  Position
	to    Position
	value int
}

func (m *Move) Sets(from, to string) {
	m.from.Sets(from)
	m.to.Sets(to)
}

func (m *Move) Set(from, to Position) {
	m.from = from
	m.to = to
}

func (m *Move) setScore(value int) {
    m.value = value
}

func (m *Move) getScore() int {
    return m.value
}

func (m *Move) Name() string {
	return fmt.Sprintf("%s-%s", m.from.Name(), m.to.Name())
}

func (m *Move) NameVal() string {
    if m.value != 0 {
	    return fmt.Sprintf("%s-%s(%d)", m.from.Name(), m.to.Name(), m.value)
    } else {
        return m.Name()
    }
}

func (m *Move) getFrom() Position {
	return m.from
}

func (m *Move) getTo() Position {
	return m.to
}
