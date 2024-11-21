package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

// tile_color would be amount * XX
// func () CreateTile(amount, tile_color) {

// }

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// screen.Draw()
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(480, 640)
	ebiten.SetWindowTitle("2048")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
