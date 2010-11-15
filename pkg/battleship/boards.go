// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package battleship


// Returns a very easy 6x6 board
func Board0() *Board {
    return NewBoard(&Clues{
        rowClues:    []int{4, 0, 2, 1, 2, 1},
        columnClues: []int{1, 0, 4, 0, 3, 2},
        shipLengths: []int{3, 2, 2, 1, 1, 1},
        initialSquares: map[Coord]Square{
            NewCoord(2, 2): WATER,
        }})
}

// Returns a very easy 8x8 board
func Board1() *Board {
    return NewBoard(&Clues{
        rowClues:    []int{2, 4, 2, 3, 2, 1, 4, 2},
        columnClues: []int{5, 0, 5, 1, 2, 1, 2, 4},
        shipLengths: []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1},
        initialSquares: map[Coord]Square{
            NewCoord(0, 2): WATER,
            NewCoord(4, 6): TOP,
        }})
}

// Returns an easy 10x10 board
func Board2() *Board {
    return NewBoard(&Clues{
        rowClues:    []int{3, 2, 3, 3, 1, 1, 2, 1, 3, 1},
        columnClues: []int{4, 0, 3, 1, 2, 2, 1, 2, 1, 4},
        shipLengths: []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1},
        initialSquares: map[Coord]Square{
            NewCoord(0, 2): WATER,
            NewCoord(3, 7): MIDDLE,
            NewCoord(5, 4): TOP,
        }})
}
