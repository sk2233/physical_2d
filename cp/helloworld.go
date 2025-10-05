package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jakecoffman/cp/v2"
)

type HelloWorld struct {
	Space  *cp.Space
	Drawer *Drawer
}

func NewHelloWorld() *HelloWorld {
	// 初始化空间
	sp := cp.NewSpace()
	sp.SetGravity(cp.Vector{Y: -100})
	// 初始化地板
	gd := cp.NewSegment(sp.StaticBody, cp.Vector{X: -20, Y: -16}, cp.Vector{X: 20, Y: -15}, 0)
	gd.SetFriction(1)
	gd.SetElasticity(1)
	sp.AddShape(gd)
	// 初始化球
	m := cp.MomentForCircle(1, 0, 5, cp.Vector{}) // 根据形状与质量计算转动惯性
	bd1 := sp.AddBody(cp.NewBody(1, m))
	bd1.SetPosition(cp.Vector{X: 10, Y: 15})
	bs := sp.AddShape(cp.NewCircle(bd1, 5, cp.Vector{}))
	bs.SetFriction(0.7)
	bs.SetElasticity(0.8)
	bd2 := sp.AddBody(cp.NewBody(1, m))
	bd2.SetPosition(cp.Vector{X: -10, Y: 15})
	bs = sp.AddShape(cp.NewCircle(bd2, 5, cp.Vector{}))
	bs.SetFriction(0.7)
	bs.SetElasticity(0.8)
	sj := cp.NewSlideJoint(bd1, bd2, cp.Vector{}, cp.Vector{}, 5, 15)
	sp.AddConstraint(sj)
	return &HelloWorld{Space: sp, Drawer: NewDrawer()}
}

func (o *HelloWorld) Update() error {
	o.Space.Step(1 / float64(ebiten.TPS()))
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		tx, ty := ebiten.CursorPosition()
		x, y := ToSpacePos(float32(tx), float32(ty))
		info := o.Space.PointQueryNearest(cp.Vector{X: float64(x), Y: float64(y)}, 5, cp.ShapeFilter{})
		if info.Shape != nil {
			fmt.Println(info)
		}
	}
	return nil
}

func (o *HelloWorld) Draw(screen *ebiten.Image) {
	o.Drawer.DrawSpace(screen, o.Space)
}

func (o *HelloWorld) Layout(w, h int) (int, int) {
	return w, h
}
