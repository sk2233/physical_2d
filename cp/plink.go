package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
)

type Plink struct {
	Space  *cp.Space
	Drawer *Drawer
}

func NewPlink() *Plink {
	sp := cp.NewSpace()
	sp.SetGravity(cp.Vector{Y: -10})

	tris := []cp.Vector{{-1.5, -1.5}, {0, 1}, {1.5, -1.5}}
	for i := -2; i <= 2; i++ {
		for j := -6; j <= 5; j++ {
			offset := cp.Vector{X: float64(j * 6), Y: float64(i * 6)}
			if i%2 == 0 {
				offset.X += 3
			}
			tmp := sp.AddShape(cp.NewPolyShape(sp.StaticBody, 3, tris, cp.NewTransformTranslate(offset), 0))
			tmp.SetElasticity(0.8)
		}
	}

	for i := 0; i < 100; i++ {
		GenBall(sp)
	}
	return &Plink{Space: sp, Drawer: NewDrawer()}
}

func GenBall(sp *cp.Space) {
	bd := sp.AddBody(cp.NewBody(1, cp.MomentForCircle(1, 0, 1, cp.Vector{})))
	bd.SetPosition(cp.Vector{X: 60 * (rand.Float64() - 0.5), Y: 30})
	tmp := sp.AddShape(cp.NewCircle(bd, 1, cp.Vector{}))
	tmp.SetElasticity(0.8)
}

func (o *Plink) Update() error {
	o.Space.Step(1 / float64(ebiten.TPS()))
	o.Space.EachBody(func(body *cp.Body) {
		if body.Position().Y < -30 {
			o.Space.RemoveBody(body)
			GenBall(o.Space)
		}
	})
	return nil
}

func (o *Plink) Draw(screen *ebiten.Image) {
	o.Drawer.DrawSpace(screen, o.Space)
}

func (o *Plink) Layout(w, h int) (int, int) {
	return w, h
}
