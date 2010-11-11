// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package battleship

import (
    "fmt"
    "container/vector"
)


// The immutable clues that, once set, never change
type Clues struct {
    rowClues           []int
    columnClues        []int
    ships              []int
    initialSquares     map[Coord]Square
}

func (clues *Clues) NumberOfRows() int {
    return len(clues.rowClues)
}

func (clues *Clues) NumberOfColumns() int {
    return len(clues.columnClues)
}


// The mutable parts of the board
type Board struct {
    *Clues
    squares            map[Coord]Square
    logging            bool // turn on to see what's happening
    numCalls_to_SetSquareAt int
    numCalls_to_GetSquareAt int
    numCalls_to_NextCoordToSolve int
    numCalls_to_CalcPossibleSquaresFor int
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

// Set to true to active logging for actions
func (board *Board) SetLogging(logging bool) {
    board.logging = logging
}

func (board *Board) GetSquareAt(coord Coord) Square {
    board.numCalls_to_GetSquareAt++
    if square, ok := board.squares[coord]; ok {
        return square
    }
    if square, ok := board.initialSquares[coord]; ok {
        return square
    }
    return OUT_OF_BOUNDS
}

func (board *Board) SetSquareAt(coord Coord, square Square) {
    board.numCalls_to_SetSquareAt++
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
                    if board.logging {
                        fmt.Printf("Middle ship at %v can't be vertical or horizontal\n", coord)
                    }
                    return false
                }
            }
        }
    }
    return true
}


// Returns the minimum and maximum number of ships in this
// row.  The minimum is the number of ships *currently*
// placed in this row.  The maximum is the number of ships
// that could be placed if all unsolved squares were filled
// with ships.
func (board *Board) MinMaxShipsInRow(coord Coord) (int, int) {
    var ships, unsolveds = 0, 0
    for column := 0; column < board.NumberOfColumns(); column++ {
        square := board.GetSquareAt(coord.WithColumn(column))
        switch {
        case square.IsShip():
            ships++
        case square.IsUnsolved():
            unsolveds++
        }
    }
    return ships, ships + unsolveds
}

// Returns the minimum and maximum number of ships in this
// column.  The minimum is the number of ships *currently*
// placed in this column.  The maximum is the number of ships
// that could be placed if all unsolved squares were filled
// with ships.
func (board *Board) MinMaxShipsInColumn(coord Coord) (int, int) {
    var ships, unsolveds = 0, 0
    for row := 0; row < board.NumberOfRows(); row++ {
        square := board.GetSquareAt(coord.WithRow(row))
        switch {
        case square.IsShip():
            ships++
        case square.IsUnsolved():
            unsolveds++
        }
    }
    return ships, ships + unsolveds
}

func (board *Board) CalcPossibleSquaresFor(coord Coord) *vector.IntVector {
    board.numCalls_to_CalcPossibleSquaresFor++
    var possibilities vector.IntVector

    var requireWater = false
    var requireShip = false

    desired := board.rowClues[coord.Row()]
    minships, maxships := board.MinMaxShipsInRow(coord)
    switch {
    case minships > desired:
        // TODO: Should this be a panic?
        fmt.Printf("Unsolvable:  too many ships in row at coord %v", coord)
        return &possibilities
    case minships == desired:
        requireWater = true
    case maxships == desired:
        requireShip = true
    case maxships < desired:
        // TODO: Should this be a panic?
        fmt.Printf("Unsolvable:  too few ships in row at coord %v", coord)
        return &possibilities
    }

    desired = board.columnClues[coord.Column()]
    minships, maxships = board.MinMaxShipsInColumn(coord)
    switch {
    case minships > desired:
        // TODO: Should this be a panic?
        fmt.Printf("Unsolvable:  too many ships in column at coord %v", coord)
        return &possibilities
    case minships == desired:
        requireWater = true
    case maxships == desired:
        requireShip = true
    case maxships < desired:
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

// Finds the next coord on the board to solve, choosing one
// of the coordinates with the fewest possibilities
func (board *Board) NextCoordToSolve() (Coord, bool) {
    board.numCalls_to_NextCoordToSolve++
    var found = false
    var minCoord Coord
    var minCount int
    for coord, square := range board.squares {
        if square.IsUnsolved() {
            possibleSquares := board.CalcPossibleSquaresFor(coord)
            if possibleSquares.Len() <= 1 {
                return coord, true
            }
            if (!found) || (possibleSquares.Len() < minCount) {
                found = true
                minCoord = coord
                minCount = possibleSquares.Len()
            }
        }
    }
    return minCoord, found
}

func (board *Board) Solve() bool {
    coord, foundOne := board.NextCoordToSolve()
    if !foundOne {
        if board.logging {
            fmt.Printf("Solved with:\n");
            fmt.Printf("%v calls to SetSquareAt\n", board.numCalls_to_SetSquareAt);
            fmt.Printf("%v calls to GetSquareAt\n", board.numCalls_to_GetSquareAt);
            fmt.Printf("%v calls to NextCoordToSolve\n", board.numCalls_to_NextCoordToSolve);
            fmt.Printf("%v calls to CalcPossibleSquaresFor\n", board.numCalls_to_CalcPossibleSquaresFor);
        }
        return true // solved!
    }
    possibleSquares := board.CalcPossibleSquaresFor(coord)
    
    if possibleSquares.Len() == 0 {
        if board.logging {
            fmt.Printf("No possibilities for %v\n", coord)
        }
        return false
    }
    if board.logging {
        fmt.Printf("Possible squares for %v are %v\n", coord, possibleSquares)
    }
    for _, possibleSquare := range *possibleSquares {
        square := Square(possibleSquare)
        board.SetSquareAt(coord, square)
        if board.logging {
            fmt.Printf("Placing %v at %v:\n%v", square, coord, board)
        }
        if board.IsValid() && board.Solve() {
            return true
        }
        board.SetSquareAt(coord, UNSOLVED)
        if board.logging {
            fmt.Printf("Undoing %v at %v\n", square, coord)
        }
    }
    if board.logging {
        fmt.Printf("Can't place anything at %v...backtracking:\n%v\n", coord, board)
        return false
    }
    return false
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
        Clues: clues,
        squares: squares}
}
