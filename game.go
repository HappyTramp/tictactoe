package main

import (
    "fmt"
	"strconv"
)

const (
	Empty = iota
	Cross = iota
	Circle = iota
	NoWinner = iota
)

type Cell byte
type Player Cell
type Game struct {
    player Player
    board [3][3]Cell
}

type Position struct {
	y, x int
}

func (g *Game) Loop() {
    var y, x int

	fmt.Println("tictactoe game, to play insert the index where to place your symbol")
	fmt.Print("ex: [row index] [column index]\n\n")

    fmt.Print(g)
    for g.winner() == Empty {
        for ok := true; ok; ok = g.setAt(y - 1, x - 1) != nil {
            fmt.Print("Enter index: ")
            n, err := fmt.Scanf("%d %d", &y, &x)
            if n != 2 || err != nil {
				fmt.Print("Input Error\n");
                continue
            }
        }
        g.next()
		fmt.Print(g)
		if g.winner() != Empty {
			break
		}
		fmt.Println("AI play")
		g.aiPlay()
        g.next()
		fmt.Print(g)
    }
	if g.winner() == NoWinner {
		fmt.Println("Tie")
	} else {
    	g.next()
		fmt.Printf("%v player won\n", g.player.String())
	}
}

func (g *Game) aiPlay() {
	var move Position
    bestScore := -2
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board[i][j] != Empty {
				continue
			}
			g.board[i][j] = Circle
			g.player = Cross
			score := g.minimax()
			if score > bestScore {
				bestScore = score
				move = Position{i, j}
			}
			g.player = Circle
			g.board[i][j] = Empty
		}
	}
	g.setAt(move.y, move.x)
}

func (g Game) minimax() int {
    switch g.winner() {
	case Cross:  return -1
	case Circle: return 1
	case NoWinner: return 0
    }

	var best int
	switch g.player {
    // mini, human
	case Cross:
        best = 2
        g.player = Circle
        for i := 0; i < 3; i++ {
            for j := 0; j < 3; j++ {
                if g.board[i][j] != Empty {
                    continue
                }
                g.board[i][j] = Cross
				if score := g.minimax(); score < best {
					best = score
				}
                g.board[i][j] = Empty
            }
        }
    	g.player = Cross

    // max, ai
	case Circle:
        best = -2
		g.player = Cross
        for i := 0; i < 3; i++ {
            for j := 0; j < 3; j++ {
                if g.board[i][j] != Empty {
                    continue
                }
                g.board[i][j] = Circle
				if score := g.minimax(); score > best {
					best = score
				}
                g.board[i][j] = Empty
            }
        }
		g.player = Circle
    }
	return best
}

func (g *Game) isTie() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board[i][j] == Empty {
				return false
			}
		}
	}
	return true
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
            return Player(b[i][i])
        }
    }
	if b[1][1] != Empty &&
	   ((b[0][0] == b[1][1] && b[0][0] == b[2][2]) ||
	    (b[0][2] == b[1][1] && b[0][2] == b[2][0])) {
        return Player(b[1][1])
    }
	if g.isTie() {
		return NoWinner
	}
    return Empty
}

func (p *Player) String() string {
	cell := Cell(*p)
    return cell.String()
}

func (g *Game) String() string {
	s := "  1 2 3\n"
    for i, row := range g.board {
        s += strconv.Itoa(i + 1)
        for _, v := range row {
			s += " " + v.String()
        }
		s += "\n"
    }
	return s
}

func (c *Cell) String() string {
    switch *c {
    case Empty:
        return "_"
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
    return fmt.Sprintf("[%v %v] is not a valid index", e.y, e.x)
}
