// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package battleship

import (
    "testing"
)

func assertCoordEquals(t *testing.T, c1 Coord, c2 Coord) {
    if c1 != c2 {
        t.Errorf("%v == %v should be true", c1, c2)
    }
}

func assertIntEquals(t *testing.T, i1 int, i2 int) {
    if i1 != i2 {
        t.Errorf("%v == %v should be true", i1, i2)
    }
}

var R1_C1 = NewCoord(1, 1)
var R1_C2 = NewCoord(1, 2)
var R1_C3 = NewCoord(1, 3)

var R2_C1 = NewCoord(2, 1)
var R2_C2 = NewCoord(2, 2)
var R2_C3 = NewCoord(2, 3)

var R3_C1 = NewCoord(3, 1)
var R3_C2 = NewCoord(3, 2)
var R3_C3 = NewCoord(3, 3)


func TestCoord_AboveAndBelow(t *testing.T) {
    assertCoordEquals(t, R2_C2.Above(), R1_C2)
    assertCoordEquals(t, R2_C2.Below(), R3_C2)
}

func TestCoord_LeftAndRight(t *testing.T) {
    assertCoordEquals(t, R2_C2.Left(), R2_C1)
    assertCoordEquals(t, R2_C2.Right(), R2_C3)
}

func TestCoord_Chaining(t *testing.T) {
    assertCoordEquals(t, R2_C2.Right().Below(), R3_C3)
    assertCoordEquals(t, R2_C2.Below().Right(), R3_C3)
}

func TestCoord_WithRow(t *testing.T) {
    assertCoordEquals(t, R1_C1.WithRow(3), R3_C1)
    assertCoordEquals(t, R1_C1.WithColumn(3), R1_C3)
}

func TestCoord_WithNegativeIndex(t *testing.T) {
    c1 := NewCoord(-1, 0)
    assertIntEquals(t, c1.Row(), -1)
    assertIntEquals(t, c1.Column(), 0)
    c2 := NewCoord(0, -1)
    assertIntEquals(t, c2.Row(), 0)
    assertIntEquals(t, c2.Column(), -1)
    c3 := NewCoord(-1, -1)
    assertIntEquals(t, c3.Row(), -1)
    assertIntEquals(t, c3.Column(), -1)
}

func TestCoord_NavigatingAroundOrigin(t *testing.T) {
    start := NewCoord(0, 0)
    finish := start.Above().Right().Below().Below().Left().Left().Above().Above().Right().Below()
    assertCoordEquals(t, start, finish)
}

func BenchmarkNewCoord(b *testing.B) {
    for i := 0; i < b.N; i++ {
        NewCoord(i, i)
    }
}

func BenchmarkWithRow(b *testing.B) {
    c := NewCoord(0, 0)
    for i := 0; i < b.N; i++ {
        c.WithRow(i)
    }
}

func BenchmarkWithColumn(b *testing.B) {
    c := NewCoord(0, 0)
    for i := 0; i < b.N; i++ {
        c.WithColumn(i)
    }
}

func BenchmarkAbove(b *testing.B) {
    c := NewCoord(0, 0)
    for i := 0; i < b.N; i++ {
        c = c.Above()
    }
}

func BenchmarkBelow(b *testing.B) {
    c := NewCoord(0, 0)
    for i := 0; i < b.N; i++ {
        c = c.Below()
    }
}

func BenchmarkLeft(b *testing.B) {
    c := NewCoord(0, 0)
    for i := 0; i < b.N; i++ {
        c = c.Left()
    }
}

func BenchmarkRight(b *testing.B) {
    c := NewCoord(0, 0)
    for i := 0; i < b.N; i++ {
        c = c.Right()
    }
}
