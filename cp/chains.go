package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jakecoffman/cp/v2"
)

type Chains struct {
	Space  *cp.Space
	Body   *cp.Body
	Drawer *Drawer
}

func NewChains() *Chains {
	sp := cp.NewSpace()
	sp.SetGravity(cp.Vector{Y: -10})
	// 设置物体
	bd := sp.AddBody(cp.NewBody(1, cp.MomentForCircle(1, 0, 5, cp.Vector{})))
	sp.AddShape(cp.NewCircle(bd, 5, cp.Vector{}))
	sj := cp.NewSlideJoint(bd, sp.StaticBody, cp.Vector{}, cp.Vector{Y: 20}, 0, 35)
	// 设置最大力与后置处理函数
	sj.SetMaxForce(1500)
	sj.PostSolve = func(c *cp.Constraint, space *cp.Space) {
		force := c.Class.GetImpulse() / space.TimeStep()
		if force > c.MaxForce() { // 大于最大力度在 space step 后移除约束
			space.AddPostStepCallback(func(space *cp.Space, key any, data any) {
				space.RemoveConstraint(data.(*cp.Constraint))
			}, nil, c)
		}
	}

	sp.AddConstraint(sj)
	return &Chains{Space: sp, Body: bd, Drawer: NewDrawer()}
}

func (o *Chains) Update() error {
	o.Space.Step(1 / float64(ebiten.TPS()))
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		o.Body.ApplyImpulseAtLocalPoint(cp.Vector{X: 10}, cp.Vector{X: 10})
	}
	return nil
}

func (o *Chains) Draw(screen *ebiten.Image) {
	o.Drawer.DrawSpace(screen, o.Space)
}

func (o *Chains) Layout(w, h int) (int, int) {
	return w, h
}
