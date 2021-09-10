SHELL=/bin/bash

EXE = trie-server

all: $(EXE)

trie-server:
	@echo "building $@ ..."
	$(MAKE) -s -f make.inc s=static

clean:
	rm -f $(EXE)

