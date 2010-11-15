# Copyright (c) 2010 Mick Killianey and Ivan Moore.
# All rights reserved.  See the LICENSE file for details.

include $(GOROOT)/src/Make.inc

TARG=battleship

GOFMT=gofmt -s -spaces=true -tabindent=false -tabwidth=4
GOTEST=gotest -v

GOFILES=\
	board.go\
	boards.go\
	coord.go\
	square.go\
	solver.go\

include $(GOROOT)/src/Make.pkg

gotest:
	$(GOTEST)

format:
	$(GOFMT) -w *.go

	