.PHONY: all clean default

default: all

all: commp2cid

%: %.go
	go build $<

clean:
	rm commp2cid
