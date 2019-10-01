package main

import (
    "fmt"
)

type Board [3][3]byte

func (b *Board) CheckWin() bool {
    var rowCheck, colCheck, diagCheck, antiDiagCheck bool

    for i := 0; i < 3; i++ {
        rowCheck = true
        colCheck = true
        for j := 0; j < 3; j++ {
            // fmt.Println(b[i][j], Empty)
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
    diagCheck, antiDiagCheck = true, true
    for i, j := 0, 0; i < 3; i, j = i + 1, j + 1 {
        if b[i][j] == Empty || b[i][j] != b[(i + 1) % 3][(j + 1) % 3] {
            diagCheck = false
        }
        if b[i][2 - j] == Empty ||
        b[i][2 - j] != b[(i + 1) % 3][positiveMod(2 - j - 1, 3)] {
            antiDiagCheck = false
        }
    }
    return diagCheck || antiDiagCheck
}

func (b *Board) Print() {
    for _, row := range b {
        for _, v := range row {
            switch v {
            case 0:
                fmt.Print("_ ")
            case 1:
                fmt.Print("O ")
            case 2:
                fmt.Print("X ")
            }
        }
        fmt.Print("\n")
    }
}

// Positive modulo, returns non negative solution to x % d
func positiveMod(a, b int) int {
    m := a % b
    if a < 0 && b < 0 {
        m -= b
    }
    if a < 0 && b > 0 {
        m += b
    }
    return m
}
