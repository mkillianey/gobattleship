// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package main

import (
    "fmt"
    "battleship"
)


func main() {
    //board := battleship.Board0()  // works!
    //board := battleship.Board1()  // works, but slow
    //board := battleship.Board2()  // requires ship knowledge
    board := battleship.Board3()
    fmt.Printf("Attempting board:\n%v\n", board)
    if board.Solve() {
        fmt.Printf("Solved:\n%v\n", board)
    } else {
        fmt.Printf("Could not solve!")
    }
}
