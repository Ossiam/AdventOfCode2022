package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	rockScore int = 1
	paperScore int = 2
	scissorsScore int = 3
)

const (
	winScore int = 6
	loseScore int = 0
	drawScore int = 3
)

const (
	toLose int = rockScore
	toWin int = scissorsScore
	toDraw int = paperScore
)

func convertToInt(moves []string) []int {
	var intMoves []int
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
			case "A", "X":
				intMoves = append(intMoves, rockScore)
			case "B", "Y":
				intMoves = append(intMoves, paperScore)
			case "C", "Z":
				intMoves = append(intMoves, scissorsScore)
		}
	}
	return intMoves
}

func compareMoves1(playerMove int, opponentMove int) int {
	
	if playerMove == opponentMove {
		return drawScore
	} else if 
		playerMove == rockScore && opponentMove == scissorsScore || 
		playerMove == paperScore && opponentMove == rockScore || 
		playerMove == scissorsScore && opponentMove == paperScore {
		return winScore
	}
	return loseScore
}

func compareMoves2(opponentMove int, playerMove int) int {
	var score int

	switch playerMove {
		case toLose:
			score += loseScore
			if opponentMove == rockScore {
				score += scissorsScore
			}else {
				score += opponentMove - 1
			}
		case toWin:
			score += winScore
			if opponentMove == scissorsScore {
				score += rockScore
			}else{
				score += opponentMove + 1
			}
		case toDraw:
			return opponentMove + drawScore
	}
	return score
}

func calculateScore(playerMoves []int, opponentMoves []int) (int, int) {
	var scoreStep1 int
	var scoreStep2 int
	for i := 0; i < len(playerMoves); i++ {
		scoreStep1 += playerMoves[i] + compareMoves1(playerMoves[i], opponentMoves[i])
		scoreStep2 += compareMoves2(opponentMoves[i], playerMoves[i])
	}
	return scoreStep1, scoreStep2
}

func parser() ([]string, []string) {
	var playerMoves []string
	var opponentMoves []string
	f, err := os.Open("input_day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		opponentMoves = append(opponentMoves, scanner.Text()[0:1])
		playerMoves = append(playerMoves, scanner.Text()[2:3])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return playerMoves, opponentMoves
}

func main() {
	playerMoves, opponentMoves := parser()
	playerMovesInt := convertToInt(playerMoves)
	opponentMovesInt := convertToInt(opponentMoves)
	scoreStep1, scoreStep2 := calculateScore(playerMovesInt, opponentMovesInt)
	fmt.Println("Score step 1: ", scoreStep1)
	fmt.Println("Score step 2: ", scoreStep2)
}