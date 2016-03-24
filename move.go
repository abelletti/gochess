package main

import "fmt"

type Move struct {
	from Position
	to   Position
}

func (m *Move) Sets(from, to string) {
	m.from.Sets(from)
	m.to.Sets(to)
}

func (m *Move) Set(from, to Position) {
	m.from = from
	m.to = to
}

func (m *Move) Name() string {
	return fmt.Sprintf("%s-%s", m.from.Name(), m.to.Name())
}

func (m *Move) getFrom() Position {
	return m.from
}

func (m *Move) getTo() Position {
	return m.to
}
