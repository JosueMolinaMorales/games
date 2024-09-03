package cookie_clicker

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/JosueMolinaMorales/game/cookie_clicker/entity"
	"github.com/JosueMolinaMorales/game/cookie_clicker/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	ScreenHeight = 600
	ScreenWidth  = 800
)

type Game struct {
	cookies       *int
	cookieImage   *ebiten.Image
	storeImg      *ebiten.Image
	store         *Store
	debugInfo     *utils.DebugInfo
	plusCount     []*entity.PlusCount
	mousex        int
	mousey        int
	ppc           int
	cookieClicked *utils.Timer
}

func NewGame() *Game {
	img, _, err := ebitenutil.NewImageFromFile("./assets/cookie.png")
	if err != nil {
		log.Fatal(err)
	}

	storeImg, _, err := ebitenutil.NewImageFromFile("./assets/store.png")
	if err != nil {
		log.Fatal(err)
	}

	cookies := new(int)
	*cookies = 0
	return &Game{
		cookieImage: img,
		storeImg:    storeImg,
		store:       NewStore(ScreenWidth, ScreenHeight),
		debugInfo: utils.NewDebugInfo(0, 0, map[utils.DebugType]interface{}{
			utils.DebugMouseX:              0,
			utils.DebugMouseY:              0,
			utils.DebugCookies:             0,
			utils.DebugPreviousCookieCount: 0,
		}),
		cookies: cookies,
	}
}

func (g *Game) removePlusIcons() {
}

func (g *Game) Update() error {
	g.store.Update()
	g.debugInfo.Update()
	// Get mouse position
	x, y := ebiten.CursorPosition()
	g.mousex = x
	g.mousey = y
	g.debugInfo.Insert(utils.DebugMouseX, x)
	g.debugInfo.Insert(utils.DebugMouseY, y)
	if g.cookieClicked != nil {
		g.cookieClicked.Update()
	}
	// Detect mouse click and check if it is within the cookie's bounds
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) && !g.store.IsOpen {
		maxX := g.cookieImage.Bounds().Max.X - (g.cookieImage.Bounds().Dx() / 2) + (ScreenWidth / 2)
		maxY := g.cookieImage.Bounds().Max.Y - (g.cookieImage.Bounds().Dy() / 2) + (ScreenHeight / 2)
		minX := g.cookieImage.Bounds().Min.X - (g.cookieImage.Bounds().Dx() / 2) + (ScreenWidth / 2)
		minY := g.cookieImage.Bounds().Min.Y - (g.cookieImage.Bounds().Dy() / 2) + (ScreenHeight / 2)
		if x >= minX && x <= maxX && y >= minY && y <= maxY {
			*g.cookies += g.store.PointsPerClick
			g.debugInfo.Insert(utils.DebugCookies, *g.cookies)
			g.plusCount = append(g.plusCount, entity.NewPlusCount(
				g.store.PointsPerClick,
				float64(x),
				float64(y),
			))
			g.cookieClicked = utils.NewTimer(time.Millisecond * 150)
		}
	}

	// Handle the click on the store button
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.store.ToggleStore(g.cookies)
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		maxX := g.storeImg.Bounds().Max.X - (g.storeImg.Bounds().Dx() / 4) + ScreenWidth/2
		maxY := g.storeImg.Bounds().Max.Y - (g.storeImg.Bounds().Dy() / 4) + (ScreenHeight - g.storeImg.Bounds().Dy()/4)
		minX := g.storeImg.Bounds().Min.X - (g.storeImg.Bounds().Dx() / 4) + ScreenWidth/2
		minY := g.storeImg.Bounds().Min.Y - (g.storeImg.Bounds().Dy() / 4) + (ScreenHeight - g.storeImg.Bounds().Dy()/4)
		if x >= minX && x <= maxX && y >= minY && y <= maxY {
			fmt.Println("Store pressed!")
			g.store.ToggleStore(g.cookies)
		}
	}

	// Update the timer of the plus counts
	for _, pc := range g.plusCount {
		pc.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{211, 211, 211, 0})
	// Draw the cookie image
	// Place cookie in middle of screen
	halfWidth := ScreenWidth / 2
	halfHeight := ScreenHeight / 2
	cookieImgOpts := &ebiten.DrawImageOptions{}

	// find the images pivot point
	imgWidth := g.cookieImage.Bounds().Dx()
	imgHeight := g.cookieImage.Bounds().Dy()

	halfW := float64(imgWidth / 2)
	halfH := float64(imgHeight / 2)

	cookieImgOpts.GeoM.Translate(-halfW, -halfH)
	cookieImgOpts.GeoM.Translate(float64(halfWidth), float64(halfHeight))

	screen.DrawImage(g.cookieImage, cookieImgOpts)

	// Draw another cookie image ontop of original cookie
	if g.cookieClicked != nil && !g.cookieClicked.IsReady() {
		imgopts := &colorm.DrawImageOptions{}
		imgopts.GeoM.Translate(-halfW, -halfH)
		imgopts.GeoM.Translate(float64(halfWidth), float64(halfHeight))
		cm := colorm.ColorM{}
		if g.cookieClicked.Percentage() < 0.50 {
			cm.Translate(1, 1, 1, -0.8+g.cookieClicked.Percentage())
		} else {
			cm.Translate(1.0, 1.0, 1.0, -0.3-g.cookieClicked.Percentage())
		}
		colorm.DrawImage(screen, g.cookieImage, cm, imgopts)
	}
	// Draw Store image
	storeImgOpts := &ebiten.DrawImageOptions{}
	storeImgOpts.GeoM.Scale(0.5, 0.5)

	// Find the image pivot point
	storeImgWidth := g.storeImg.Bounds().Dx() / 2
	storeImgHeight := g.storeImg.Bounds().Dy() / 2
	halfH = float64(storeImgHeight / 2)
	halfW = float64(storeImgWidth / 2)
	storeImgOpts.GeoM.Translate(-halfW, -halfH)
	storeImgOpts.GeoM.Translate(ScreenWidth/2, ScreenHeight-halfH)
	screen.DrawImage(g.storeImg, storeImgOpts)

	// Open the store only if it isnt already open
	if g.store.IsOpen {
		g.store.Draw(screen)
	}

	// Draw a +1 when a click is performed
	for _, pc := range g.plusCount {
		pc.Draw(screen)
	}

	// Draw the click amount on top
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(float64(halfWidth), 0)
	opts.ColorScale.ScaleWithColor(color.RGBA{R: 255, G: 165, B: 0})
	opts.PrimaryAlign = text.AlignCenter
	text.Draw(screen, fmt.Sprintf("%06d", *g.cookies), utils.PixelSquareNormalFace, opts)

	g.debugInfo.Draw(screen)
	// ebitenutil.DebugPrint(screen, "Cookies: "+strconv.Itoa(g.cookies))
	// ebitenutil.DebugPrintAt(screen, "ScreenWidth: "+strconv.Itoa(ScreenWidth), 0, 10)
	// ebitenutil.DebugPrintAt(screen, "Mouse x @ "+strconv.Itoa(g.mouseX), 0, 20)
	// ebitenutil.DebugPrintAt(screen, "Mouse y @ "+strconv.Itoa(g.mouseY), 0, 30)
	// ebitenutil.DebugPrintAt(screen, "Is Store opened: "+strconv.FormatBool(g.store.IsOpen), 0, 40)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
