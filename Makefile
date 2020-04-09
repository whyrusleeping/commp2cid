.PHONY: all clean default

default: all

all: commp2cid

commp2cid: main.go
	go build -o $@ $^

clean:
	rm -f commp2cid
