package game_test

import (
	"WorkWhileAssignment/pkg/game"
	"WorkWhileAssignment/pkg/models"
	"testing"
)

func newBoard() [][]models.CellState {
	b, _ := game.InitializeBoard(3, 3)
	return b
}

func TestWinHorizontal(t *testing.T) {
	board := newBoard()
	board[0][0] = models.X
	board[0][1] = models.X
	board[0][2] = models.X

	if !game.WinHorizontal(0, 2, board) {
		t.Errorf("expected win but got false")
	}

	board = newBoard()
	board[1][0] = models.X
	board[1][1] = models.O
	board[1][2] = models.X

	if game.WinHorizontal(1, 2, board) {
		t.Errorf("expected no win but got true")
	}
}

func TestWinVertical(t *testing.T) {
	board := newBoard()
	board[0][0] = models.O
	board[1][0] = models.O
	board[2][0] = models.O

	if !game.WinVertical(2, 0, board) {
		t.Errorf("expected win but got false")
	}

	board = newBoard()
	board[0][1] = models.X
	board[1][1] = models.O
	board[2][1] = models.X

	if game.WinVertical(2, 1, board) {
		t.Errorf("expected no win but got true")
	}
}

func TestWinLeftUpRightDown(t *testing.T) {
	board := newBoard()
	board[0][0] = models.X
	board[1][1] = models.X
	board[2][2] = models.X

	if !game.WinLeftUpRightDown(2, 2, board) {
		t.Errorf("expected win but got false")
	}

	board = newBoard()
	board[0][0] = models.X
	board[1][1] = models.O
	board[2][2] = models.X

	if game.WinLeftUpRightDown(2, 2, board) {
		t.Errorf("expected no win but got true")
	}
}

func TestWinLeftDownRightUpDiagonal(t *testing.T) {
	board := newBoard()
	board[2][0] = models.O
	board[1][1] = models.O
	board[0][2] = models.O

	if !game.WinLeftDownRightUpDiagonal(0, 2, board) {
		t.Errorf("expected win but got false")
	}

	board = newBoard()
	board[2][0] = models.X
	board[1][1] = models.O
	board[0][2] = models.X

	if game.WinLeftDownRightUpDiagonal(0, 2, board) {
		t.Errorf("expected no win but got true")
	}
}

func TestGameIsComplete_Draw(t *testing.T) {
	board := newBoard()
	// tied board
	board = [][]models.CellState{
		{models.X, models.O, models.X},
		{models.X, models.X, models.O},
		{models.O, models.X, models.O},
	}
	var avail [][]int

	if !game.GameIsComplete(2, 2, board, avail) {
		t.Errorf("expected game complete but got false")
	}
}

func TestGameIsComplete_Win(t *testing.T) {
	board := newBoard()
	board[0][0] = models.X
	board[1][1] = models.X
	board[2][2] = models.X

	avail := [][]int{{0, 1}, {0, 2}} // don't need to pass all available bc the game shouldn't advance anyway
	if !game.GameIsComplete(2, 2, board, avail) {
		t.Errorf("expected game complete but got false")
	}
}
