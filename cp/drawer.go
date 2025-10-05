package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/jakecoffman/cp/v2"
)

type Drawer struct {
	Screen *ebiten.Image
	Image  *ebiten.Image
	Option *ebiten.DrawTrianglesOptions
}

func NewDrawer() *Drawer {
	img := ebiten.NewImage(1, 1)
	img.Set(0, 0, color.White)
	return &Drawer{Image: img, Option: &ebiten.DrawTrianglesOptions{}}
}

func (d *Drawer) DrawSpace(screen *ebiten.Image, space *cp.Space) {
	d.Screen = screen
	cp.DrawSpace(space, d)
}

func (d *Drawer) DrawCircle(pos cp.Vector, angle, radius float64, outline, fill cp.FColor, data any) {
	x, y := ToScreenPos(float32(pos.X), float32(pos.Y))
	r := ToScreenLen(float32(radius))
	vector.DrawFilledCircle(d.Screen, x, y, r, ToColor(fill), false)
	vector.StrokeCircle(d.Screen, x, y, r, 1, ToColor(outline), false)
	ex, ey := x+float32(math.Cos(-angle))*r, y+float32(math.Sin(-angle))*r
	vector.StrokeLine(d.Screen, x, y, ex, ey, 1, ToColor(outline), false)
}

func (d *Drawer) DrawSegment(a, b cp.Vector, fill cp.FColor, data any) {
	ax, ay := ToScreenPos(float32(a.X), float32(a.Y))
	bx, by := ToScreenPos(float32(b.X), float32(b.Y))
	vector.StrokeLine(d.Screen, ax, ay, bx, by, 1, ToColor(fill), false)
}

func (d *Drawer) DrawFatSegment(a, b cp.Vector, radius float64, outline, fill cp.FColor, data any) {
	ax, ay := ToScreenPos(float32(a.X), float32(a.Y))
	bx, by := ToScreenPos(float32(b.X), float32(b.Y))
	vector.StrokeLine(d.Screen, ax, ay, bx, by, 1, ToColor(fill), false)
}

func (d *Drawer) DrawPolygon(count int, verts []cp.Vector, radius float64, outline, fill cp.FColor, data any) {
	// 绘制内部
	vs := make([]ebiten.Vertex, 0)
	idx := make([]uint16, 0)
	for _, item := range verts {
		x, y := ToScreenPos(float32(item.X), float32(item.Y))
		vs = append(vs, ebiten.Vertex{
			DstX:   x,
			DstY:   y,
			ColorR: fill.R,
			ColorG: fill.G,
			ColorB: fill.B,
			ColorA: fill.A,
		})
	}
	for i := 2; i < count; i++ {
		idx = append(idx, 0, uint16(i-1), uint16(i))
	}
	d.Screen.DrawTriangles(vs, idx, d.Image, d.Option)
	// 绘制边框
	for i := 0; i < count; i++ {
		p0, p1 := verts[i], verts[(i+1)%count]
		x0, y0 := ToScreenPos(float32(p0.X), float32(p0.Y))
		x1, y1 := ToScreenPos(float32(p1.X), float32(p1.Y))
		vector.StrokeLine(d.Screen, x0, y0, x1, y1, 1, ToColor(outline), false)
	}
}

func (d *Drawer) DrawDot(size float64, pos cp.Vector, fill cp.FColor, data any) {
	x, y := ToScreenPos(float32(pos.X), float32(pos.Y))
	//l := ToScreenLen(float32(size)) // 防止标识点太大
	vector.DrawFilledCircle(d.Screen, x, y, float32(size), ToColor(fill), false)
}

func (d *Drawer) Flags() uint {
	return cp.DRAW_SHAPES | cp.DRAW_CONSTRAINTS | cp.DRAW_COLLISION_POINTS
}

func (d *Drawer) OutlineColor() cp.FColor {
	return cp.FColor{R: 200.0 / 255.0, G: 210.0 / 255.0, B: 230.0 / 255.0, A: 1}
}

func (d *Drawer) ShapeColor(shape *cp.Shape, data any) cp.FColor {
	val := shape.HashId()
	// scramble the bits up using Robert Jenkins' 32 bit integer hash function
	val = (val + 0x7ed55d16) + (val << 12)
	val = (val ^ 0xc761c23c) ^ (val >> 19)
	val = (val + 0x165667b1) + (val << 5)
	val = (val + 0xd3a2646c) ^ (val << 9)
	val = (val + 0xfd7046c5) + (val << 3)
	val = (val ^ 0xb55a4f09) ^ (val >> 16)
	r := float32((val >> 0) & 0xFF)
	g := float32((val >> 8) & 0xFF)
	b := float32((val >> 16) & 0xFF)
	return cp.FColor{R: r / 255, G: g / 255, B: b / 255, A: 1}
}

func (d *Drawer) ConstraintColor() cp.FColor {
	return cp.FColor{G: 0.75, A: 1}
}

func (d *Drawer) CollisionPointColor() cp.FColor {
	return cp.FColor{R: 1, A: 1}
}

func (d *Drawer) Data() any {
	return nil
}
