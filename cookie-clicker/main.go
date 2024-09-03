package main

import (
	_ "image/png"
	"log"

	"github.com/JosueMolinaMorales/game/cookie_clicker"
	"github.com/JosueMolinaMorales/game/cookie_clicker/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := cookie_clicker.NewGame()
	ebiten.SetWindowSize(cookie_clicker.ScreenWidth, cookie_clicker.ScreenHeight)
	ebiten.SetWindowTitle("Cookie Clicker")
	utils.InitFonts()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
