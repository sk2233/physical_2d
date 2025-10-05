package main

import "github.com/hajimehoshi/ebiten/v2"

// https://github.com/jakecoffman/cp

func main() {
	ebiten.SetWindowSize(ScreenW, ScreenH)
	err := ebiten.RunGame(NewTheojansen())
	HandleErr(err)
}
