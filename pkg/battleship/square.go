// Copyright (c) 2010 Mick Killianey and Ivan Moore.
// All rights reserved.  See the LICENSE file for details.

package battleship

import (
    "fmt"
)

type Square int

const (
    UNSOLVED = Square(iota)
    WATER
    TOP
    BOTTOM
    LEFT
    RIGHT
    SINGLE
    MIDDLE
    OUT_OF_BOUNDS
)

// The solved squares
var SQUARES = []Square{WATER, TOP, BOTTOM, LEFT, RIGHT, SINGLE, MIDDLE}

var SQUARE_NAMES = map[Square]string{
    UNSOLVED:      "UNSOLVED",
    WATER:         "WATER",
    TOP:           "TOP",
    BOTTOM:        "BOTTOM",
    LEFT:          "LEFT",
    RIGHT:         "RIGHT",
    SINGLE:        "SINGLE",
    MIDDLE:        "MIDDLE",
    OUT_OF_BOUNDS: "OUT_OF_BOUNDS",
}


func (square Square) String() string {
    if name, ok := SQUARE_NAMES[square]; ok {
        return name
    }
    return fmt.Sprintf("Unrecognized square: %v", int(square))
}

func (square Square) IsShip() bool {
    switch square {
    case TOP,
        BOTTOM,
        LEFT,
        RIGHT,
        MIDDLE,
        SINGLE:
        return true
    }
    return false
}

func (square Square) IsWater() bool {
    return square == WATER
}

func (square Square) IsUnsolved() bool {
    return square == UNSOLVED
}

func (square Square) IsOutOfBounds() bool {
    return square == OUT_OF_BOUNDS
}

func (this Square) CanAppearAbove(that Square) bool {
    if this.IsUnsolved() || that.IsUnsolved() {
        return true
    }
    switch this {
    case WATER, OUT_OF_BOUNDS:
        return that != BOTTOM
    case LEFT, RIGHT, BOTTOM, SINGLE:
        return that == WATER || that == OUT_OF_BOUNDS
    case TOP:
        return that == MIDDLE || that == BOTTOM
    case MIDDLE:
        return that == MIDDLE || that == BOTTOM || that == WATER || that == OUT_OF_BOUNDS
    }
    return false
}

func (this Square) CanAppearBelow(that Square) bool {
    return that.CanAppearAbove(this)
}

func (this Square) CanAppearLeftOf(that Square) bool {
    if this.IsUnsolved() || that.IsUnsolved() {
        return true
    }
    switch this {
    case WATER, OUT_OF_BOUNDS:
        return that != RIGHT
    case TOP, BOTTOM, RIGHT, SINGLE:
        return that.IsWater() || that.IsOutOfBounds()
    case LEFT:
        return that == MIDDLE || that == RIGHT
    case MIDDLE:
        return that == MIDDLE || that == RIGHT || that.IsWater() || that.IsOutOfBounds()
    }
    return false
}

func (this Square) CanAppearRightOf(that Square) bool {
    return that.CanAppearLeftOf(this)
}

func (this Square) CanAppearDiagonallyAdjacentTo(that Square) bool {
    return !(this.IsShip() && that.IsShip())
}
