package battleship

import (
    "fmt"
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
        return true // solved!
    }
    possibleSquares := board.CalcPossibleSquaresFor(coord)

    if possibleSquares.Len() == 0 {
        return false
    }
    for _, possibleSquare := range *possibleSquares {
        square := Square(possibleSquare)
        board.SetSquareAt(coord, square)
        if board.IsValid() && board.Solve() {
            return true
        }
        board.SetSquareAt(coord, UNSOLVED)
    }
    return false
}
