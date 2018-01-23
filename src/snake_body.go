package main

type Snake [2]SnakeSegment //[Head,Tail]

type SnakeSegment interface {
	Position() (int32, int32)
	Move(int32, int32)
}

type Head struct {
	X    int32
	Y    int32
	vX   int32
	vY   int32
	Next *BodySegment
}

type BodySegment struct {
	X    int32
	Y    int32
	Next *BodySegment
}

func (h *Head) Position() (x, y int32) {
	return h.X, h.Y
}

func (h *Head) Move(newX, newY int32) {
	h.X = newX
	h.Y = newY
}

func (b *BodySegment) Position() (x, y int32) {
	return b.X, b.Y
}

func (b *BodySegment) Move(newX, newY int32) {
	b.X = newX
	b.Y = newY
}

func Next(s SnakeSegment) (nextSegment SnakeSegment) {
	switch segment := s.(type) {
	case *Head:
		return segment.Next
	case *BodySegment:
		return segment.Next
	default:
		return
	}
}

func NewSnake() (snake Snake) {
	snake[0] = &Head{}
	snake[1] = &BodySegment{}
	return
}

func (s *Snake) AddSegment(x, y int32) {
	currentTail := s[1].(*BodySegment)
	newTail := &BodySegment{X: x, Y: y}
	currentTail.Next = newTail
	s[1] = SnakeSegment(newTail)
}

func (s *Snake) moveSnakeSegment(segment SnakeSegment, newX, newY int32) {
	currentX, currentY := segment.Position()
	segment.Move(newX, newY)
	if segment != s[1] {
		s.moveSnakeSegment(SnakeSegment(Next(segment)), currentX, currentY)
	}
}

func (s *Snake) AdvancePosition() {
	currentHeadX, currentHeadY := s[0].Position()
	s[0].Move(currentHeadX+s[0].(*Head).vX, currentHeadY+s[0].(*Head).vY)
	s.moveSnakeSegment(SnakeSegment(Next(s[0])), currentHeadX, currentHeadY)
}

func (s *Snake) ChangeSnakeDirection(direction string) {
	snakeHead := s[0].(*Head)
	switch direction {
	case "left":
		snakeHead.vX, snakeHead.vY = -1, 0
	case "right":
		snakeHead.vX, snakeHead.vY = 1, 0
	case "up":
		snakeHead.vX, snakeHead.vY = 0, -1
	case "down":
		snakeHead.vX, snakeHead.vY = 0, 1
	}
}
