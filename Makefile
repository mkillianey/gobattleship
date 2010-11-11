# Copyright (c) 2010 Mick Killianey and Ivan Moore.
# All rights reserved.  See the LICENSE file for details.

include $(GOROOT)/src/Make.inc

TARG = battleship 

PKG_DIR = pkg
CMD_DIR = cmd

GODOC_PORT=8000
GODOC = godoc -http=:$(GODOC_PORT) -path=`pwd`/pkg

all: clean test command bench

command: 
	$(MAKE) -C pkg/battleship 
	$(MAKE) -C pkg/battleship install
	$(MAKE) -C cmd/battleship clean
	$(MAKE) -C cmd/battleship

clean: 
	$(MAKE) -C pkg/battleship clean
	# Uninstalling the battleship package from $(GOROOT)/pkg
	rm -f $(GOROOT)/pkg/$(GOOS)_$(GOARCH)/battleship.a
	$(MAKE) -C cmd/battleship clean

test: 
	$(MAKE) -C pkg/battleship test 

gotest: 
	$(MAKE) -C pkg/battleship gotest 

bench:
	$(MAKE) -C pkg/battleship bench

format:
	$(MAKE) -C pkg/battleship format
	$(MAKE) -C cmd/battleship format

doc:
	echo "Running on http://localhost:$(GODOC_PORT)"
	$(GODOC)


	