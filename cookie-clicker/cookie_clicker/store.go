package cookie_clicker

import (
	"fmt"
	"image/color"
	"math"

	"github.com/JosueMolinaMorales/game/cookie_clicker/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Store struct {
	PointsPerClick  int
	PointsPerSecond int
	IsOpen          bool
	upgrades        []Upgrade
	CookieCount     *int
	storeImage      *ebiten.Image
	dx              float64
	dy              float64
	debugInfo       *utils.DebugInfo
}

func NewStore(screenWidth, screenHeight int) *Store {
	upgrades := []Upgrade{
		NewClickUpgrade(10, 2, 5),
		NewClickUpgrade(100, 10, 10),
		NewClickUpgrade(500, 50, 100),
	}
	storeImg := ebiten.NewImage(screenWidth/2, screenHeight/2)
	storeImg.Fill(color.RGBA{245, 245, 220, 255}) // Light beige

	// Draw a rectange
	dx := float64(screenWidth / 4)
	dy := float64(screenHeight / 4)

	debugInfo := utils.NewDebugInfo(int(dx), int(dy), map[utils.DebugType]interface{}{
		utils.DebugMouseX: 0,
		utils.DebugMouseY: 0,
	})

	return &Store{
		IsOpen:          false,
		upgrades:        upgrades,
		storeImage:      storeImg,
		dx:              dx,
		dy:              dy,
		debugInfo:       debugInfo,
		PointsPerClick:  1,
		PointsPerSecond: 0,
	}
}

func (s *Store) ToggleStore(cookies *int) {
	s.IsOpen = !s.IsOpen
	s.CookieCount = cookies
}

func (s *Store) CloseStore() {
	s.IsOpen = false
}

func (s *Store) Update() error {
	if !s.IsOpen {
		return nil
	}
	// Detect if a click happens outside the store
	x, y := ebiten.CursorPosition()
	s.debugInfo.Insert(utils.DebugMouseX, x)
	s.debugInfo.Insert(utils.DebugMouseY, y)
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		maxX := s.storeImage.Bounds().Max.X + int(s.dx)
		maxY := s.storeImage.Bounds().Max.Y + int(s.dy)
		minX := s.storeImage.Bounds().Min.X + int(s.dx)
		minY := s.storeImage.Bounds().Min.Y + int(s.dy)
		if x < minX || x > maxX || y < minY || y > maxY {
			s.IsOpen = false
		}
	}

	for i, key := range []ebiten.Key{ebiten.Key1, ebiten.Key2, ebiten.Key3} {
		if inpututil.IsKeyJustPressed(key) && *s.CookieCount >= s.upgrades[i].Cost() {
			s.upgrades[i].Apply(&s.PointsPerClick, s.CookieCount)
		}
	}

	return nil
}

func (s *Store) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(s.dx, s.dy)
	screen.DrawImage(s.storeImage, opts)

	textOpts := &text.DrawOptions{}
	textOpts.GeoM.Translate(2*s.dx, s.dy+10)
	textOpts.ColorScale.ScaleWithColor(color.Black)
	textOpts.PrimaryAlign = text.AlignCenter
	text.Draw(screen, "SHOP", utils.PixelSquareNormalFace, textOpts)

	for i, upgrade := range s.upgrades {
		textOpts = &text.DrawOptions{}
		textOpts.GeoM.Translate(s.dx+40, s.dy+float64(50+(30*i)))
		textOpts.ColorScale.ScaleWithColor(color.Black)
		text.Draw(screen, fmt.Sprintf("INC CLICK: %02d (Press %d)", upgrade.Cost(), i+1), utils.PixelSquareNormalFace, textOpts)
	}

	s.debugInfo.Draw(screen)
}

func (s *Store) Layout(outsideWidth, outsideHeight int) (width, height int) {
	// The store should be half the size of the parent screen
	return outsideWidth / 2, outsideHeight / 2
}

type Upgrade interface {
	Cost() int
	Apply(curr *int, total *int)
}

type ClickUpgrade struct {
	cost int
	inc  int
	rate float64
}

func NewClickUpgrade(cost, inc int, rate float64) *ClickUpgrade {
	return &ClickUpgrade{
		cost,
		inc,
		rate,
	}
}
func (cu *ClickUpgrade) Cost() int { return cu.cost }
func (cu *ClickUpgrade) Apply(curr *int, total *int) {
	*curr += cu.inc
	*total -= cu.cost

	cu.cost += int(math.Ceil(float64(cu.cost) * cu.rate))
}
