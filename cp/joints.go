package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
)

type Joints struct {
	Space  *cp.Space
	Drawer *Drawer
	Draper *Draper
}

func NewJoints() *Joints {
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
	// 准备两个球
	bd1 := sp.AddBody(cp.NewBody(1, cp.MomentForCircle(1, 0, 3, cp.Vector{})))
	bd1.SetPosition(cp.Vector{X: -10, Y: 0})
	sp1 := sp.AddShape(cp.NewCircle(bd1, 3, cp.Vector{}))
	sp1.SetFriction(0.7)
	sp1.SetFilter(cp.NewShapeFilter(1, cp.ALL_CATEGORIES, cp.ALL_CATEGORIES))
	bd2 := sp.AddBody(cp.NewBody(1, cp.MomentForCircle(1, 0, 3, cp.Vector{})))
	bd2.SetPosition(cp.Vector{X: 10, Y: 0})
	sp2 := sp.AddShape(cp.NewCircle(bd2, 3, cp.Vector{}))
	sp2.SetFriction(0.7)
	sp2.SetFilter(cp.NewShapeFilter(1, cp.ALL_CATEGORIES, cp.ALL_CATEGORIES))
	// NewPinJoint 固定两个对象的两个位置
	//con := cp.NewPinJoint(bd1, bd2, cp.Vector{}, cp.Vector{})
	// NewSlideJoint 与 NewPinJoint 类似不过有一定容忍
	//con := cp.NewSlideJoint(bd1, bd2, cp.Vector{}, cp.Vector{}, 10, 20)
	// NewPivotJoint2 两个对象只能绕着其局部坐标下的某一点旋转
	//con := cp.NewPivotJoint2(bd1, bd2, cp.Vector{X: 10}, cp.Vector{X: -10})
	// NewGrooveJoint 与 NewPivotJoint2 类似不过一个锚点在线上
	//con := cp.NewGrooveJoint(bd1, bd2, cp.Vector{X: 10, Y: -10}, cp.Vector{X: 10, Y: 10}, cp.Vector{X: -10})
	// NewDampedSpring 弹簧
	//con := cp.NewDampedSpring(bd1, bd2, cp.Vector{}, cp.Vector{}, 10, 5, 0.3)
	// NewRotaryLimitJoint 限制旋转角度的最大差值
	//con := cp.NewRotaryLimitJoint(bd1, bd2, 0, math.Pi/2)
	// NewRatchetJoint 单向旋转限制同步
	//con := cp.NewRatchetJoint(bd1, bd2, 0, math.Pi/2)
	// NewGearJoint 变速齿轮
	//con := cp.NewGearJoint(bd1, bd2, 0, 2)
	// NewSimpleMotor 马达？
	//con := cp.NewSimpleMotor(bd1, bd2, math.Pi*9)
	//sp.AddConstraint(con)
	// 小车
	bd := sp.AddBody(cp.NewBody(1, cp.MomentForBox(1, 12, 4)))
	sp3 := sp.AddShape(cp.NewBox(bd, 12, 4, 0))
	sp3.SetFriction(0.7)
	// 同一组内不会互相碰撞
	sp3.SetFilter(cp.NewShapeFilter(1, cp.ALL_CATEGORIES, cp.ALL_CATEGORIES))
	sp.AddConstraint(cp.NewPivotJoint2(bd, bd1, cp.Vector{X: -5}, cp.Vector{}))
	sp.AddConstraint(cp.NewPivotJoint2(bd, bd2, cp.Vector{X: 5}, cp.Vector{}))
	return &Joints{Space: sp, Drawer: NewDrawer(), Draper: NewDraper()}
}

func (o *Joints) Update() error {
	o.Space.Step(1 / float64(ebiten.TPS()))
	o.Draper.Update(o.Space)
	return nil
}

func (o *Joints) Draw(screen *ebiten.Image) {
	o.Drawer.DrawSpace(screen, o.Space)
}

func (o *Joints) Layout(w, h int) (int, int) {
	return w, h
}
