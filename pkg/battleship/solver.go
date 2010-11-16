package battleship

import (
    //    "fmt"
    "container/vector"
)


// Returns the minimum and maximum number of ships in this
// row.  The minimum is the number of ships *currently*
// placed in this row.  The maximum is the number of ships
// that could be placed if all unsolved squares were filled
// with ships.
func (board *Board) MinMaxShipsInRow(coord Coord) (int, int) {
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
func (board *Board) MinMaxShipsInColumn(coord Coord) (int, int) {
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


func (board *Board) CalcPossibleSquaresFor(coord Coord) *vector.IntVector {
    var possibilities vector.IntVector

    var requireWater = false
    var requireShip = false

    desired := board.rowClues[coord.Row()]
    minships, maxships := board.MinMaxShipsInRow(coord)
    switch {
    case minships > desired:
        // TODO: Should this be a panic?
        //fmt.Printf("Unsolvable:  too many ships in row at coord %v", coord)
        return &possibilities
    case minships == desired:
        requireWater = true
    case maxships == desired:
        requireShip = true
    case maxships < desired:
        // TODO: Should this be a panic?
        //fmt.Printf("Unsolvable:  too few ships in row at coord %v", coord)
        return &possibilities
    }

    desired = board.columnClues[coord.Column()]
    minships, maxships = board.MinMaxShipsInColumn(coord)
    switch {
    case minships > desired:
        // TODO: Should this be a panic?
        //fmt.Printf("Unsolvable:  too many ships in column at coord %v", coord)
        return &possibilities
    case minships == desired:
        requireWater = true
    case maxships == desired:
        requireShip = true
    case maxships < desired:
        // TODO: Should this be a panic?
        //fmt.Printf("Unsolvable:  too few ships in column at coord %v", coord)
        return &possibilities
    }

    for _, square := range SQUARES {
        switch {
        case requireWater && !square.IsWater():
        case requireShip && !square.IsShip():
        case !board.SquareAt(coord.Above()).CanAppearAbove(square):
        case !board.SquareAt(coord.Below()).CanAppearBelow(square):
        case !board.SquareAt(coord.Right()).CanAppearRightOf(square):
        case !board.SquareAt(coord.Left()).CanAppearLeftOf(square):
        case !board.SquareAt(coord.Above().Left()).CanAppearDiagonallyAdjacentTo(square):
        case !board.SquareAt(coord.Above().Right()).CanAppearDiagonallyAdjacentTo(square):
        case !board.SquareAt(coord.Below().Left()).CanAppearDiagonallyAdjacentTo(square):
        case !board.SquareAt(coord.Below().Right()).CanAppearDiagonallyAdjacentTo(square):
        default:
            possibilities.Push(int(square))
        }
    }
    return &possibilities
}


// Finds the next coord on the board to solve, choosing one
// of the coordinates with the fewest possibilities
func (board *Board) NextCoordToSolve() (Coord, bool) {
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
    if !board.IsValid() {
        return false
    }
    coord, foundOne := board.NextCoordToSolve()
    if !foundOne {
        return true // solved!
    }
    possibleSquares := board.CalcPossibleSquaresFor(coord)

    for _, possibleSquare := range *possibleSquares {
        square := Square(possibleSquare)
        board.SetSquareAt(coord, square)
        //fmt.Printf("Trying %v at %v\n%v\n", square, coord, board)
        if board.Solve() {
            return true
        }
        //fmt.Printf("Undoing %v at %v.\n", square, coord)
        board.SetSquareAt(coord, UNSOLVED)
    }
    //fmt.Printf("Backtracking from coord %v:\n", coord)
    return false
}
