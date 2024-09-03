package utils

import (
	"fmt"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type DebugType string

const (
	DebugMouseX              DebugType = "Mouse X @"
	DebugMouseY              DebugType = "Mouse Y @"
	DebugCookies             DebugType = "Cookies"
	DebugPreviousCookieCount DebugType = "Previous Cookie Count"
)

type DebugInfo struct {
	dx   int
	dy   int
	keys []DebugType
	Info map[DebugType]interface{}
}

func NewDebugInfo(dx, dy int, info map[DebugType]interface{}) *DebugInfo {
	keys := make([]DebugType, len(info))
	for k := range info {
		keys = append(keys, k)
	}
	return &DebugInfo{
		Info: info,
		keys: keys,
		dx:   dx,
		dy:   dy,
	}
}

func (di *DebugInfo) Update() error {
	return nil
}

func (di *DebugInfo) Draw(screen *ebiten.Image) {
	yLocation := 0

	for _, k := range di.keys {
		v := di.Info[k]
		vstr := ""
		switch val := v.(type) {
		case int:
			vstr = strconv.Itoa(val)
		case bool:
			vstr = strconv.FormatBool(val)
		case string:
			vstr = val
		default:
			continue
		}
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s: %s", k, vstr), di.dx, di.dy+yLocation)
		yLocation += 10
	}
}

func (di *DebugInfo) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 200, 300
}

func (di *DebugInfo) Insert(key DebugType, value interface{}) bool {
	_, ok := di.Info[key]
	if !ok {
		di.keys = append(di.keys, key)
	}
	di.Info[key] = value

	return ok
}
