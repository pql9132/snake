package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"os"
	"sync"
)

const (
	Width  int32 = 450
	Height int32 = 450
)

var (
	Running bool = true
)

func handleEvents(e sdl.Event) {
	switch event := e.(type) {
	case *sdl.QuitEvent:
		Running = false
	case *sdl.KeyboardEvent:
		if event.State == sdl.PRESSED {
			handleKeyPressEvent(event)
		}
	}
}

func run() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Snake", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, Width, Height, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		return err
	}
	defer renderer.Destroy()

	errorChannel := make(chan error)
	runMutex := &sync.Mutex{}

	go runGame(renderer, errorChannel)

	//Check for errors from gameloop
	go func() {
		err := <-errorChannel
		if err != nil {
			runMutex.Lock()
			fmt.Fprintln(os.Stderr, err)
			Running = false
			runMutex.Unlock()
		}
	}()

	//Window loop
	for Running {
		runMutex.Lock()
		handleEvents(sdl.PollEvent())
		runMutex.Unlock()
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
