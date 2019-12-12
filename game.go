package main

import (
    "fmt"
)

type Player Cell

type Game struct {
    player Player
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
				fmt.Print("Input Error\n");
                continue
            }
        }
        g.next()
        g.board.Print()
    }
    g.next()
    fmt.Printf("%v player won\n", g.player.String())
}

func (g *Game) next() {
    switch g.player {
    case Cross:
        g.player = Circle
    case Circle:
        g.player = Cross
    }
}

func (g *Game) setAt(y, x int) error {
    if !inBorder(y, x) || g.board[y][x] != Empty {
        return &BoardIndexError{y, x}
    }
    g.board[y][x] = Cell(g.player)
    return nil
}

func inBorder(y, x int) bool {
    return y >= 0 && y < 3 && x >= 0 && x < 3
}

func (p *Player) String() string {
    switch *p {
    case Cross:
		return "X"
    case Circle:
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
