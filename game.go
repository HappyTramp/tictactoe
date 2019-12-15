package main

import (
    "fmt"
	// "math/rand"
)

const (
	Empty = iota
	Cross = iota
	Circle = iota
)

type Cell byte
type Board [3][3]Cell

type Player Cell
type Game struct {
    player Player
    board Board
}

type Position struct {
	y, x int
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
		if g.board.CheckWin() {
			break
		}
		g.aiPlay()
        g.next()
        g.board.Print()
    }
    g.next()
    fmt.Printf("%v player won\n", g.player.String())
}

func (g *Game) aiPlay() {
	// var possible [9]Position;
	// possibleIndex := 0;

    best := MinInt32
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (g.board[i][j] == Empty) {

                g.board[i][j] = Circle
                g.player = Cross
                best = max(best, minimax(g))
                g.player = Circle
                g.board[i][j] = Empty
				// possible[possibleIndex] = Position{i, j};
				// possibleIndex++;
			}
		}
	}
	// picked := possible[rand.Intn(possibleIndex)]
	// g.setAt(picked.y, picked.x)
}

func min(x, y int) int {
    if x < y {
        return x
    } else {
        return y
    }
}

func max(x, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}

func (g Game) minimax() int {
    switch g.CheckWin() {
        case Cross:
            return 1
        case Circle:
            return -1
    }
    // mini
    if (g.player == Cross) {
        best := MaxInt32
        for i := 0; i < 3; i++ {
            for j := 0; j < 3; j++ {
                if g.board[i][j] != Empty {
                    continue
                }
                g.board[i][j] = Cross
                g.player = Circle
                best = min(best, minimax(g))
                g.player = Cross
                g.board[i][j] = Empty
            }
        }
        return best
    }
    // max
    else if g.player == Circle {
        best := MinInt32
        for i := 0; i < 3; i++ {
            for j := 0; j < 3; j++ {
                if g.board[i][j] != Empty {
                    continue
                }
                g.board[i][j] = Circle
                g.player = Cross
                best = max(best, minimax(g))
                g.player = Circle
                g.board[i][j] = Empty
            }
        }
        return best
    }
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
    return Cell(p).String()
}

func (g *Game) winner() Player {
    b := g.board
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
            return b[i][i]
        }
    }
	if b[1][1] != Empty &&
	   ((b[0][0] == b[1][1] && b[0][0] == b[2][2]) ||
	    (b[0][2] == b[1][1] && b[0][2] == b[2][0])) {
        return b[1][1]
    }
    return nil
}

func (g *Game) String() string {
    fmt.Println("1 2 3")
    for i, row := range b {
        fmt.Printf("%v ", i + 1)
        for _, v := range row {
            fmt.Print(v)
        }
        fmt.Print("\n")
    }
}

func (c *Cell) String() string {
    switch c {
    case Empty:
        return "_ "
    case Cross:
        return "X "
    case Circle:
        return "O "
}

type BoardIndexError struct {
    y, x int
}

func (e *BoardIndexError) Error() string {
    return fmt.Sprintf("[%v %v] is not a valid index", e.y, e.x)
}