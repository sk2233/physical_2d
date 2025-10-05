package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/jakecoffman/cp/v2"
)

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ToColor(clr cp.FColor) color.Color {
	return color.RGBA{
		R: uint8(clr.R * 255),
		G: uint8(clr.G * 255),
		B: uint8(clr.B * 255),
		A: uint8(clr.A * 255),
	}
}

func ToScreenPos(x float32, y float32) (float32, float32) {
	x = (x + SpaceW/2) * ScreenW / SpaceW
	y = (SpaceH/2 - y) * ScreenH / SpaceH
	return x, y
}

func ToScreenLen(val float32) float32 {
	return val * ScreenW / SpaceW
}

func ToSpacePos(x, y float32) (float32, float32) {
	x = x*SpaceW/ScreenW - SpaceW/2
	y = SpaceH/2 - y*SpaceH/ScreenH
	return x, y
}

func DrawBB(screen *ebiten.Image, bb cp.BB, clr color.Color) {
	x0, y0 := ToScreenPos(float32(bb.L), float32(bb.T))
	x1, y1 := ToScreenPos(float32(bb.R), float32(bb.B))
	vector.StrokeLine(screen, x0, y0, x1, y0, 1, clr, false)
	vector.StrokeLine(screen, x0, y1, x1, y1, 1, clr, false)
	vector.StrokeLine(screen, x0, y0, x0, y1, 1, clr, false)
	vector.StrokeLine(screen, x1, y0, x1, y1, 1, clr, false)
}
