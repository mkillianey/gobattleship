// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package battleship

import (
    "testing"
)

func checkSquareCanAppearAbove(t *testing.T, expected bool, s1 Square, s2 Square) {
    if actual := s1.CanAppearAbove(s2); actual != expected {
        t.Errorf("%v.CanAppearAbove(%v) should be %v, was %v", s1, s2, expected, actual)
    }
}

func checkSquareCanAppearLeftOf(t *testing.T, expected bool, s1 Square, s2 Square) {
    if actual := s1.CanAppearLeftOf(s2); actual != expected {
        t.Errorf("%v.CanAppearLeftOf(%v) should be %v, was %v", s1, s2, expected, actual)
    }
}


func TestSquare_String(t *testing.T) {
    if s := WATER.String(); s != "WATER" {
        t.Errorf("Wrong string for WATER: %v", s)
    }
    if s := LEFT.String(); s != "LEFT" {
        t.Errorf("Wrong string for LEFT: %v", s)
    }
}

func TestSquare_CanAppearAbove(t *testing.T) {
    checkSquareCanAppearAbove(t, true, UNSOLVED, UNSOLVED)
    checkSquareCanAppearAbove(t, true, UNSOLVED, WATER)
    checkSquareCanAppearAbove(t, true, UNSOLVED, TOP)
    checkSquareCanAppearAbove(t, true, UNSOLVED, BOTTOM)
    checkSquareCanAppearAbove(t, true, UNSOLVED, LEFT)
    checkSquareCanAppearAbove(t, true, UNSOLVED, RIGHT)
    checkSquareCanAppearAbove(t, true, UNSOLVED, SINGLE)
    checkSquareCanAppearAbove(t, true, UNSOLVED, MIDDLE)
    checkSquareCanAppearAbove(t, true, UNSOLVED, OUT_OF_BOUNDS)

    checkSquareCanAppearAbove(t, true, WATER, UNSOLVED)
    checkSquareCanAppearAbove(t, true, WATER, WATER)
    checkSquareCanAppearAbove(t, true, WATER, TOP)
    checkSquareCanAppearAbove(t, false, WATER, BOTTOM)
    checkSquareCanAppearAbove(t, true, WATER, LEFT)
    checkSquareCanAppearAbove(t, true, WATER, RIGHT)
    checkSquareCanAppearAbove(t, true, WATER, SINGLE)
    checkSquareCanAppearAbove(t, true, WATER, MIDDLE)
    checkSquareCanAppearAbove(t, true, WATER, OUT_OF_BOUNDS)

    checkSquareCanAppearAbove(t, true, TOP, UNSOLVED)
    checkSquareCanAppearAbove(t, false, TOP, WATER)
    checkSquareCanAppearAbove(t, false, TOP, TOP)
    checkSquareCanAppearAbove(t, true, TOP, BOTTOM)
    checkSquareCanAppearAbove(t, false, TOP, LEFT)
    checkSquareCanAppearAbove(t, false, TOP, RIGHT)
    checkSquareCanAppearAbove(t, false, TOP, SINGLE)
    checkSquareCanAppearAbove(t, true, TOP, MIDDLE)
    checkSquareCanAppearAbove(t, false, TOP, OUT_OF_BOUNDS)

    checkSquareCanAppearAbove(t, true, BOTTOM, UNSOLVED)
    checkSquareCanAppearAbove(t, true, BOTTOM, WATER)
    checkSquareCanAppearAbove(t, false, BOTTOM, TOP)
    checkSquareCanAppearAbove(t, false, BOTTOM, BOTTOM)
    checkSquareCanAppearAbove(t, false, BOTTOM, LEFT)
    checkSquareCanAppearAbove(t, false, BOTTOM, RIGHT)
    checkSquareCanAppearAbove(t, false, BOTTOM, SINGLE)
    checkSquareCanAppearAbove(t, false, BOTTOM, MIDDLE)
    checkSquareCanAppearAbove(t, true, BOTTOM, OUT_OF_BOUNDS)

    checkSquareCanAppearAbove(t, true, LEFT, UNSOLVED)
    checkSquareCanAppearAbove(t, true, LEFT, WATER)
    checkSquareCanAppearAbove(t, false, LEFT, TOP)
    checkSquareCanAppearAbove(t, false, LEFT, BOTTOM)
    checkSquareCanAppearAbove(t, false, LEFT, LEFT)
    checkSquareCanAppearAbove(t, false, LEFT, RIGHT)
    checkSquareCanAppearAbove(t, false, LEFT, SINGLE)
    checkSquareCanAppearAbove(t, false, LEFT, MIDDLE)
    checkSquareCanAppearAbove(t, true, LEFT, OUT_OF_BOUNDS)

    checkSquareCanAppearAbove(t, true, RIGHT, UNSOLVED)
    checkSquareCanAppearAbove(t, true, RIGHT, WATER)
    checkSquareCanAppearAbove(t, false, RIGHT, TOP)
    checkSquareCanAppearAbove(t, false, RIGHT, BOTTOM)
    checkSquareCanAppearAbove(t, false, RIGHT, LEFT)
    checkSquareCanAppearAbove(t, false, RIGHT, RIGHT)
    checkSquareCanAppearAbove(t, false, RIGHT, SINGLE)
    checkSquareCanAppearAbove(t, false, RIGHT, MIDDLE)
    checkSquareCanAppearAbove(t, true, RIGHT, OUT_OF_BOUNDS)

    checkSquareCanAppearAbove(t, true, SINGLE, UNSOLVED)
    checkSquareCanAppearAbove(t, true, SINGLE, WATER)
    checkSquareCanAppearAbove(t, false, SINGLE, TOP)
    checkSquareCanAppearAbove(t, false, SINGLE, BOTTOM)
    checkSquareCanAppearAbove(t, false, SINGLE, LEFT)
    checkSquareCanAppearAbove(t, false, SINGLE, RIGHT)
    checkSquareCanAppearAbove(t, false, SINGLE, SINGLE)
    checkSquareCanAppearAbove(t, false, SINGLE, MIDDLE)
    checkSquareCanAppearAbove(t, true, SINGLE, OUT_OF_BOUNDS)

    checkSquareCanAppearAbove(t, true, MIDDLE, UNSOLVED)
    checkSquareCanAppearAbove(t, true, MIDDLE, WATER)
    checkSquareCanAppearAbove(t, false, MIDDLE, TOP)
    checkSquareCanAppearAbove(t, true, MIDDLE, BOTTOM)
    checkSquareCanAppearAbove(t, false, MIDDLE, LEFT)
    checkSquareCanAppearAbove(t, false, MIDDLE, RIGHT)
    checkSquareCanAppearAbove(t, false, MIDDLE, SINGLE)
    checkSquareCanAppearAbove(t, true, MIDDLE, MIDDLE)
    checkSquareCanAppearAbove(t, true, MIDDLE, OUT_OF_BOUNDS)

    checkSquareCanAppearAbove(t, true, OUT_OF_BOUNDS, UNSOLVED)
    checkSquareCanAppearAbove(t, true, OUT_OF_BOUNDS, WATER)
    checkSquareCanAppearAbove(t, true, OUT_OF_BOUNDS, TOP)
    checkSquareCanAppearAbove(t, false, OUT_OF_BOUNDS, BOTTOM)
    checkSquareCanAppearAbove(t, true, OUT_OF_BOUNDS, LEFT)
    checkSquareCanAppearAbove(t, true, OUT_OF_BOUNDS, RIGHT)
    checkSquareCanAppearAbove(t, true, OUT_OF_BOUNDS, SINGLE)
    checkSquareCanAppearAbove(t, true, OUT_OF_BOUNDS, MIDDLE)
    checkSquareCanAppearAbove(t, true, OUT_OF_BOUNDS, OUT_OF_BOUNDS)
}

