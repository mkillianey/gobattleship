// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package battleship

import (
    "testing"
)

func assertCanAppearAbove(t *testing.T, s1 Square, s2 Square) {
    if !s1.CanAppearAbove(s2) {
        t.Errorf("%v.CanAppearAbove(%v) should be true", s1, s2)
    }
}

func assertCannotAppearAbove(t *testing.T, s1 Square, s2 Square) {
    if s1.CanAppearAbove(s2) {
        t.Errorf("%v.CanAppearAbove(%v) should be true", s1, s2)
    }
}

func TestSquareCanAppearAbove(t *testing.T) {
    assertCanAppearAbove(t, WATER, LEFT)
    assertCannotAppearAbove(t, WATER, BOTTOM)
}

func doSolve(t *testing.T, board *Board) {
    if !board.Solve() {
        t.Error("Could not find solution")
    } else {
        t.Logf("Found solution:\n%v\n", board)
    }
}

func TestSolveBoard0(t *testing.T) {
    doSolve(t, Board0())
}

func TestSolveBoard1(t *testing.T) {
    doSolve(t, Board1())
}

func TestSolveBoard2(t *testing.T) {
    doSolve(t, Board2())
}

func BenchmarkSolveBoard0(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Board0().Solve()
    }
}

func BenchmarkSolveBoard1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Board1().Solve()
    }
}

