/*
* <2023-05-06>
* hihi it's edelstine back at it again.
*
 */

package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var img *ebiten.Image

var userPlayer Player = Player{
	name:     "fartMan",
	position: drawingParams{0, 0, 0, 0, nil, nil},
}

func init() {
	var err error

	userPlayer.position.sprite, _, err = ebitenutil.NewImageFromFile("player.png")

	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

/* Main Game loop */
func (g *Game) Draw(screen *ebiten.Image) {

	userPlayer.updatePosition()
	userPlayer.gravityCalculation()
	userPlayer.collisionCheck()
	userPlayer.userInputs()

	screen.DrawImage(userPlayer.position.sprite, userPlayer.position.options)
}

//////////////////
//      pp      //
/* No clue lmfao*/
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Fake game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
