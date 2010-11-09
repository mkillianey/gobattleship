package battleship

import (
    "container/vector"
)

func Board0() *Board{
    return NewBoard(
        []int{4, 0, 2, 1, 2, 1},
        []int{1, 0, 4, 0, 3, 2},
        []int{3, 2, 2, 1, 1, 1},
        []struct {
                row, column int
                square Square
        }{
           {row: 2, column: 2, square: WATER},
        })
}

func Board1() *Board{
    return NewBoard(
        []int{2, 4, 2, 3, 2, 1, 4, 2},
        []int{5, 0, 5, 1, 2, 1, 2, 4},
        []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1},
        []struct {
                row, column int
                square Square
        }{
           {row: 0, column: 2, square: WATER},
           {row: 4, column: 6, square: TOP},
        })
}

func Board2() *Board{
    return NewBoard(
        []int{3, 2, 3, 3, 1, 1, 2, 1, 3, 1},
        []int{4, 0, 3, 1, 2, 2, 1, 2, 1, 4},
        []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1},
        []struct {
                row, column int
                square Square
        }{
           {row: 0, column: 2, square: WATER},
           {row: 3, column: 7, square: MIDDLE},
           {row: 5, column: 4, square: TOP},
        })
}

var BOARDS vector.Vector

func init() {
    BOARDS.Push(Board0())
    BOARDS.Push(Board1())
    BOARDS.Push(Board2())
}
