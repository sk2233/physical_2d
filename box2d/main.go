package main

import "github.com/hajimehoshi/ebiten/v2"

// 长度单位m(米) 总量单位k(千克) 时间单位s(秒)

func main() {
	ebiten.SetWindowSize(ScreenW, ScreenH)
	err := ebiten.RunGame(NewGame())
	HandleErr(err)
}
