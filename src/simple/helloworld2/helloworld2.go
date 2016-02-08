package main

import (
	"os"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/slapresta/dit"
)

type Triangle struct {
	Character              rune
	Foreground, Background termbox.Attribute
}

func (t Triangle) Draw(c dit.Context) {
	width, height := c.Size()
	for i := 0; i < height; i++ {
		if width%2 == 1 {
			c.SetCell((width/2)+1, i, t.Character, t.Foreground, t.Background)
		}
		fill := (width / 2) * (i + 1) / (height + 1)
		for j := width/2 - fill; j < width/2; j++ {
			c.SetCell(j+1, i, t.Character, t.Foreground, t.Background)
			c.SetCell(width-j, i, t.Character, t.Foreground, t.Background)
		}
	}
}

type VerticalFlip struct {
	dit.Drawable
}

func (f VerticalFlip) Draw(c dit.Context) {
	f.Drawable.Draw(verticalFlipContext{f, c})
}

type verticalFlipContext struct {
	VerticalFlip
	dit.Context
}

func (fc verticalFlipContext) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	_, height := fc.Size()
	fc.Context.SetCell(x, height-y-1, ch, fg, bg)
}

func makeTriangle(x, y int, color termbox.Attribute, flip bool) dit.Drawable {
	var triangle dit.Drawable
	triangle = Triangle{
		Background: color, Foreground: color,
		Character: '▄',
	}

	if flip {
		triangle = VerticalFlip{triangle}
	}

	view := dit.View{
		X: x, Y: y,
		Width: 20, Height: 10,
		Drawables: []dit.Drawable{triangle},
	}

	return view
}

func main() {
	termbox.Init()
	defer termbox.Close()

	arg := ""
	if len(os.Args) > 1 {
		arg = " " + os.Args[1]
	}

	quit := make(chan bool, 1)
	go pollForQuit(quit)

	b := dit.Box{
		X: 10, Y: 3,
		Width: 20, Height: 6,
		Background: termbox.ColorGreen, Foreground: termbox.ColorWhite,
		Content:   "Welcome to dit!\n\nPlease stay calm and enjoy the winds~",
		Alignment: dit.RightAlignment,
	}

	dit.Draw(b)

	positions := []struct {
		x    int
		y    int
		flip bool
	}{
		{0, 0, false},
		{12, 0, true},
		{24, 0, false},
		{24, 11, true},
		{12, 11, false},
		{0, 11, true},
	}

	colors := []termbox.Attribute{
		termbox.ColorCyan,
		termbox.ColorGreen,
		termbox.ColorYellow,
		termbox.ColorRed,
		termbox.ColorMagenta,
		termbox.ColorBlue,
	}

	for j := 0; ; j++ {
		group := dit.Group{X: 3, Y: 2}

		for i, position := range positions {
			color := (i + j) % len(colors)
			group.Drawables = append(group.Drawables, makeTriangle(
				position.x, position.y, colors[color], position.flip))
		}

		dit.Draw(dit.Group{Drawables: []dit.Drawable{
			dit.Text{Content: "HellowWorld(2!)" + arg},
			group,
		}})

		select {
		case <-quit:
			termbox.Close()
			os.Exit(0)
		case <-time.After(time.Millisecond * 80):
		}
	}
}

func pollForQuit(quit chan bool) {
	for {
		e := termbox.PollEvent()
		if e.Key == termbox.KeyCtrlC {
			quit <- true
		}
	}
}
