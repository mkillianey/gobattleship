package battleship

import (
    //"fmt"
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
    var numPossibilities = 0

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
        case square == MIDDLE &&
            !((squareAbove.CouldBeShip() && squareBelow.CouldBeShip()) ||
              (squareLeft.CouldBeShip() && squareRight.CouldBeShip())):
        default:
            possibilities[numPossibilities] = square
            numPossibilities++
        }
    }
    return possibilities[:numPossibilities]
}


// SquareSolver makes a copy of the board and fills in all forced
// squares, then guesses what to fill in for the most constrained
// unsolved square and calls itself recursively to see if it can
// find a solution.
type SquareSolver struct{}

func (solver *SquareSolver) solve(board *Board) (*Board, bool) {
    workingCopy := board.Copy()
    //fmt.Printf("Solving:\n%v\n", workingCopy)

    var foundUnsolved = false
    var minCoord Coord
    var minCount int
    for coord, square := range workingCopy.squares {
        if square.IsUnsolved() {
            candidates := calcPossibleSquares(workingCopy, coord)
            count := len(candidates)
            switch count {
            case 0:
                //fmt.Printf("%v is unsolved: No solution\n", coord)
                return nil, false
            case 1:
                //fmt.Printf("%v is unsolved, but can only be %v\n", coord, candidates)
                workingCopy.SetSquareAt(coord, Square(candidates[0]))
            default:
                //fmt.Printf("%v is unsolved: %v\n", coord, candidates)
                if !foundUnsolved || minCount > len(candidates) {
                    foundUnsolved = true
                    minCoord = coord
                    minCount = len(candidates)
                }
            }
        }
    }
    if !workingCopy.IsValid() {
        //fmt.Printf("Something went wrong...this board is now invalid\n%v\n", workingCopy);
        return nil, false
    }
    if !foundUnsolved {
        return workingCopy, true
    }
    for _, square := range calcPossibleSquares(workingCopy, minCoord) {
        workingCopy.SetSquareAt(minCoord, square)
        //fmt.Printf("Trying %v at %v\n", square, minCoord)
        if solution, ok := solver.solve(workingCopy); ok {
            return solution, true
        }
        //fmt.Printf("Undoing %v at %v\n", square, minCoord)
    }
    return nil, false
}

func (solver *SquareSolver) SolveClues(clues *Clues) (solution *Board, ok bool) {
    solution, ok = solver.solve(NewBoard(clues))
    return solution, ok
}

func NewSolver() Solver {
    return &SquareSolver{}
}

