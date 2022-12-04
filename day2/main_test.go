package main

import (
	"testing"
)

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		playerMoves, opponentMoves := parser()
		playerMovesInt := convertToInt(playerMoves)
		opponentMovesInt := convertToInt(opponentMoves)
		calculateScore(playerMovesInt, opponentMovesInt)
	}
}