package main

import (
	"log/slog"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

const (
	gameSpeed = 200 * time.Millisecond
)

type Entity interface {
	Locatable
	Drawable
	Updatable
	Collidable
	Controllable
}

type Locatable interface {
	X() int
	Y() int
}

type Drawable interface {
	Draw(tcell.Screen)
}

type Updatable interface {
	Update(int, int)
}

type Collidable interface {
	Collide([]Entity)
}

type Controllable interface {
	Velocity(int, int)
}

type Game struct {
	Entity []Entity
	// Snake  Snake
	Screen tcell.Screen
}

func (g *Game) Run() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	g.Screen.SetStyle(defStyle)
	width, height := g.Screen.Size()

	for {
		g.Screen.Clear()
		for _, entity := range g.Entity {
			entity.Update(width, height)
			entity.Collide(g.Entity)
			entity.Draw(g.Screen)
		}
		time.Sleep(gameSpeed)
		g.Screen.Show()
	}
}

type SnakePart struct {
	X int
	Y int
}

type Snake struct {
	x      int
	y      int
	XSpeed int
	YSpeed int
	Parts  []SnakePart
}

func (s *Snake) Draw(screen tcell.Screen) {
	snakeStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)

	for _, part := range s.Parts {
		screen.SetContent(part.X, part.Y, rune(len(s.Parts)), nil, snakeStyle)
	}
}

func (s *Snake) Velocity(y, x int) {
	s.XSpeed = x
	s.YSpeed = y
}

func (s *Snake) Collide(entities []Entity) {
	for _, entity := range entities {
		if entity == s {
			continue
		}

		if entity.X() == s.X() && entity.Y() == s.Y() {
			s.Parts = append(s.Parts, SnakePart{X: s.X(), Y: s.Y()})
		}
	}
}

func (s *Snake) X() int {
	return s.x
}

func (s *Snake) Y() int {
	return s.y
}

func (s *Snake) Update(width int, height int) {
	// fmt.Println(s.Parts)
	s.x = (s.x + s.XSpeed) % width
	if s.x < 0 {
		s.x += width
	}
	s.y = (s.y + s.YSpeed) % height
	if s.y < 0 {
		s.y += height
	}

	s.Parts = append(s.Parts, SnakePart{s.x, s.y})

	if len(s.Parts) > 1 {
		s.Parts = s.Parts[1:]
	}
}

type Food struct {
	x, y int
}

func (f *Food) Draw(screen tcell.Screen) {
	foodStyle := tcell.StyleDefault.Background(tcell.ColorRed).Foreground(tcell.ColorRed)
	screen.SetContent(f.x, f.y, 'F', nil, foodStyle)
}

func (f *Food) Update(width int, height int) {

}

func (f *Food) Velocity(width int, height int) {

}

func (f *Food) Collide(entities []Entity) {
	for _, entity := range entities {
		if entity == f {
			continue
		}

		// TODO this need to know where it can be placed randomly
		if entity.X() == f.x && entity.Y() == f.y {
			f.x = rand.Intn(50)
			f.y = rand.Intn(50)
		}
	}
}

func (f *Food) X() int {
	return f.x
}

func (f *Food) Y() int {
	return f.y
}

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		slog.Error("creating a new screen", "error", err)
		return
	}

	if err := screen.Init(); err != nil {
		slog.Error("initializing screen", "error", err)
		return
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)

	snake := Snake{
		x:      10,
		y:      10,
		XSpeed: 1,
		YSpeed: 0,
		Parts: []SnakePart{
			{X: 10, Y: 10},
			{X: 10, Y: 10},
			{X: 10, Y: 10},
			{X: 10, Y: 10},
			{X: 10, Y: 10},
			{X: 10, Y: 10},
			{X: 10, Y: 10},
			{X: 10, Y: 10},
			{X: 10, Y: 10},
		},
	}

	food := Food{
		x: 20,
		y: 20,
	}

	game := Game{
		Entity: []Entity{&snake, &food},
		Screen: screen,
	}

	go game.Run()

	for {
		switch event := screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				game.Screen.Fini()
				os.Exit(0)
			} else if event.Key() == tcell.KeyUp {
				snake.Velocity(-1, 0)
			} else if event.Key() == tcell.KeyDown {
				snake.Velocity(1, 0)
			} else if event.Key() == tcell.KeyLeft {
				snake.Velocity(0, -1)
			} else if event.Key() == tcell.KeyRight {
				snake.Velocity(0, 1)
			}
		}

	}
}
