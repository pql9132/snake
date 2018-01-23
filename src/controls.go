package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func handleKeyPressEvent(event *sdl.KeyboardEvent) {
	switch event.Keysym.Sym {
	case sdl.K_LEFT:
		setSnakeDirection("left")
	case sdl.K_RIGHT:
		setSnakeDirection("right")
	case sdl.K_UP:
		setSnakeDirection("up")
	case sdl.K_DOWN:
		setSnakeDirection("down")
	}
}
