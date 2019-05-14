// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package main

import (
    "fmt"
    "github.com/mkillianey/gobattleship/battleship"
)


func main() {
    solver := battleship.NewSolver()

    for _, clues := range battleship.SampleClues() {
        fmt.Println("Attempting to solve:")
        fmt.Println(clues)

        if solution, ok := solver.SolveClues(clues); ok {
            fmt.Printf("Solved:\n%v\n", solution)
        } else {
            fmt.Printf("Could not solve '%v'!", clues.Title())
        }
    }
}
