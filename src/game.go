package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"sync"
	"time"
)

type Coordinate [2]int32

const (
	gridSize  int32  = 15
	speed     int32  = 1
	gameSpeed uint32 = 120
)

var (
	foodCoord  Coordinate
	snake      = NewSnake()
	snakeMutex = &sync.Mutex{}
)

func initializeSnake() {
	head := snake[0].(*Head)
	tail := snake[1].(*BodySegment)
	head.X, head.Y = Width/2/15, Height/2/15
	head.vX = -speed
	head.Next = tail
	tail.X, tail.Y = head.X+1, head.Y
	(&snake).AddSegment(head.X+2, head.Y)
}

func setSnakeDirection(direction string) {
	snakeMutex.Lock()
	snake.ChangeSnakeDirection(direction)
	snakeMutex.Unlock()
}

func setFood() {
	foodCoord[0] = rand.Int31n(Width / gridSize)
	foodCoord[1] = rand.Int31n(Width / gridSize)
}

//Returns array of positions, with tail first and head last
func getSnakePositions(startSegment SnakeSegment) (positions []Coordinate) {
	x, y := startSegment.Position()
	if startSegment == snake[1] {
		return []Coordinate{Coordinate{x, y}}
	}
	return append(getSnakePositions(Next(startSegment)), Coordinate{x, y})
}

func getSnakeRects(snakePositions []Coordinate) []sdl.Rect {
	snakeRects := []sdl.Rect{}
	for _, coordinate := range snakePositions {
		snakeRects = append(snakeRects, sdl.Rect{X: coordinate[0] * 15, Y: coordinate[1] * 15, W: gridSize, H: gridSize})
	}
	return snakeRects
}

func foodCollision() {
	if x, y := snake[0].Position(); x == foodCoord[0] && y == foodCoord[1] {
		snakeHead := snake[0].(*Head)
		x, y = snake[1].Position()
		(&snake).AddSegment(x-snakeHead.vX, y-snakeHead.vY)
		setFood()
	}
}

func offscreen() bool {
	x, y := snake[0].Position()
	return x < 0 || x >= Width/gridSize || y < 0 || y >= Height/gridSize
}

func selfCollision(snakePositions []Coordinate) bool {
	positions := map[Coordinate]bool{}
	for _, coordinate := range snakePositions {
		if _, exists := positions[coordinate]; exists {
			return true
		} else {
			positions[coordinate] = true
		}
	}
	return false
}

func checkGameOver(snakePositions []Coordinate) bool {
	return offscreen() || selfCollision(snakePositions)
}

func drawGame(renderer *sdl.Renderer, snakePositions []Coordinate) {
	renderer.Clear()
	renderer.SetDrawColor(200, 200, 200, 255)

	var rects = getSnakeRects(snakePositions)
	rects = append(rects, sdl.Rect{X: foodCoord[0]*gridSize + 3, Y: foodCoord[1]*gridSize + 3, W: 9, H: 9})
	renderer.FillRects(rects)
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Present()
}

func gameloop(renderer *sdl.Renderer) error {
	for gameRunning := true; gameRunning; {
		snakeMutex.Lock()
		snakePositions := getSnakePositions(snake[0])
		drawGame(renderer, snakePositions)
		snake.AdvancePosition()
		foodCollision()
		gameRunning = !checkGameOver(snakePositions)
		snakeMutex.Unlock()
		sdl.Delay(gameSpeed)
	}
	return nil
}

func runGame(renderer *sdl.Renderer, errorChannel chan error) {
	rand.Seed(time.Now().UnixNano())
	setFood()
	initializeSnake()
	if err := gameloop(renderer); err != nil {
		errorChannel <- err
	}
	errorChannel <- nil
}
