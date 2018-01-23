#!/bin/bash
DIRECTORY=$(pwd)
export GOPATH=$DIRECTORY
if [ ! -e "src/github.com/veandco/go-sdl2" ]; then
    go get github.com/veandco/go-sdl2/sdl
fi
go fmt src/*.go
go run src/*.go
