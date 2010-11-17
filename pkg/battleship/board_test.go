// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package battleship

import (
    "testing"
)

func doSolve(t *testing.T, board *Board) {
    if !board.Solve() {
        t.Error("Could not find solution")
    } else {
        t.Logf("Found solution:\n%v\n", board)
    }
}

func TestBoard_SolveBoard0(t *testing.T) {
    doSolve(t, Board0())
}

func TestBoard_SolveBoard1(t *testing.T) {
    doSolve(t, Board1())
}

func TestBoard_SolveBoard2(t *testing.T) {
    doSolve(t, Board2())
}

func TestBoard_SolveBoard3(t *testing.T) {
    doSolve(t, Board3())
}

func BenchmarkBoard_SolveBoard0(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Board0().Solve()
    }
}

func BenchmarkBoard_SolveBoard1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Board1().Solve()
    }
}

func BenchmarkBoard_SolveBoard2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Board2().Solve()
    }
}

func BenchmarkBoard_SolveBoard3(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Board3().Solve()
    }
}
