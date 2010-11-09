# Copyright (c) 2010 Mick Killianey and Ivan Moore.
# All rights reserved.  See the LICENSE file for details.

include $(GOROOT)/src/Make.inc

TARG=battleship

GOFMT=gofmt -s -spaces=true -tabindent=false -tabwidth=4

GOFILES=\
	main.go\

include $(GOROOT)/src/Make.cmd

format:
	$(GOFMT) -w main.go
