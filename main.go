package main

import (
	// go packages
	"image/color"
	_ "image/png"
	"log"
	"time"

	// ebiten packages
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct{}

var img *ebiten.Image

var jumpLoop int = 0
var user userPlayer = userPlayer{
	Xpos:      25,
	Ypos:      50,
	isJumping: false,
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

	user = user.playerCollision()
	user = user.playerGravity()
	user = user.playerMovement()

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

/* Player stuffs */
type userPlayer struct {
	Xpos      float64
	Ypos      float64
	isJumping bool
}

// WORK IN PROGRESS
/*
func (x userPlayer) thePlayerFunction() userPlayer {
	user.playerMovement()
	user.playerCollision()

	return x
}
*/

func (x userPlayer) playerMovement() userPlayer {
	// left-right movement
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		x.Xpos = x.Xpos + 1
	} else if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		x.Xpos = x.Xpos - 1
	}

	// for debugging?
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		x.Ypos = x.Ypos + 1
	} else if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		x.Ypos = x.Ypos - 1
	}

	// jump
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		for jumpLoop <= 25 {
			x.Ypos = x.Ypos - 1
			jumpLoop++
		}
		jumpLoop = 0
	}

	return x

}

func (x userPlayer) playerCollision() userPlayer {
	for x.Ypos > 100 {
		x.Ypos = x.Ypos - 1
		time.Sleep(1 * time.Second)
	}

	return x
}

func (x userPlayer) playerGravity() userPlayer {
	for x.Ypos < 100 {
		x.Ypos = x.Ypos + 0.13
	}

	return x
}