func TestSquare_CanAppearLeftOf(t *testing.T) {
    checkSquareCanAppearLeftOf(t, true, UNSOLVED, UNSOLVED)
    checkSquareCanAppearLeftOf(t, true, UNSOLVED, WATER)
    checkSquareCanAppearLeftOf(t, true, UNSOLVED, TOP)
    checkSquareCanAppearLeftOf(t, true, UNSOLVED, BOTTOM)
    checkSquareCanAppearLeftOf(t, true, UNSOLVED, LEFT)
    checkSquareCanAppearLeftOf(t, true, UNSOLVED, RIGHT)
    checkSquareCanAppearLeftOf(t, true, UNSOLVED, SINGLE)
    checkSquareCanAppearLeftOf(t, true, UNSOLVED, MIDDLE)
    checkSquareCanAppearLeftOf(t, true, UNSOLVED, OUT_OF_BOUNDS)

    checkSquareCanAppearLeftOf(t, true, WATER, UNSOLVED)
    checkSquareCanAppearLeftOf(t, true, WATER, WATER)
    checkSquareCanAppearLeftOf(t, true, WATER, TOP)
    checkSquareCanAppearLeftOf(t, true, WATER, BOTTOM)
    checkSquareCanAppearLeftOf(t, true, WATER, LEFT)
    checkSquareCanAppearLeftOf(t, false, WATER, RIGHT)
    checkSquareCanAppearLeftOf(t, true, WATER, SINGLE)
    checkSquareCanAppearLeftOf(t, true, WATER, MIDDLE)
    checkSquareCanAppearLeftOf(t, true, WATER, OUT_OF_BOUNDS)

    checkSquareCanAppearLeftOf(t, true, TOP, UNSOLVED)
    checkSquareCanAppearLeftOf(t, true, TOP, WATER)
    checkSquareCanAppearLeftOf(t, false, TOP, TOP)
    checkSquareCanAppearLeftOf(t, false, TOP, BOTTOM)
    checkSquareCanAppearLeftOf(t, false, TOP, LEFT)
    checkSquareCanAppearLeftOf(t, false, TOP, RIGHT)
    checkSquareCanAppearLeftOf(t, false, TOP, SINGLE)
    checkSquareCanAppearLeftOf(t, false, TOP, MIDDLE)
    checkSquareCanAppearLeftOf(t, true, TOP, OUT_OF_BOUNDS)

    checkSquareCanAppearLeftOf(t, true, BOTTOM, UNSOLVED)
    checkSquareCanAppearLeftOf(t, true, BOTTOM, WATER)
    checkSquareCanAppearLeftOf(t, false, BOTTOM, TOP)
    checkSquareCanAppearLeftOf(t, false, BOTTOM, BOTTOM)
    checkSquareCanAppearLeftOf(t, false, BOTTOM, LEFT)
    checkSquareCanAppearLeftOf(t, false, BOTTOM, RIGHT)
    checkSquareCanAppearLeftOf(t, false, BOTTOM, SINGLE)
    checkSquareCanAppearLeftOf(t, false, BOTTOM, MIDDLE)
    checkSquareCanAppearLeftOf(t, true, BOTTOM, OUT_OF_BOUNDS)

    checkSquareCanAppearLeftOf(t, true, LEFT, UNSOLVED)
    checkSquareCanAppearLeftOf(t, false, LEFT, WATER)
    checkSquareCanAppearLeftOf(t, false, LEFT, TOP)
    checkSquareCanAppearLeftOf(t, false, LEFT, BOTTOM)
    checkSquareCanAppearLeftOf(t, false, LEFT, LEFT)
    checkSquareCanAppearLeftOf(t, true, LEFT, RIGHT)
    checkSquareCanAppearLeftOf(t, false, LEFT, SINGLE)
    checkSquareCanAppearLeftOf(t, true, LEFT, MIDDLE)
    checkSquareCanAppearLeftOf(t, false, LEFT, OUT_OF_BOUNDS)

    checkSquareCanAppearLeftOf(t, true, RIGHT, UNSOLVED)
    checkSquareCanAppearLeftOf(t, true, RIGHT, WATER)
    checkSquareCanAppearLeftOf(t, false, RIGHT, TOP)
    checkSquareCanAppearLeftOf(t, false, RIGHT, BOTTOM)
    checkSquareCanAppearLeftOf(t, false, RIGHT, LEFT)
    checkSquareCanAppearLeftOf(t, false, RIGHT, RIGHT)
    checkSquareCanAppearLeftOf(t, false, RIGHT, SINGLE)
    checkSquareCanAppearLeftOf(t, false, RIGHT, MIDDLE)
    checkSquareCanAppearLeftOf(t, true, RIGHT, OUT_OF_BOUNDS)

    checkSquareCanAppearLeftOf(t, true, SINGLE, UNSOLVED)
    checkSquareCanAppearLeftOf(t, true, SINGLE, WATER)
    checkSquareCanAppearLeftOf(t, false, SINGLE, TOP)
    checkSquareCanAppearLeftOf(t, false, SINGLE, BOTTOM)
    checkSquareCanAppearLeftOf(t, false, SINGLE, LEFT)
    checkSquareCanAppearLeftOf(t, false, SINGLE, RIGHT)
    checkSquareCanAppearLeftOf(t, false, SINGLE, SINGLE)
    checkSquareCanAppearLeftOf(t, false, SINGLE, MIDDLE)
    checkSquareCanAppearLeftOf(t, true, SINGLE, OUT_OF_BOUNDS)

    checkSquareCanAppearLeftOf(t, true, MIDDLE, UNSOLVED)
    checkSquareCanAppearLeftOf(t, true, MIDDLE, WATER)
    checkSquareCanAppearLeftOf(t, false, MIDDLE, TOP)
    checkSquareCanAppearLeftOf(t, false, MIDDLE, BOTTOM)
    checkSquareCanAppearLeftOf(t, false, MIDDLE, LEFT)
    checkSquareCanAppearLeftOf(t, true, MIDDLE, RIGHT)
    checkSquareCanAppearLeftOf(t, false, MIDDLE, SINGLE)
    checkSquareCanAppearLeftOf(t, true, MIDDLE, MIDDLE)
    checkSquareCanAppearLeftOf(t, true, MIDDLE, OUT_OF_BOUNDS)

    checkSquareCanAppearLeftOf(t, true, OUT_OF_BOUNDS, UNSOLVED)
    checkSquareCanAppearLeftOf(t, true, OUT_OF_BOUNDS, WATER)
    checkSquareCanAppearLeftOf(t, true, OUT_OF_BOUNDS, TOP)
    checkSquareCanAppearLeftOf(t, true, OUT_OF_BOUNDS, BOTTOM)
    checkSquareCanAppearLeftOf(t, true, OUT_OF_BOUNDS, LEFT)
    checkSquareCanAppearLeftOf(t, false, OUT_OF_BOUNDS, RIGHT)
    checkSquareCanAppearLeftOf(t, true, OUT_OF_BOUNDS, SINGLE)
    checkSquareCanAppearLeftOf(t, true, OUT_OF_BOUNDS, MIDDLE)
    checkSquareCanAppearLeftOf(t, true, OUT_OF_BOUNDS, OUT_OF_BOUNDS)
}
