package utils

import (
	"bytes"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	pixelSquareFaceSource *text.GoTextFaceSource
	PixelSquareNormalFace *text.GoTextFace
	// mplusBigFace          *text.GoTextFace
)

func InitFonts() {
	f, err := os.ReadFile("./assets/fonts/Pixel_Square.ttf")
	if err != nil {
		log.Fatal("Failed to load fonts")
	}

	s, err := text.NewGoTextFaceSource(bytes.NewReader(f))
	if err != nil {
		log.Fatal(err)
	}
	pixelSquareFaceSource = s

	PixelSquareNormalFace = &text.GoTextFace{
		Source: pixelSquareFaceSource,
		Size:   24,
	}
	// mplusBigFace = &text.GoTextFace{
	// 	Source: mplusFaceSource,
	// 	Size:   32,
	// }
}
