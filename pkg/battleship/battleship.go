// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package battleship

import (
    "fmt"
    "container/vector"
)

type Board struct {
    squares      [][]Square // [row][column] order
    rowClues     []int
    columnClues  []int
    ships        []int
    turn_counter int // just for fun, number of attempts
}

func (board *Board) String() string {
    var s = ""
    for rowIndex, rowClue := range board.rowClues {
        row := board.squares[rowIndex]
        for _, square := range row {
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

func (board *Board) NumberOfRows() int {
    return len(board.rowClues)
}

func (board *Board) NumberOfColumns() int {
    return len(board.columnClues)
}

func (board *Board) GetSquareAt(coord Coord) Square {
    row := coord.Row()
    column := coord.Column()
    if (column < 0) ||
        (column >= board.NumberOfColumns()) ||
        (row < 0) ||
        (row >= board.NumberOfRows()) {
        return OUT_OF_BOUNDS
    }
    return board.squares[row][column]
}

func (board Board) SetSquareAt(coord Coord, square Square) {
    board.squares[coord.Row()][coord.Column()] = square
}

func (board *Board) GetCoordOfUnsolvedSquare() (Coord, bool) {
    for rowIndex, row := range board.squares {
        for columnIndex, square := range row {
            if square.IsUnsolved() {
                //fmt.Printf("%v,%v\n",x,y)
                return NewCoord(rowIndex, columnIndex), true
            }
        }
    }
    return 0, false
}

func (board *Board) IsValid() bool {
    return true
}


func (board *Board) ShipCountInRow(row int) int {
    var count = 0
    for i := 0; i < board.NumberOfColumns(); i++ {
        if board.squares[row][i].IsShip() {
            count++
        }
    }
    return count
}

func (board *Board) UnsolvedCountInRow(row int) int {
    var count = 0
    for i := 0; i < board.NumberOfColumns(); i++ {
        if board.squares[row][i].IsUnsolved() {
            count++
        }
    }
    return count
}

func (board *Board) ShipCountInColumn(column int) int {
    var count = 0
    for i := 0; i < board.NumberOfRows(); i++ {
        if board.squares[i][column].IsShip() {
            count++
        }
    }
    return count
}

func (board *Board) UnsolvedCountInColumn(column int) int {
    var count = 0
    for i := 0; i < board.NumberOfRows(); i++ {
        if board.squares[i][column].IsUnsolved() {
            count++
        }
    }
    return count
}

func (board *Board) CalcPossibleSquaresFor(coord Coord) *vector.IntVector {
    var possibilities vector.IntVector

    var requireWater = false
    var requireShip = false

    desired := board.rowClues[coord.Row()]
    actual := board.ShipCountInRow(coord.Row())
    unsolved := board.UnsolvedCountInRow(coord.Row())
    switch {
    case actual > desired:
        // TODO: Should this be a panic?
        fmt.Printf("Unsolvable:  too many ships in row at coord %v", coord)
        return &possibilities
    case actual == desired:
        requireWater = true
    case actual+unsolved == desired:
        requireShip = true
    case actual+unsolved < desired:
        // TODO: Should this be a panic?
        fmt.Printf("Unsolvable:  too few ships in row at coord %v", coord)
        return &possibilities
    }

    desired = board.columnClues[coord.Column()]
    actual = board.ShipCountInColumn(coord.Column())
    unsolved = board.UnsolvedCountInColumn(coord.Column())
    switch {
    case actual > desired:
        // TODO: Should this be a panic?
        fmt.Printf("Unsolvable:  too many ships in column at coord %v", coord)
        return &possibilities
    case actual == desired:
        requireWater = true
    case actual+unsolved == desired:
        requireShip = true
    case actual+unsolved < desired:
        // TODO: Should this be a panic?
        fmt.Printf("Unsolvable:  too few ships in column at coord %v", coord)
        return &possibilities
    }

    for _, square := range SQUARES {
        switch {
        case requireWater && !square.IsWater():
        case requireShip && !square.IsShip():
        case !board.GetSquareAt(coord.Above()).CanAppearAbove(square):
        case !board.GetSquareAt(coord.Below()).CanAppearBelow(square):
        case !board.GetSquareAt(coord.Right()).CanAppearRightOf(square):
        case !board.GetSquareAt(coord.Left()).CanAppearLeftOf(square):
        case !board.GetSquareAt(coord.Above().Left()).CanAppearDiagonallyAdjacentTo(square):
        case !board.GetSquareAt(coord.Above().Right()).CanAppearDiagonallyAdjacentTo(square):
        case !board.GetSquareAt(coord.Below().Left()).CanAppearDiagonallyAdjacentTo(square):
        case !board.GetSquareAt(coord.Below().Right()).CanAppearDiagonallyAdjacentTo(square):
        default:
            possibilities.Push(int(square))
        }
    }
    return &possibilities
}

func (board *Board) Solve() bool {
    coord, ok := board.GetCoordOfUnsolvedSquare()
    if !ok {
        return true // all solved!
    }

    var possibilities = board.CalcPossibleSquaresFor(coord)

    for _, possibility := range *possibilities {
        square := Square(possibility)
        board.turn_counter++
        //if board.turn_counter % 1000 == 0 {
        //    fmt.Printf("At turn %v possibilities for %v are %v, trying %v (%v)\n%v\n",
        //              board.turn_counter, coord, possibilities,
        //              possibility, square, board)
        //}
        board.SetSquareAt(coord, square)
        if board.IsValid() && board.Solve() {
            return true
        }
        board.SetSquareAt(coord, UNSOLVED)
    }
    return false
}

func NewBoard(rowClues []int, columnClues []int, ships []int, initialSquares []struct {
    row    int
    column int
    square Square
}) *Board {
    squares := make([][]Square, len(rowClues))
    for rowIndex, _ := range squares {
        row := make([]Square, len(columnClues))
        squares[rowIndex] = row
        for columnIndex, _ := range row {
            row[columnIndex] = UNSOLVED
        }
    }
    for _, square := range initialSquares {
        squares[square.row][square.column] = square.square
    }
    return &Board{squares: squares, columnClues: columnClues, rowClues: rowClues, ships: ships}
}
