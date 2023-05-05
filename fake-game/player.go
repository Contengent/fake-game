package main

import (
	_ "fmt"

	_ "github.com/hajimehoshi/ebiten"
)

type player struct {
	Xpos float64
	Ypos float64
}

/*
func (x player) playerMovement() player {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		x.Xpos = x.Xpos + 1
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		x.Xpos = x.Xpos - 1
	} else if ebiten.IsKeyPressed(ebiten.KeyW) {
		x.Ypos = x.Ypos + 1
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		x.Ypos = x.Ypos - 1
	}

	return x

}
*/
