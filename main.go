package main

import (
	"WorkWhileAssignment/pkg/game"
	"WorkWhileAssignment/pkg/models"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	board, availableSpaces := game.InitializeBoard(3, 3)
	game.Play(models.X, board, availableSpaces)
	duration := time.Since(start)

	fmt.Printf("num of unique completed games: %d\n", game.GetNumUniqueCompletedGames())

	// just out of curiosity
	_ = duration
	//fmt.Printf("duration: %d ms\n", duration.Milliseconds())
}
