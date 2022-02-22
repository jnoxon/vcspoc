GO=go1.18

all: clean vcspoc1 vcspoc2

vcspoc1: vcspoc.go
	$(GO) build vcspoc.go
	./vcspoc

vcspoc2: vcspoc.go
	$(GO) build
	./vcspoc

clean:
	rm -f vcspoc
