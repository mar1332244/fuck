.PHONY: all
all:
	./build

run:
	./bin/main

.PHONY: clean
clean:
	rm -f ./bin/*
