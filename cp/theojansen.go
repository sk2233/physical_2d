package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
)

type Theojansen struct {
	Space  *cp.Space
	Drawer *Drawer
	Draper *Draper
}

func NewTheojansen() *Theojansen {
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

	bd1 := sp.AddBody(cp.NewBody(2, cp.MomentForBox(2, 20, 2)))
	sp1 := sp.AddShape(cp.NewBox(bd1, 10, 1, 0))
	sp1.SetFilter(cp.NewShapeFilter(1, cp.ALL_CATEGORIES, cp.ALL_CATEGORIES))
	bd2 := sp.AddBody(cp.NewBody(2, cp.MomentForCircle(2, 0, 3, cp.Vector{})))
	sp2 := sp.AddShape(cp.NewCircle(bd2, 3, cp.Vector{}))
	sp2.SetFilter(cp.NewShapeFilter(1, cp.ALL_CATEGORIES, cp.ALL_CATEGORIES))
	sp.AddConstraint(cp.NewPivotJoint2(bd1, bd2, cp.Vector{}, cp.Vector{}))

	for i := 0; i < 3; i++ {
		makeLeg(sp, -10, bd1, bd2, 2*math.Pi*float32(i*2)/3/2)
		makeLeg(sp, 10, bd1, bd2, 2*math.Pi*float32(i*2+1)/3/2)
	}
	return &Theojansen{Space: sp, Drawer: NewDrawer(), Draper: NewDraper()}
}

func makeLeg(sp *cp.Space, offset int, bd1 *cp.Body, bd2 *cp.Body, angle float32) {

}

func (o *Theojansen) Update() error {
	o.Space.Step(1 / float64(ebiten.TPS()))
	o.Draper.Update(o.Space)
	return nil
}

func (o *Theojansen) Draw(screen *ebiten.Image) {
	o.Drawer.DrawSpace(screen, o.Space)
}

func (o *Theojansen) Layout(w, h int) (int, int) {
	return w, h
}
