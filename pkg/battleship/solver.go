package battleship

import (
//    "fmt"
)


type Solver interface {
    SolveClues(clues *Clues) (solution *Board, ok bool)
}

// Returns the minimum and maximum number of ships in this
// row.  The minimum is the number of ships *currently*
// placed in this row.  The maximum is the number of ships
// that could be placed if all unsolved squares were filled
// with ships.
func minMaxShipsInRow(board *Board, coord Coord) (int, int) {
    var ships, unsolveds = 0, 0
    for column := 0; column < board.NumberOfColumns(); column++ {
        square := board.SquareAt(coord.WithColumn(column))
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
func minMaxShipsInColumn(board *Board, coord Coord) (int, int) {
    var ships, unsolveds = 0, 0
    for row := 0; row < board.NumberOfRows(); row++ {
        square := board.SquareAt(coord.WithRow(row))
        switch {
        case square.IsShip():
            ships++
        case square.IsUnsolved():
            unsolveds++
        }
    }
    return ships, ships + unsolveds
}


var NO_SOLUTIONS_FOR_SQUARE []Square = []Square{}
var SQUARE_MUST_BE_WATER []Square = []Square{WATER}

func calcPossibleSquares(board *Board, coord Coord) []Square {
    var requireShip = false
    var requireWater = false
    desired := board.rowClues[coord.Row()]
    alreadyPlacedShips, maxPossibleShips := minMaxShipsInRow(board, coord)
    switch {
    case alreadyPlacedShips == desired:
        requireWater = true
    case maxPossibleShips == desired:
        requireShip = true
    case alreadyPlacedShips > desired:
        // TODO: Should this be a panic?
        //fmt.Printf("Unsolvable:  too many ships in row at coord %v", coord)
        return NO_SOLUTIONS_FOR_SQUARE
    case maxPossibleShips < desired:
        // TODO: Should this be a panic?
        //fmt.Printf("Unsolvable:  too few ships in row at coord %v", coord)
        return NO_SOLUTIONS_FOR_SQUARE
    }

    desired = board.columnClues[coord.Column()]
    alreadyPlacedShips, maxPossibleShips = minMaxShipsInColumn(board, coord)
    switch {
    case alreadyPlacedShips == desired:
        requireWater = true
    case maxPossibleShips == desired:
        requireShip = true
    case alreadyPlacedShips > desired:
        // TODO: Should this be a panic?
        //fmt.Printf("Unsolvable:  too many ships in column at coord %v", coord)
        return NO_SOLUTIONS_FOR_SQUARE
    case maxPossibleShips < desired:
        // TODO: Should this be a panic?
        //fmt.Printf("Unsolvable:  too few ships in column at coord %v", coord)
        return NO_SOLUTIONS_FOR_SQUARE
    }

    if requireShip && requireWater {
        return NO_SOLUTIONS_FOR_SQUARE
    }

    var possibilities []Square = make([]Square, len(SQUARES))
    var possibilityIndex = 0

    var above = coord.Above()
    var below = coord.Below()
    var left = coord.Left()
    var right = coord.Right()

    var squareAbove = board.SquareAt(above)
    var squareBelow = board.SquareAt(below)
    var squareLeft = board.SquareAt(left)
    var squareRight = board.SquareAt(right)

    for _, square := range SQUARES {
        switch {
        case requireWater && !square.IsWater():
        case requireShip && !square.IsShip():
        case !squareAbove.CanAppearAbove(square):
        case !squareBelow.CanAppearBelow(square):
        case !squareLeft.CanAppearLeftOf(square):
        case !squareRight.CanAppearRightOf(square):
        case square.IsShip() &&
            (board.SquareAt(above.Left()).IsShip() ||
                board.SquareAt(above.Right()).IsShip() ||
                board.SquareAt(below.Left()).IsShip() ||
                board.SquareAt(below.Right()).IsShip()):
            requireWater = true
        default:
            possibilities[possibilityIndex] = square
            possibilityIndex++
        }
    }
    return possibilities[:possibilityIndex]
}


// Finds the next coord on the board to solve, choosing one
// of the coordinates with the fewest possibilities
func nextCoordToSolve(board *Board) (Coord, bool) {
    var found = false
    var minCoord Coord
    var minCount int
    for coord, square := range board.squares {
        if square.IsUnsolved() {
            possibleSquares := calcPossibleSquares(board, coord)
            if len(possibleSquares) <= 1 {
                return coord, true
            }
            if (!found) || (len(possibleSquares) < minCount) {
                found = true
                minCoord = coord
                minCount = len(possibleSquares)
            }
        }
    }
    return minCoord, found
}

type BasicSolver struct{}

func (solver *BasicSolver) recursiveSolve(board *Board) (*Board, bool) {
    if !board.IsValid() {
        return nil, false
    }
    coord, foundOne := nextCoordToSolve(board)
    if !foundOne {
        return board, true // solved!
    }
    possibleSquares := calcPossibleSquares(board, coord)

    for _, possibleSquare := range possibleSquares {
        square := Square(possibleSquare)
        board.SetSquareAt(coord, square)
        //fmt.Printf("Trying %v at %v\n%v\n", square, coord, board)
        if solution, ok := solver.recursiveSolve(board); ok {
            return solution, ok
        }
        //fmt.Printf("Undoing %v at %v.\n", square, coord)
    }
    board.SetSquareAt(coord, UNSOLVED)
    //fmt.Printf("Backtracking from coord %v:\n", coord)
    return nil, false
}

func (solver *BasicSolver) SolveClues(clues *Clues) (solution *Board, ok bool) {
    solution, ok = solver.recursiveSolve(NewBoard(clues))
    return solution, ok
}

func NewSolver() Solver {
    solver := BasicSolver{}
    return &solver
}
