// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package battleship


// Returns an array of various samples.
func SampleClues() []*Clues {
    return _CLUES
}

var _CLUES = []*Clues{
    &Clues{
        title:       "rectangular by Mick",
        rowClues:    []int{2, 2},
        columnClues: []int{1, 0, 2, 0, 1},
        shipLengths: []int{2, 1, 1},
        initialSquares: map[Coord]Square{
            NewCoord(1, 0): WATER,
        },
        // Solution is:
        //   ~~^~o
        //   o~v~~
    },

    &Clues{
        title:       "6x6 example from conceptis",
        rowClues:    []int{4, 0, 2, 1, 2, 1},
        columnClues: []int{1, 0, 4, 0, 3, 2},
        shipLengths: []int{3, 2, 2, 1, 1, 1},
        initialSquares: map[Coord]Square{
            NewCoord(2, 2): WATER,
        },
    },

    &Clues{
        title:       "8x8 easy from conceptis",
        rowClues:    []int{2, 4, 2, 3, 2, 1, 4, 2},
        columnClues: []int{5, 0, 5, 1, 2, 1, 2, 4},
        shipLengths: []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1},
        initialSquares: map[Coord]Square{
            NewCoord(0, 2): WATER,
            NewCoord(4, 6): TOP,
        },
    },

    &Clues{
        title:       "8x8 difficult fron conceptis",
        rowClues:    []int{1, 3, 1, 4, 2, 4, 0, 5},
        columnClues: []int{4, 3, 2, 1, 4, 0, 5, 1},
        shipLengths: []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1},
        initialSquares: map[Coord]Square{
            NewCoord(7, 2): SINGLE,
        },
    },

    &Clues{
        title:       "10x10 easy from conceptis",
        rowClues:    []int{3, 2, 3, 3, 1, 1, 2, 1, 3, 1},
        columnClues: []int{4, 0, 3, 1, 2, 2, 1, 2, 1, 4},
        shipLengths: []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1},
        initialSquares: map[Coord]Square{
            NewCoord(0, 2): WATER,
            NewCoord(3, 7): MIDDLE,
            NewCoord(5, 4): TOP,
        },
    },

    &Clues{
        title:       "GAMES World of Puzzles, Nov 2009, 1-Seaman",
        rowClues:    []int{4, 2, 1, 1, 1, 1, 1, 3, 5, 1},
        columnClues: []int{0, 7, 0, 5, 0, 3, 0, 2, 0, 3},
        shipLengths: []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1},
        initialSquares: map[Coord]Square{
            NewCoord(4, 9): TOP,
            NewCoord(6, 1): WATER,
        },
    },

    &Clues{
        title:       "GAMES World of Puzzles, Nov 2009, 2-Petty Officer",
        rowClues:    []int{0, 2, 4, 3, 2, 2, 3, 1, 2, 1},
        columnClues: []int{3, 0, 0, 4, 0, 3, 2, 4, 0, 4},
        shipLengths: []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1},
        initialSquares: map[Coord]Square{
            NewCoord(4, 3): MIDDLE,
            NewCoord(4, 5): MIDDLE,
            NewCoord(7, 0): WATER,
        },
    },

    &Clues{
        title:       "GAMES World of Puzzles, Nov 2009, 3-Ensign",
        rowClues:    []int{2, 1, 1, 1, 0, 1, 0, 6, 2, 6},
        columnClues: []int{0, 5, 0, 5, 0, 2, 2, 2, 2, 2},
        shipLengths: []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1},
        initialSquares: map[Coord]Square{
            NewCoord(0, 1): TOP,
            NewCoord(2, 3): WATER,
            NewCoord(5, 5): SINGLE,
            NewCoord(7, 5): WATER,
        },
    },

    &Clues{
        title:       "GAMES World of Puzzles, Nov 2009, 4-Captain",
        rowClues:    []int{0, 2, 2, 5, 1, 1, 6, 1, 2, 0},
        columnClues: []int{5, 0, 3, 1, 1, 3, 1, 2, 1, 3},
        shipLengths: []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1},
        initialSquares: map[Coord]Square{
            NewCoord(2, 0): BOTTOM,
            NewCoord(2, 9): MIDDLE,
            NewCoord(8, 4): SINGLE,
        },
    },

    &Clues{
        title:       "GAMES World of Puzzles, Nov 2009, 5-Commodore",
        rowClues:    []int{4, 2, 1, 0, 2, 2, 2, 4, 0, 3},
        columnClues: []int{1, 0, 5, 0, 4, 2, 3, 2, 0, 3},
        shipLengths: []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1},
        initialSquares: map[Coord]Square{
            NewCoord(0, 5): LEFT,
            NewCoord(0, 9): SINGLE,
            NewCoord(5, 9): BOTTOM,
        },
    },

    &Clues{
        title:       "GAMES World of Puzzles, Nov 2009, 6-Admiral",
        rowClues:    []int{2, 1, 4, 3, 3, 2, 3, 1, 1, 0},
        columnClues: []int{4, 0, 1, 1, 2, 4, 0, 3, 0, 5},
        shipLengths: []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1},
        initialSquares: map[Coord]Square{
            NewCoord(3, 0): BOTTOM,
            NewCoord(2, 3): MIDDLE,
            NewCoord(5, 9): WATER,
        },
    },
}
