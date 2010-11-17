// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package battleship

import (
    "testing"
)

func doSolve(t *testing.T, clues *Clues) {
    board := NewBoard(clues)
    t.Logf("Solving: %v\n%v\n", clues.Title(), board)
    if !board.Solve() {
        t.Error("Could not find solution")
    } else {
        t.Logf("Found solution:\n%v\n", board)
    }
}

func TestBoard_Solve_SampleClues_All(t *testing.T) {
    for _, clues := range SampleClues() {
        doSolve(t, clues)
    }
}

func TestBoard_Solve_SampleClues_0(t *testing.T) {
    doSolve(t, SampleClues()[0])
}

func TestBoard_Solve_SampleClues_1(t *testing.T) {
    doSolve(t, SampleClues()[1])
}

func TestBoard_Solve_SampleClues_2(t *testing.T) {
    doSolve(t, SampleClues()[2])
}

func BenchmarkBoard_Solve_SampleClues_0(b *testing.B) {
    clues := SampleClues()[0]
    for i := 0; i < b.N; i++ {
        NewBoard(clues).Solve()
    }
}

func BenchmarkBoard_Solve_SampleClues_1(b *testing.B) {
    clues := SampleClues()[1]
    for i := 0; i < b.N; i++ {
        NewBoard(clues).Solve()
    }
}

func BenchmarkBoard_Solve_SampleClues_2(b *testing.B) {
    clues := SampleClues()[2]
    for i := 0; i < b.N; i++ {
        NewBoard(clues).Solve()
    }
}

func BenchmarkBoard_Solve_SampleClues_All(b *testing.B) {
    for i := 0; i < b.N; i++ {
        for _, clues := range SampleClues() {
            NewBoard(clues).Solve()
        }
    }
}
