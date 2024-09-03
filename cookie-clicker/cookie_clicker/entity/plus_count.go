package entity

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/JosueMolinaMorales/game/cookie_clicker/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Quadrant = int

const (
	UpperLeft Quadrant = iota
	UpperRight
	BottomLeft
	BottomRight
	Center
)

// TODO: Rename
var Quads = []Quadrant{
	UpperLeft,
	UpperRight,
	BottomLeft,
	BottomRight,
}

type PlusCount struct {
	Dx    float64
	Dy    float64
	Count int
	Timer *utils.Timer
	Quad  Quadrant
}

func NewPlusCount(count int, dx, dy float64) *PlusCount {
	return &PlusCount{
		Count: count,
		Timer: utils.NewTimer(time.Second * 2),
		Dx:    dx,
		Dy:    dy,
		Quad:  Quads[rand.Intn(len(Quads))],
	}
}

func (pc *PlusCount) Update() error {
	pc.Timer.Update()
	// Every tick move the icon to the right
	switch pc.Quad {
	case UpperLeft:
		pc.Dx += -1.5
		pc.Dy += -1.5
	case UpperRight:
		pc.Dx += 1.5
		pc.Dy += 1.5
	case BottomLeft:
		pc.Dx += -1.5
		pc.Dy += 1.5
	case BottomRight:
		pc.Dx += 1.5
		pc.Dy += -1.5
	}
	return nil
}

func (pc *PlusCount) Draw(screen *ebiten.Image) {
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(pc.Dx, pc.Dy)
	opts.ColorScale.ScaleWithColor(color.White)
	if !pc.Timer.IsReady() {
		text.Draw(screen, fmt.Sprintf("+%d", pc.Count), utils.PixelSquareNormalFace, opts)
	}
}

func (pc *PlusCount) Layout(outsideWidth, outsideHeight int) (width, height int) {
	return outsideWidth, outsideHeight
}
