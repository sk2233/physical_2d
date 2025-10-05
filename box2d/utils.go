package main

import (
	"image/color"

	b2 "github.com/oliverbestmann/box2d-go"
)

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ToScreenPos(x, y float32) (float32, float32) { // 注意y 方向上上下颠倒
	x = (x + SpaceW/2) * ScreenW / SpaceW
	y = (-y + SpaceH/2) * ScreenH / SpaceH
	return x, y
}

func ToScreenLen(val float32) float32 {
	return val * ScreenW / SpaceW
}

func ToSpacePos(x, y float32) (float32, float32) { // 注意y 方向上上下颠倒
	x = x*SpaceW/ScreenW - SpaceW/2
	y = SpaceH/2 - y*SpaceH/ScreenH
	return x, y
}

func ToColor(clr b2.HexColor) color.RGBA {
	return color.RGBA{
		R: uint8(clr & 0xFF),
		G: uint8((clr >> 8) & 0xFF),
		B: uint8((clr >> 16) & 0xFF),
		A: 0xFF,
	}
}

func NewWorld() b2.World {
	def := b2.DefaultWorldDef()
	def.Gravity = b2.Vec2{Y: -10}
	return b2.CreateWorld(def)
}

func NewBox(world b2.World, pos0, pos1 b2.Vec2) {
	def := b2.DefaultBodyDef() // XxxDef 就是用来创建 Xxx 的参数集合
	def.Position = b2.Vec2{X: (pos0.X + pos1.X) / 2, Y: (pos0.Y + pos1.Y) / 2}
	body := world.CreateBody(def)
	// 可以添加多个形状信息
	body.CreatePolygonShape(b2.DefaultShapeDef(), b2.MakeBox((pos1.X-pos0.X)/2, (pos1.Y-pos0.Y)/2))
}
