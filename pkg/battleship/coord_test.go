// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package battleship

import (
    "testing"
)

func assertEquals(t *testing.T, c1 *Coord, c2 *Coord) {
    if !c1.Equals(c2) {
        t.Errorf("%v.Equals(%v) should be true", c1, c2)
    }
}

var R1_C1 = &Coord{row: 1, column: 1}
var R1_C2 = &Coord{row: 1, column: 2}
var R1_C3 = &Coord{row: 1, column: 3}

var R2_C1 = &Coord{row: 2, column: 1}
var R2_C2 = &Coord{row: 2, column: 2}
var R2_C3 = &Coord{row: 2, column: 3}

var R3_C1 = &Coord{row: 3, column: 1}
var R3_C2 = &Coord{row: 3, column: 2}
var R3_C3 = &Coord{row: 3, column: 3}


func TestCoord_AboveAndBelow(t *testing.T) {
    assertEquals(t, R2_C2.Above(), R1_C2)
    assertEquals(t, R2_C2.Below(), R3_C2)
}

func TestCoord_LeftAndRight(t *testing.T) {
    assertEquals(t, R2_C2.Left(), R2_C1)
    assertEquals(t, R2_C2.Right(), R2_C3)
}

func TestCoord_Chaining(t *testing.T) {
    assertEquals(t, R2_C2.Right().Below(), R3_C3)
    assertEquals(t, R2_C2.Below().Right(), R3_C3)
}

func TestCoord_WithRow(t *testing.T) {
    assertEquals(t, R1_C1.WithRow(3), R3_C1)
    assertEquals(t, R1_C1.WithColumn(3), R1_C3)
}
