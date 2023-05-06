/*
* <2023-05-06>
* hihi it's edelstine back at it again.
*
 */

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

type userPlayer struct {
	Xpos      float64
	Ypos      float64
	isJumping bool
}

type gravityStruct struct {
	constant float64
	terminal float64
}

var img *ebiten.Image

var gravLoop float64 = 0
var jumpLoop float64 = 0

var gravity gravityStruct = gravityStruct{
	constant: 0.5,
	terminal: 3,
}

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

/* UPDATE (time: just now): I just thought of a way to use external files without breaking the program bc ebiten doesn't like it when u import it twice */
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
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || user.isJumping {
		if jumpLoop <= 2 {
			x.isJumping = true
			x.Ypos = x.Ypos - jumpLoop
			time.Sleep(1 * time.Nanosecond)
			jumpLoop = jumpLoop - 0.1
		} else if jumpLoop <= 0 {
			x.isJumping = false
			jumpLoop = 0
		}
	}

	return x

}

func (x userPlayer) playerCollision() userPlayer {
	for x.Ypos >= 100 {
		x.Ypos = x.Ypos - 0.1
	}

	return x
}

func (x userPlayer) playerGravity(g gravityStruct) userPlayer {
	if x.Ypos < 100 && !x.isJumping {
		x.Ypos = x.Ypos + gravLoop
		if gravLoop <= g.terminal {
			gravLoop = gravLoop + g.constant
		}
		time.Sleep(2 * time.Nanosecond)
	}

	return x
}

/* render the guy */
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 0xff})

	op := &ebiten.DrawImageOptions{}

	user = user.playerGravity(gravity)
	user = user.playerMovement()
	user = user.playerCollision()

	// the heart
	op.GeoM.Translate(user.Xpos, user.Ypos)
	screen.DrawImage(img, op)
	// end of the heart

	ebitenutil.DebugPrint(screen, "walao")
}
