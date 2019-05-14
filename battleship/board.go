// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package battleship

import (
    "fmt"
)


// The immutable clues that, once set, never change
type Clues struct {
    title          string
    rowClues       []int
    columnClues    []int
    shipLengths    []int
    initialSquares map[Coord]Square
    solution       string
}

func (clues *Clues) String() string {
    return fmt.Sprintf("%v: ships=%v\n%v", clues.title, clues.shipLengths, NewBoard(clues))
}

func (clues *Clues) Title() string {
    return clues.title
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

func (clues *Clues) Solution() string {
    return clues.solution
}


// The mutable parts of the board
type Board struct {
    *Clues
    squares map[Coord]Square
}

func (board *Board) Copy() *Board {
    squares := make(map[Coord]Square, len(board.squares))
    for coord, square := range board.squares {
        squares[coord] = square
    }
    return &Board{Clues: board.Clues, squares: squares}
}

func (board *Board) String() string {
    var s = ""
    for row := 0; row < board.NumberOfRows(); row++ {
        for column := 0; column < board.NumberOfColumns(); column++ {
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
        s += fmt.Sprintf("%v\n", board.RowClue(row))
    }
    for column := 0; column < board.NumberOfColumns(); column++ {
        s += fmt.Sprintf("%v", board.ColumnClue(column))
    }
    s += "\n"
    return s
}

func (board *Board) SquareAt(coord Coord) Square {
    if square, ok := board.squares[coord]; ok {
        return square
    }
    return OUT_OF_BOUNDS
}

func (board *Board) SetSquareAt(coord Coord, square Square) {
    board.squares[coord] = square
}

func (board *Board) IsValid() bool {
    // Make sure each middle could either be a vertical or horizontal ship
    for coord, square := range board.squares {
        if square == MIDDLE {
            above := board.SquareAt(coord.Above())
            below := board.SquareAt(coord.Below())
            if !(above.CouldBeShip() && below.CouldBeShip()) {
                left := board.SquareAt(coord.Left())
                right := board.SquareAt(coord.Right())
                if !(left.CouldBeShip() && right.CouldBeShip()) {
                    return false
                }
            }
        }
    }

    // Make sure we don't find more ships than we have
    shipsFound := make([]bool, board.NumberOfShips())
    for coord, square := range board.squares {
        found, length := false, 0 // measuring from top-left of ship
        switch square {
        case SINGLE:
            found, length = true, 1
        case TOP:
            candidate := 1
            nextCoord := coord.Below()
            for board.SquareAt(nextCoord) == MIDDLE {
                candidate++
                nextCoord = nextCoord.Below()
            }
            if board.SquareAt(nextCoord) == BOTTOM {
                found, length = true, candidate+1
            }
        case LEFT:
            candidate := 1
            nextCoord := coord.Right()
            for board.SquareAt(nextCoord) == MIDDLE {
                candidate++
                nextCoord = nextCoord.Right()
            }
            if board.SquareAt(nextCoord) == RIGHT {
                found, length = true, candidate+1
            }
        }
        if found {
            markedShipAsFound := false
            for index, foundYet := range shipsFound {
                if (!foundYet) && (board.ShipLength(index) == length) {
                    shipsFound[index] = true
                    markedShipAsFound = true
                    break
                }
            }
            if !markedShipAsFound {
                return false // too many ships of this length
            }
        }
    }

    // Make sure each row/column doesn't have too few/many ships
    shipsInRow := make([]int, board.NumberOfRows())
    shipsInColumn := make([]int, board.NumberOfColumns())
    watersInRow := make([]int, board.NumberOfRows())
    watersInColumn := make([]int, board.NumberOfColumns())
    for row := 0; row < board.NumberOfRows(); row++ {
        shipsInRow[row] = 0
        watersInRow[row] = 0
    }
    for column := 0; column < board.NumberOfColumns(); column++ {
        shipsInColumn[column] = 0
        watersInColumn[column] = 0
    }
    for coord, square := range board.squares {
        switch {
        case square.IsWater():
            watersInRow[coord.Row()] += 1
            watersInColumn[coord.Column()] += 1
        case square.IsShip():
            shipsInRow[coord.Row()] += 1
            shipsInColumn[coord.Column()] += 1
        }
    }
    for row := 0; row < board.NumberOfRows(); row++ {
        clue := board.RowClue(row)
        switch {
        case shipsInRow[row] > clue:
            return false
        case board.NumberOfColumns()-watersInRow[row] < clue:
            return false
        }
    }
    for column := 0; column < board.NumberOfColumns(); column++ {
        clue := board.ColumnClue(column)
        switch {
        case shipsInColumn[column] > clue:
            return false
        case board.NumberOfRows()-watersInColumn[column] < clue:
            return false
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
