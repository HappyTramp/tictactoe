package main

import (
    "fmt"
)

type Player byte
const (
    Empty = 0
    crossPlayer = 1
    circlePlayer = 2
)

type Game struct {
    currentPlayer Player
    board Board
}

func (g *Game) Loop() {
    var y, x int

    g.board.Print()
    for !g.board.CheckWin() {
        for keep := true; keep; keep = g.setAt(y, x) != nil {
            fmt.Print("\nEnter index: ")
            n, err := fmt.Scanf("%d %d", &y, &x)
            if n != 2 || err != nil {
                continue
            }
        }
        g.next()
        g.board.Print()
    }
    fmt.Printf("%v player won\n", g.currentPlayer.String())
}

func (g *Game) next() {
    switch g.currentPlayer {
    case crossPlayer:
        g.currentPlayer = circlePlayer
    case circlePlayer:
        g.currentPlayer = crossPlayer
    }
}

func (g *Game) setAt(y, x int) error {
    if !inBorder(y, x) || g.board[y][x] != Empty {
        return &BoardIndexError{y, x}
    }
    g.board[y][x] = byte(g.currentPlayer)
    return nil
}

func inBorder(y, x int) bool {
    return y >= 0 && y < 3 && x >= 0 && x < 3
}

func (p *Player) String() string {
    switch *p {
    case crossPlayer:
        return "X"
    case circlePlayer:
        return "O"
    }
    return " "
}

type BoardIndexError struct {
    y, x int
}

func (e *BoardIndexError) Error() string {
    return fmt.Sprintf("%v %v is not a valid index", e.y, e.x)
}
