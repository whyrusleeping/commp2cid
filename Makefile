.PHONY: all clean default setup

default: all

all: commp2cid

setup:
	asdf install
	asdf reshim

commp2cid: main.go
	go build -o $@ $^

clean:
	rm -f commp2cid
