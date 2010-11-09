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
