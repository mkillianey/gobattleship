// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package battleship

import (
    "fmt"
)

type Coord int

func NewCoord(row int, column int) Coord {
    return Coord((row << 16) | (column & 0xffff))
}

// Returns a new coord in this coord's row, but with the given column
func (coord Coord) WithColumn(column int) Coord {
    return Coord((int(coord) &^ 0xffff) | (column & 0xffff))
}

// Returns a new coord in this coord's column, but with the given row
func (coord Coord) WithRow(row int) Coord {
    return Coord((row << 16) | (int(coord) & 0xffff))
}

// Returns the row of this coord
func (coord Coord) Row() int {
    return int(int16(coord >> 16))
}

// Returns the column of this coord
func (coord Coord) Column() int {
    return int(int16(coord))
}

func (coord Coord) String() string {
    return fmt.Sprintf("(%v,%v)", coord.Row(), coord.Column())
}

// Returns the coord adjacent immediately above this coord
func (coord Coord) Above() Coord {
    //return coord.WithRow(coord.Row() - 1)
    return Coord(int(coord) - 0x10000)
}

// Returns the coord adjacent immediately below this coord
func (coord Coord) Below() Coord {
    //return coord.WithRow(coord.Row() + 1)
    return Coord(int(coord) + 0x10000)
}

// Returns the coord immediately to the left of this coord
func (coord Coord) Left() Coord {
    //return coord.WithColumn(coord.Column() - 1)
    i := int(coord)
    if i & 0xffff == 0 {
        i += 0x10000
    }
    return Coord(i - 1)
}

// Returns the coord immediately to the right of this coord
func (coord Coord) Right() Coord {
    // return coord.WithColumn(coord.Column() + 1)
    i := int(coord)
    if i & 0xffff == 0xffff {
        i -= 0x10000
    }
    return Coord(i + 1)
}
