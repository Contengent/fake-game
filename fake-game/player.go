package main

import (
	_ "math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	name     string
	position drawingParams
}

type drawingParams struct {
	xPos float64
	yPos float64

	xVelocity float64
	yVelocity float64

	sprite  *ebiten.Image
	options *ebiten.DrawImageOptions
}

func (p *Player) updatePosition() {

	p.position.options = &ebiten.DrawImageOptions{}

	p.position.xPos += p.position.xVelocity
	p.position.yPos += p.position.yVelocity

	p.position.options.GeoM.Translate(p.position.xPos, p.position.yPos)

}

func (p *Player) gravityCalculation() {
	p.position.yVelocity = 2.6
}

func (p *Player) collisionCheck() {
	if p.position.yPos > 300 {
		p.position.yPos = 300
	}

	if p.position.xPos < 0 {
		p.position.xVelocity = 0
		p.position.xPos = 0
	} else if p.position.xPos > 600 {
		p.position.xVelocity = 0
		p.position.xPos = 600
	}
}

func (p *Player) userInputs() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.position.xVelocity = -4
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.position.xVelocity = 4
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.position.yVelocity += 10
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.position.yVelocity -= 1
	} else {
		p.position.xVelocity = 0
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.position.yVelocity -= 40
	}
}
