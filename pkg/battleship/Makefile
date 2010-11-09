# Copyright (c) 2010 Mick Killianey and Ivan Moore.
# All rights reserved.  See the LICENSE file for details.

include $(GOROOT)/src/Make.inc

TARG=battleship

GOFMT=gofmt -s -spaces=true -tabindent=false -tabwidth=4

GOFILES=\
	battleship.go\
	boards.go\

include $(GOROOT)/src/Make.pkg

format:
	$(GOFMT) -w battleship.go
	$(GOFMT) -w battleship_test.go
	$(GOFMT) -w boards.go

	