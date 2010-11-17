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

    for _, clues := range battleship.SampleClues() {
        board := battleship.NewBoard(clues)
        fmt.Printf("Attempting to solve: %v\n%v\n", clues.Title(), board)
        if board.Solve() {
            fmt.Printf("Solved:\n%v\n", board)
        } else {
            fmt.Printf("Could not solve '%v'!", clues.Title())
        }
    }
}
