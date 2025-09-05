package models

const (
	NumNeededForWin = 3
)

type CellState int

const (
	Empty CellState = iota
	X
	O
)

type Move struct {
	NextPlayer     CellState
	AvailableSpots []int
	PreviousState  []CellState
}
