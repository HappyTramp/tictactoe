package main

import (
    "fmt"
)

type Cell byte

const (
	Empty = 0
	Cross = 1
	Circle = 2
)

type Board [3][3]Cell

func (b *Board) CheckWin() bool {
    for i := 0; i < 3; i++ {
		rowCheck := true
		colCheck := true
        for j := 0; j < 3; j++ {
            if b[i][j] == Empty || b[i][j] != b[i][(j + 1) % 3] {
                rowCheck = false
            }
            if b[j][i] == Empty || b[j][i] != b[(j + 1) % 3][i] {
                colCheck = false
            }
        }
        if rowCheck || colCheck {
            return true
        }
    }
	return b[1][1] != Empty &&
	       ((b[0][0] == b[1][1] && b[0][0] == b[2][2]) ||
	        (b[0][2] == b[1][1] && b[0][2] == b[2][0]))
}

func (b *Board) Print() {
    for _, row := range b {
        for _, v := range row {
            switch v {
            case Empty:
				fmt.Print("_ ")
            case Cross:
				fmt.Print("X ")
            case Circle:
				fmt.Print("O ")
            }
        }
        fmt.Print("\n")
    }
}

// func (b Board) Minimax() {

// }
