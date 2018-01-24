#!/bin/bash
DIRECTORY=$(pwd)
export GOPATH=$DIRECTORY
export GOBIN=$DIRECTORY/bin
if [ ! -e "src/github.com/veandco/go-sdl2" ]; then
    go get github.com/veandco/go-sdl2/sdl
fi
go fmt src/*.go
if [ $# -eq 1 ]; then
	if [ "$1" == "install" ]; then
		go build -o bin/snake src/*.go
	else
		echo "Invalid argument: $1"
	fi
else
	go run src/*.go
fi