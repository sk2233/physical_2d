package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jakecoffman/cp/v2"
)

type BBQuery struct {
	Space  *cp.Space
	Drawer *Drawer
}

func NewBBQuery() *BBQuery {
	// 初始化空间
	sp := cp.NewSpace()
	for i := 0; i < 50; i++ {
		bd := sp.AddBody(cp.NewBody(1, cp.MomentForBox(1, 2, 2)))
		bd.SetPosition(cp.Vector{X: (rand.Float64() - 0.5) * SpaceW, Y: (rand.Float64() - 0.5) * SpaceH})
		bd.SetAngle(rand.Float64() * math.Pi * 2)
		sp.AddShape(cp.NewBox(bd, 2, 2, 0))
	}
	return &BBQuery{Space: sp, Drawer: NewDrawer()}
}

func (o *BBQuery) Update() error {
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

func (o *BBQuery) Draw(screen *ebiten.Image) {
	o.Drawer.DrawSpace(screen, o.Space)

	x, y := ebiten.CursorPosition()
	tx, ty := ToSpacePos(float32(x), float32(y))
	bb := cp.NewBBForCircle(cp.Vector{}, cp.Vector{X: float64(tx), Y: float64(ty)}.Length())
	DrawBB(screen, bb, color.White)
	o.Space.BBQuery(bb, cp.SHAPE_FILTER_ALL, func(shape *cp.Shape, data any) {
		DrawBB(screen, shape.BB(), color.White)
	}, nil)
}

func (o *BBQuery) Layout(w, h int) (int, int) {
	return w, h
}
