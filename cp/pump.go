package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
)

type Pump struct {
	Space  *cp.Space
	Drawer *Drawer
	Draper *Draper
}

func NewPump() *Pump {
	sp := cp.NewSpace()
	sp.SetGravity(cp.Vector{Y: -10})
	// 构建围墙
	tmp := sp.AddShape(cp.NewSegment(sp.StaticBody, cp.Vector{X: -SpaceW / 2, Y: -SpaceH / 2}, cp.Vector{X: SpaceW / 2, Y: -SpaceH / 2}, 0))
	tmp.SetElasticity(0.9)
	tmp.SetFriction(0.5)
	tmp = sp.AddShape(cp.NewSegment(sp.StaticBody, cp.Vector{X: SpaceW / 2, Y: -SpaceH / 2}, cp.Vector{X: SpaceW / 2, Y: SpaceH / 2}, 0))
	tmp.SetElasticity(0.9)
	tmp.SetFriction(0.5)
	tmp = sp.AddShape(cp.NewSegment(sp.StaticBody, cp.Vector{X: -SpaceW / 2, Y: -SpaceH / 2}, cp.Vector{X: -SpaceW / 2, Y: SpaceH / 2}, 0))
	tmp.SetElasticity(0.9)
	tmp.SetFriction(0.5)

	tmp = sp.AddShape(cp.NewSegment(sp.StaticBody, cp.Vector{X: -2.5, Y: 0}, cp.Vector{X: -2.5, Y: -SpaceH / 2}, 0))
	tmp = sp.AddShape(cp.NewSegment(sp.StaticBody, cp.Vector{X: 2.5, Y: 0}, cp.Vector{X: 2.5, Y: -SpaceH / 2}, 0))

	bd1 := sp.AddBody(cp.NewBody(1, cp.MomentForBox(1, 5, 2)))
	bd1.SetPosition(cp.Vector{Y: -10})
	sp.AddShape(cp.NewBox(bd1, 5, 2, 0))
	bd2 := sp.AddBody(cp.NewBody(1, cp.MomentForBox(1, 5, 2)))
	bd2.SetPosition(cp.Vector{Y: -15})
	sp.AddShape(cp.NewBox(bd2, 5, 2, 0))
	sp.AddConstraint(cp.NewDampedSpring(bd1, bd2, cp.Vector{X: -2}, cp.Vector{X: -2}, 10, 100, 0.7))
	sp.AddConstraint(cp.NewDampedSpring(bd1, bd2, cp.Vector{X: 2}, cp.Vector{X: 2}, 10, 100, 0.7))
	return &Pump{Space: sp, Drawer: NewDrawer(), Draper: NewDraper()}
}

func (o *Pump) Update() error {
	o.Space.Step(1 / float64(ebiten.TPS()))
	o.Draper.Update(o.Space)
	return nil
}

func (o *Pump) Draw(screen *ebiten.Image) {
	o.Drawer.DrawSpace(screen, o.Space)
}

func (o *Pump) Layout(w, h int) (int, int) {
	return w, h
}
