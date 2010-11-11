// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package battleship


// Returns a very easy 6x6 board
func Board0() *Board {
    return NewBoard(
        []int{4, 0, 2, 1, 2, 1},
        []int{1, 0, 4, 0, 3, 2},
        []int{3, 2, 2, 1, 1, 1},
        []struct {
            row, column int
            square      Square
        }{
            {row: 2, column: 2, square: WATER},
        })
}

// Returns a very easy 8x8 board
func Board1() *Board {
    return NewBoard(
        []int{2, 4, 2, 3, 2, 1, 4, 2},
        []int{5, 0, 5, 1, 2, 1, 2, 4},
        []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1},
        []struct {
            row, column int
            square      Square
        }{
            {row: 0, column: 2, square: WATER},
            {row: 4, column: 6, square: TOP},
        })
}

// Returns an easy 10x10 board
func Board2() *Board {
    return NewBoard(
        []int{3, 2, 3, 3, 1, 1, 2, 1, 3, 1},
        []int{4, 0, 3, 1, 2, 2, 1, 2, 1, 4},
        []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1},
        []struct {
            row, column int
            square      Square
        }{
            {row: 0, column: 2, square: WATER},
            {row: 3, column: 7, square: MIDDLE},
            {row: 5, column: 4, square: TOP},
        })
}
