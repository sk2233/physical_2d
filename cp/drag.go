package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
)

type Drag struct {
	Space  *cp.Space
	Drawer *Drawer
	Draper *Draper
}

func NewDrag() *Drag {
	sp := cp.NewSpace()
	sp.SetGravity(cp.Vector{Y: -10})
	// 构建围墙
	tmp := sp.AddShape(cp.NewSegment(sp.StaticBody, cp.Vector{X: -SpaceW / 2, Y: -SpaceH / 2}, cp.Vector{X: SpaceW / 2, Y: -SpaceH / 2}, 0))
	tmp.SetElasticity(0.9)
	tmp = sp.AddShape(cp.NewSegment(sp.StaticBody, cp.Vector{X: SpaceW / 2, Y: -SpaceH / 2}, cp.Vector{X: SpaceW / 2, Y: SpaceH / 2}, 0))
	tmp.SetElasticity(0.9)
	tmp = sp.AddShape(cp.NewSegment(sp.StaticBody, cp.Vector{X: -SpaceW / 2, Y: -SpaceH / 2}, cp.Vector{X: -SpaceW / 2, Y: SpaceH / 2}, 0))
	tmp.SetElasticity(0.9)
	// 添加球
	bd := sp.AddBody(cp.NewBody(1, cp.MomentForCircle(1, 0, 3, cp.Vector{})))
	tmp = sp.AddShape(cp.NewCircle(bd, 3, cp.Vector{}))
	tmp.SetElasticity(0.9)
	return &Drag{Space: sp, Drawer: NewDrawer(), Draper: NewDraper()}
}

func (o *Drag) Update() error {
	o.Space.Step(1 / float64(ebiten.TPS()))
	o.Draper.Update(o.Space)
	return nil
}

func (o *Drag) Draw(screen *ebiten.Image) {
	o.Drawer.DrawSpace(screen, o.Space)
}

func (o *Drag) Layout(w, h int) (int, int) {
	return w, h
}
