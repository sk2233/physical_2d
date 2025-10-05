package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jakecoffman/cp/v2"
)

type Draper struct {
	Mouse *cp.Body
	Joint *cp.Constraint
}

func (d *Draper) Update(space *cp.Space) {
	tx, ty := ebiten.CursorPosition()
	x, y := ToSpacePos(float32(tx), float32(ty))
	mouse := cp.Vector{X: float64(x), Y: float64(y)}
	d.Mouse.SetPosition(mouse)

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		info := space.PointQueryNearest(mouse, 1, cp.SHAPE_FILTER_ALL)
		if info.Shape == nil {
			return
		}
		body := info.Shape.Body()
		if body.GetType() != cp.BODY_DYNAMIC {
			return
		}
		if info.Distance > 0 { // 鼠标在外面 info.Point 距离形状最近的点
			// NewPivotJoint 使用的是世界坐标
			// NewPivotJoint2 使用的是双方的局部坐标
			d.Joint = cp.NewPivotJoint2(d.Mouse, body, cp.Vector{}, body.WorldToLocal(info.Point))
		} else { // 鼠标在内部
			d.Joint = cp.NewPivotJoint2(d.Mouse, body, cp.Vector{}, body.WorldToLocal(d.Mouse.Position()))
		}
		space.AddConstraint(d.Joint)
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if d.Joint == nil {
			return
		}
		space.RemoveConstraint(d.Joint)
		d.Joint = nil
	}
}

func NewDraper() *Draper {
	return &Draper{Mouse: cp.NewKinematicBody()}
}
