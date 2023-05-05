package main

import (
	// go packages
	"image/color"
	_ "image/png"
	"log"

	// ebiten packages
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

var img *ebiten.Image

var user player = player{
	Xpos: 25,
	Ypos: 50,
}

func init() {

	var err error

	img, _, err = ebitenutil.NewImageFromFile("player.png")
	if err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 0xff})

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(user.Xpos, user.Ypos)
	screen.DrawImage(img, op)

	//user = user.playerMovement()

	ebitenutil.DebugPrint(screen, "walao")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Movement demo")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
