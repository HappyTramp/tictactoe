package main

import (
	// "sync"
	"fmt"
)

type bestMove struct {
	score int
	move Position
}

func (g *Game) aiPlay() {
	best := g.minimax()
	fmt.Println("final ", best);
	g.setAt(best.move.y, best.move.x)
}

func (g *Game) minimax() bestMove {
	bestMovesChan := make(chan bestMove, 9)
	g.minimax_rec(bestMovesChan)
	return <- bestMovesChan
}

func (g Game) minimax_rec(parentPossibleMovesChan chan<- bestMove) {
	var best bestMove
	possibleMovesChan := make(chan bestMove, 8)
	currentPlayer := g.player
	g.next()
	counter := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board[i][j] != Empty {
				continue
			}
			g.board[i][j] = Cell(currentPlayer)

			// fmt.Println(i, j, g.String(), possibleMovesChan)
			if winner := g.winner(); winner != Empty {
				possibleMovesChan <- bestMove{winner.value(), Position{i, j}}
				counter++
				// fmt.Println("self", i, j, possibleMovesChan)
				g.board[i][j] = Empty
				if winner == currentPlayer {
					break
				}
				continue
			}
			g.minimax_rec(possibleMovesChan)
			counter++
			g.board[i][j] = Empty
		}
	}
	// close(possibleMovesChan)

	best = bestMove{-2 * currentPlayer.value(), Position{-1, -1}}
	// for b := range possibleMovesChan {
	for c := 0; c < counter; c++ {
		b := <- possibleMovesChan
		fmt.Println(b, currentPlayer)
		if (currentPlayer == Circle && b.score > best.score) ||
			(currentPlayer == Cross && b.score < best.score) {
			best = b
		}
	}
	fmt.Println(">>", best)
	parentPossibleMovesChan <- best
	// fmt.Println("parent", parentPossibleMovesChan)
}

func (p *Player) value() int {
	switch *p {
	case Circle: return 1
	case Cross: return -1
	}
	return 0
}
