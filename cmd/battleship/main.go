// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package main

import (
    fmt        "fmt"
    battleship "battleship"
)


func main() {
    //board := battleship.Board0()  // works!
    //board := battleship.Board1()  // works, but slow
    board := battleship.Board2() // doesn't get ships right

    fmt.Printf("Solving board:\n%v\n", board)
    if board.Solve() {
        fmt.Println("Solved!")
    } else {
        fmt.Println("Could not solve!")
    }
}
