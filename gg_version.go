package main

import (
	"fmt"
	"math/rand"
)

// BoardState - структура для представления состояния доски
type BoardState struct {
	board [][]int // Массив массивов, представляющий доску
	turn  int     // Текущий ход (0 - крестики, 1 - нолики)
}

// NewBoardState - функция для создания нового состояния доски
func NewBoardState() *BoardState {
	return &BoardState{
		board: make([][]int, 3),
		turn:  0,
	}
}

// IsGameOver - функция для проверки, закончилась ли игра
func (bs *BoardState) IsGameOver() bool {
	for i := 0; i < 3; i++ {
		if bs.board[i][0]+bs.board[i][1]+bs.board[i][2] == 3 || bs.board[0][i]+bs.board[1][i]+bs.board[2][i] == 3 {
			return true
		}
	}
	if bs.board[0][0]+bs.board[1][1]+bs.board[2][2] == 3 {
		return true
	}
	if bs.board[0][2]+bs.board[1][1]+bs.board[2][0] == 3 {
		return true
	}

	return false
}

// MakeMove - функция для совершения хода
func (bs *BoardState) MakeMove(row, col int) {
	bs.board[row][col] = bs.turn
	bs.turn = 1 - bs.turn
}

// GetAvailableMoves - функция для получения доступных ходов
func (bs *BoardState) GetAvailableMoves() [][]int {
	moves := make([][]int, 0)
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if bs.board[row][col] == 0 {
				moves = append(moves, []int{row, col})
			}
		}
	}
	return moves
}

// Minimax - функция для поиска оптимального хода
func Minimax(bs *BoardState, depth int, isMaximizingPlayer bool) int {
	if bs.IsGameOver() {
		if bs.turn == 0 {
			return 1 // Крестики выиграли
		} else {
			return -1 // Нолики выиграли
		}
	}

	if isMaximizingPlayer {
		bestScore := -1000 // Максимальный проигрыш для максимизирующего игрока
		for _, move := range bs.GetAvailableMoves() {
			bs.MakeMove(move[0], move[1])
			score := Minimax(bs, depth+1, !isMaximizingPlayer)
			bs.MakeMove(-1, -1) // Отменяем ход
			if score > bestScore {
				bestScore = score
			}
		}
		return bestScore
	} else {
		bestScore := 1000 // Максимальный выигрыш для минимизирующего игрока
		for _, move := range bs.GetAvailableMoves() {
			bs.MakeMove(move[0], move[1])
			score := Minimax(bs, depth+1, !isMaximizingPlayer)
			bs.MakeMove(-1, -1) // Отменяем ход
			if score < bestScore {
				bestScore = score
			}
		}
		return bestScore
	}
}

// PlayGame - функция для игры
func PlayGame() {
	bs := NewBoardState()
	for !bs.IsGameOver() {
		fmt.Println("Current board state:")
		for _, row := range bs.board {
			for _, val := range row {
				if val == 0 {
					fmt.Print(".")
				} else if val == 1 {
					fmt.Print("O")
				} else {
					fmt.Print("X")
				}
			}
			fmt.Println()
		}

		if bs.turn == 0 {
			// Крестики делают ход
			moves := bs.GetAvailableMoves()
			if len(moves) == 0 {
				break
			}
			move := moves[Minimax(bs, 0, true)]
			bs.MakeMove(move[0], move[1])
		} else {
			// Нолики делают случайный ход
			var move []int
			for {
				move = bs.GetAvailableMoves()[rand.Intn(len(bs.GetAvailableMoves()))]
				if bs.board[move[0]][move[1]] == 0 {
					break
				}
			}
			bs.MakeMove(move[0], move[1])
		}

	}

	if bs.IsGameOver() && bs.turn == 0 {
		fmt.Println("Крестики выиграли!")
	} else if bs.IsGameOver() {
		fmt.Println("Нолики выиграли!")
	} else {
		fmt.Println("Ничья!")
	}
}

func main() {
	PlayGame()
}
