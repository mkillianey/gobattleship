package battleship

import (
    testing "testing"
)

func assertCanAppearAbove(t *testing.T, s1 Square, s2 Square) {
    if (!s1.CanAppearAbove(s2)) {
        t.Errorf("%v.CanAppearAbove(%v) should be true", s1, s2);
    }
}

func assertCannotAppearAbove(t *testing.T, s1 Square, s2 Square) {
    if (s1.CanAppearAbove(s2)) {
        t.Errorf("%v.CanAppearAbove(%v) should be true", s1, s2);
    }
}

func TestSquareCanAppearAbove(t *testing.T) {
    assertCanAppearAbove(t, WATER, LEFT)
    assertCannotAppearAbove(t, WATER, BOTTOM)
}

func TestSolveBoard0(t *testing.T) {
    if !board0().Solve() {
        t.Error("Could not find solution")
    }
}

func TestSolveBoard1(t *testing.T) {
    if !board1().Solve() {
        t.Error("Could not find solution")
    }
}

func TestSolveBoard2(t *testing.T) {
    if !board2().Solve() {
        t.Error("Could not find solution")
    }
}

