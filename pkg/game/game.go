package game

import (
	"WorkWhileAssignment/pkg/models"
	"fmt"
	"strings"
)

var (
	BoardHeight    int
	BoardWidth     int
	CompletedGames map[string]struct{} // set of unique completed board states
	visited        map[string]struct{} // seen board state + player positions to avoid searching down the same path
)

func InitializeBoard(boardHeight, boardWidth int) ([][]models.CellState, [][]int) {
	BoardWidth = boardWidth
	BoardHeight = boardHeight

	// initialize to avoid panic LOL oops
	CompletedGames = make(map[string]struct{})
	board := make([][]models.CellState, boardHeight)
	availableSpaces := make([][]int, 0, boardHeight*boardWidth)
	visited = make(map[string]struct{})

	for i := 0; i < boardHeight; i++ {
		board[i] = make([]models.CellState, boardWidth)
		for j := 0; j < boardWidth; j++ {
			board[i][j] = models.Empty
			availableSpaces = append(availableSpaces, []int{i, j})
		}
	}

	return board, availableSpaces
}

// Play recursively checks all possible games sequences given a current board state and a current player.
// It uses DFS and follows each possible move to completion, backtracking before trying an alternate path.
// To prevent redundant checks, it stops going down a path if it's already been traversed before or is the board is in a "completed" state.
func Play(currentPlayer models.CellState, boardState [][]models.CellState, availableSpaces [][]int) {
	// skip previously seen board states
	k := getVisitedState(boardState, currentPlayer)
	if _, ok := visited[k]; ok {
		return
	}
	visited[k] = struct{}{}

	// remove current move from available spaces
	for i, availableSpace := range availableSpaces {
		row, col := availableSpace[0], availableSpace[1]
		newAvailable := make([][]int, 0, len(availableSpaces)-1)
		newAvailable = append(newAvailable, availableSpaces[:i]...)
		newAvailable = append(newAvailable, availableSpaces[i+1:]...)

		// make current move
		boardState[row][col] = currentPlayer

		if GameIsComplete(row, col, boardState, newAvailable) {
			recordState(boardState)
		} else {
			Play(NextPlayer(currentPlayer), boardState, newAvailable)
		}
		// backtrack so subsequent recursions see the original board state
		boardState[row][col] = models.Empty
	}
}

// GameIsComplete checks for completion by checking if no further moves are possible (tie)
// or if the last move resulted in 3 consecutive pieces of the same time in any direction (win)
func GameIsComplete(latestMoveRow, latestMoveCol int, boardState [][]models.CellState, availableSpaces [][]int) bool {
	// check tie
	if len(availableSpaces) == 0 {
		// nothing left to do, tied game
		return true
	}

	// check win
	if WinHorizontal(latestMoveRow, latestMoveCol, boardState) ||
		WinVertical(latestMoveRow, latestMoveCol, boardState) ||
		WinLeftUpRightDown(latestMoveRow, latestMoveCol, boardState) ||
		WinLeftDownRightUpDiagonal(latestMoveRow, latestMoveCol, boardState) {
		return true
	}

	// game on!
	return false

}

func WinHorizontal(latestMoveRow, latestMoveCol int, boardState [][]models.CellState) bool {
	lastPlayer := boardState[latestMoveRow][latestMoveCol]
	count := 1 // we start assuming that the last move counts as a potential winning move in the sequence

	r, c := latestMoveRow, latestMoveCol-1
	for isValidBoardPosition(r, c) && boardState[r][c] == lastPlayer {
		count++
		c--
	}

	// Oops you need to continue checking along the path for the case where the
	// last position was placed in the middle of a valid path 🙃
	r, c = latestMoveRow, latestMoveCol+1
	for isValidBoardPosition(r, c) && boardState[r][c] == lastPlayer {
		count++
		c++
	}

	return count >= models.NumNeededForWin
}

func WinVertical(latestMoveRow, latestMoveCol int, boardState [][]models.CellState) bool {
	lastPlayer := boardState[latestMoveRow][latestMoveCol]
	count := 1 // we start assuming that the last move counts as a potential winning move in the sequence

	r, c := latestMoveRow-1, latestMoveCol
	for isValidBoardPosition(r, c) && boardState[r][c] == lastPlayer {
		count++
		r--
	}

	r, c = latestMoveRow+1, latestMoveCol
	for isValidBoardPosition(r, c) && boardState[r][c] == lastPlayer {
		count++
		r++
	}

	return count >= models.NumNeededForWin
}

func WinLeftUpRightDown(latestMoveRow, latestMoveCol int, boardState [][]models.CellState) bool {
	lastPlayer := boardState[latestMoveRow][latestMoveCol]
	count := 1 // we start assuming that the last move counts as a potential winning move in the sequence

	// up-left
	r, c := latestMoveRow-1, latestMoveCol-1
	for isValidBoardPosition(r, c) && boardState[r][c] == lastPlayer {
		count++
		r--
		c--
	}

	// down-right
	r, c = latestMoveRow+1, latestMoveCol+1
	for isValidBoardPosition(r, c) && boardState[r][c] == lastPlayer {
		count++
		r++
		c++
	}

	return count >= models.NumNeededForWin
}

func WinLeftDownRightUpDiagonal(latestMoveRow, latestMoveCol int, boardState [][]models.CellState) bool {
	lastPlayer := boardState[latestMoveRow][latestMoveCol]
	count := 1 // we start assuming that the last move counts as a potential winning move in the sequence

	// down-left
	r, c := latestMoveRow+1, latestMoveCol-1
	for isValidBoardPosition(r, c) && boardState[r][c] == lastPlayer {
		count++
		r++
		c--
	}

	// up-right
	r, c = latestMoveRow-1, latestMoveCol+1
	for isValidBoardPosition(r, c) && boardState[r][c] == lastPlayer {
		count++
		r--
		c++
	}

	return count >= models.NumNeededForWin
}

func NextPlayer(current models.CellState) models.CellState {
	if current == models.X {
		return models.O
	} else {
		return models.X
	}
}

func GetNumUniqueCompletedGames() int {
	return len(CompletedGames)
}

func recordState(boardState [][]models.CellState) {
	flattened := flattenBoard(boardState)
	CompletedGames[flattened] = struct{}{}
}

func flattenBoard(boardState [][]models.CellState) string {
	var b strings.Builder

	for _, row := range boardState {
		for _, cell := range row {
			switch cell {
			case models.Empty:
				b.WriteRune('.')
			case models.X:
				b.WriteRune('X')
			case models.O:
				b.WriteRune('O')
			}
		}
		b.WriteRune(';')
	}

	return b.String()
}

func getVisitedState(boardState [][]models.CellState, currentPlayer models.CellState) string {
	return fmt.Sprintf("%s;%d", flattenBoard(boardState), currentPlayer)
}

func isValidBoardPosition(row, col int) bool {
	if col < 0 || col >= BoardWidth {
		return false
	}
	if row < 0 || row >= BoardHeight {
		return false
	}
	return true
}
