GOPATH:=$(CURDIR)
export GOPATH

all: build

fmt:
	gofmt -l -w -s src/

test:
	go test -v app/common
	go test -v app/iowrapper

build:dep
	go build -o bin/bonus-conf src/app/main.go

clean:
	rm -rf pkg
	rm -rf status
	rm -rf output

dep:fmt