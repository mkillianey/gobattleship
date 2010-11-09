# Copyright (c) 2010 Mick Killianey and Ivan Moore.
# All rights reserved.  See the LICENSE file for details.

TARG = battleship 

PKG_DIR = pkg
CMD_DIR = cmd

all: 
	$(MAKE) -C pkg/battleship 
	$(MAKE) -C pkg/battleship install
	$(MAKE) -C cmd/battleship

clean: 
	$(MAKE) -C pkg/battleship clean 
	$(MAKE) -C cmd/battleship clean

test: 
	$(MAKE) -C pkg/battleship test

format:
	$(MAKE) -C pkg/battleship format
	$(MAKE) -C cmd/battleship format
	