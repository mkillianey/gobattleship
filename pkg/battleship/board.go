// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package battleship

import (
    "fmt"
)


// The immutable clues that, once set, never change
type Clues struct {
    rowClues       []int
    columnClues    []int
    shipLengths    []int
    initialSquares map[Coord]Square
}

func (clues *Clues) NumberOfRows() int {
    return len(clues.rowClues)
}

func (clues *Clues) NumberOfColumns() int {
    return len(clues.columnClues)
}

func (clues *Clues) NumberOfShips() int {
    return len(clues.shipLengths)
}

func (clues *Clues) RowClue(row int) int {
    return clues.rowClues[row]
}

func (clues *Clues) ColumnClue(column int) int {
    return clues.columnClues[column]
}

func (clues *Clues) ShipLength(index int) int {
    return clues.shipLengths[index]
}


// The mutable parts of the board
type Board struct {
    *Clues
    squares map[Coord]Square
}

func (board *Board) String() string {
    var s = ""
    for row, rowClue := range board.rowClues {
        for column, _ := range board.columnClues {
            square := board.squares[NewCoord(row, column)]
            switch square {
            case UNSOLVED:
                s += "."
            case WATER:
                s += "~"
            case TOP:
                s += "^"
            case BOTTOM:
                s += "v"
            case LEFT:
                s += "<"
            case RIGHT:
                s += ">"
            case MIDDLE:
                s += "#"
            case SINGLE:
                s += "O"
            default:
                s += "?"
            }
        }
        s += fmt.Sprintf("%v\n", rowClue)
    }
    for _, columnClue := range board.columnClues {
        s += fmt.Sprintf("%v", columnClue)
    }
    s += "\n"
    return s
}

func (board *Board) GetSquareAt(coord Coord) Square {
    if square, ok := board.squares[coord]; ok {
        return square
    }
    if square, ok := board.initialSquares[coord]; ok {
        return square
    }
    return OUT_OF_BOUNDS
}

func (board *Board) SetSquareAt(coord Coord, square Square) {
    board.squares[coord] = square
}


func (board *Board) IsValid() bool {
    for coord, square := range board.squares {
        if square == MIDDLE {
            above := board.GetSquareAt(coord.Above())
            below := board.GetSquareAt(coord.Below())
            canBeVertical := (above.IsShip() || above.IsUnsolved()) &&
                (below.IsShip() || below.IsUnsolved())
            if !canBeVertical {
                left := board.GetSquareAt(coord.Left())
                right := board.GetSquareAt(coord.Right())
                canBeHorizontal := (left.IsShip() || left.IsUnsolved()) &&
                    (right.IsShip() || right.IsUnsolved())
                if !canBeHorizontal {
                    return false
                }
            }
        }
    }
    return true
}


func NewBoard(clues *Clues) *Board {
    squares := make(map[Coord]Square, len(clues.rowClues)*len(clues.columnClues))
    for row, _ := range clues.rowClues {
        for column, _ := range clues.columnClues {
            coord := NewCoord(row, column)
            if square, ok := clues.initialSquares[coord]; ok {
                squares[coord] = square
            } else {
                squares[coord] = UNSOLVED
            }
        }
    }
    return &Board{
        Clues:   clues,
        squares: squares}
}
