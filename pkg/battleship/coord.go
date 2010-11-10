// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package battleship

import (
    "fmt"
)

type Coord struct {
    row, column int
}

func (coord *Coord) String() string {
    return fmt.Sprintf("(%v,%v)", coord.row, coord.column)
}

func (coord *Coord) Equals(other *Coord) bool {
    return coord.row == other.row && other.column == other.column
}

// Returns the coord adjacent immediately above this coord
func (coord *Coord) Above() *Coord {
    return &Coord{row: coord.row - 1, column: coord.column}
}

// Returns the coord adjacent immediately below this coord
func (coord *Coord) Below() *Coord {
    return &Coord{row: coord.row + 1, column: coord.column}
}

// Returns the coord immediately to the left of this coord
func (coord *Coord) Left() *Coord {
    return &Coord{row: coord.row, column: coord.column - 1}
}

// Returns the coord immediately to the right of this coord
func (coord *Coord) Right() *Coord {
    return &Coord{row: coord.row, column: coord.column + 1}
}

// Returns a new coord in this coord's row, but with the given column
func (coord *Coord) WithColumn(column int) *Coord {
    return &Coord{row: coord.row, column: column}
}

// Returns a new coord in this coord's column, but with the given row
func (coord *Coord) WithRow(row int) *Coord {
    return &Coord{row: row, column: coord.column}
}
