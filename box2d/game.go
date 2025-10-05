package main

import (
	"image/color"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	b2 "github.com/oliverbestmann/box2d-go"
)

type Game struct {
	World  *b2.World
	Body   *b2.Body
	Option *ebiten.DrawTrianglesOptions
}

func NewGame() *Game {
	// 创建世界
	world := NewWorld()
	// 创建对象
	NewBox(world, b2.Vec2{X: -16, Y: -9}, b2.Vec2{X: -15, Y: 9})
	NewBox(world, b2.Vec2{X: 15, Y: -9}, b2.Vec2{X: 16, Y: 9})
	NewBox(world, b2.Vec2{X: -16, Y: -9}, b2.Vec2{X: 16, Y: -8})

	def := b2.DefaultBodyDef()
	def.Type1 = b2.DynamicBody // 动态对象
	def.Position = b2.Vec2{X: 0, Y: 6}
	body := world.CreateBody(def)
	body.CreateCircleShape(b2.DefaultShapeDef(), b2.Circle{Center: b2.ZeroVec2, Radius: 2})
	pos := [8]b2.Vec2{{X: 0, Y: 8}, {X: -2, Y: 0}, {X: 2, Y: 0}}
	body.CreatePolygonShape(b2.DefaultShapeDef(), b2.MakePolygon(b2.Hull{Points: pos, Count: 3}, 0))
	temp := body.GetMassData() // 调整重心
	temp.Center.Y = -1
	body.SetMassData(temp)
	//sd := b2.DefaultShapeDef()
	//sd.Material.Restitution = 0.5 // 反弹性
	//body.CreateCircleShape(sd, b2.Circle{Center: b2.Vec2{X: 0, Y: 0}, Radius: 1})
	//sd = b2.DefaultShapeDef()
	//sd.Material.Restitution = 0.5 // 反弹性
	//body.CreatePolygonShape(sd, b2.MakeOffsetBox(0.5, 2, b2.Vec2{X: 0, Y: 1}, b2.IdentityRot))
	return &Game{World: &world, Body: &body, Option: &ebiten.DrawTrianglesOptions{}}
}

func (g *Game) Update() error {
	g.World.Step(1.0/60, 4)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		//g.Body.SetMassData(b2.MassData{}) // 设置质量重心等信息
		//fmt.Println("Mouse button left", g.Body.GetMass()) // 获取质量
		// 添加力的作用 改变加速度
		//g.Body.ApplyForce(b2.Vec2{Y: 100}, g.Body.GetPosition(), 1)
		// 直接叠加到当前速度上 改变速度 角度是传入对应的 cos sin 值，方便内部计算
		//g.Body.ApplyLinearImpulse(b2.Vec2{Y: 5}, g.Body.GetPosition(), 1)
		// 添加扭力
		g.Body.ApplyTorque(200, 1)
		//g.World.OverlapAABB()
		//g.Body.ComputeAABB() // 计算包围盒
	}
	return nil
}

var (
	empty = ebiten.NewImage(1, 1)
)

func init() {
	empty.Set(0, 0, color.White)
}

func (g *Game) Draw(screen *ebiten.Image) {
	//t := g.Body2.GetTransform() // 获取数据自定义绘制
	//g.World.OverlapShape(b2.MakeBox(2, 2), nil, nil)

	g.World.Draw(b2.DebugDraw{
		DrawShapes: true, // 4个点
		DrawJoints: true,
		DrawMass:   true,
		DrawSolidPolygon: func(t b2.Transform, vs []b2.Vec2, r float32, c b2.HexColor) {
			ts := make([]ebiten.Vertex, 0)
			clr := ToColor(c)
			mat := mgl32.Translate2D(t.P.X, t.P.Y).Mul3(mgl32.HomogRotate2D(t.Q.Angle()))
			for _, item := range vs {
				pos := mat.Mul3x1(mgl32.Vec3{item.X, item.Y, 1})
				x, y := ToScreenPos(pos[0], pos[1])
				ts = append(ts, ebiten.Vertex{
					DstX:   x,
					DstY:   y,
					ColorA: 1, ColorR: float32(clr.R) / 255, ColorB: float32(clr.B) / 255, ColorG: float32(clr.G) / 255,
				})
			}
			idxes := make([]uint16, 0)
			for i := 2; i < len(vs); i++ {
				idxes = append(idxes, 0, uint16(i-1), uint16(i))
			}
			screen.DrawTriangles(ts, idxes, empty, g.Option)
		},
		DrawSolidCircle: func(t b2.Transform, r float32, c b2.HexColor) {
			clr := ToColor(c)
			x, y := ToScreenPos(t.P.X, t.P.Y)
			r = ToScreenLen(r)
			vector.DrawFilledCircle(screen, x, y, r, clr, false)
		},
		DrawString: func(p b2.Vec2, s string, c b2.HexColor) {
			x, y := ToScreenPos(p.X, p.Y)
			ebitenutil.DebugPrintAt(screen, s, int(x)-25, int(y))
		},
		DrawSegment: func(p1 b2.Vec2, p2 b2.Vec2, c b2.HexColor) {
			x0, y0 := ToScreenPos(p1.X, p1.Y)
			x1, y1 := ToScreenPos(p2.X, p2.Y)
			clr := ToColor(c)
			vector.StrokeLine(screen, x0, y0, x1, y1, 5, clr, false)
		},
	}) // 自带的调试绘制器
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}
