package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
)

type CollisionHandler struct {
	Space  *cp.Space
	Drawer *Drawer
}

func NewCollisionHandler() *CollisionHandler {
	sp := cp.NewSpace()
	sp.SetGravity(cp.Vector{Y: -10})
	// 设置地面
	sp.StaticBody.SetPosition(cp.Vector{Y: -15}) // 不推荐改公共的
	tmp := sp.AddShape(cp.NewBox(sp.StaticBody, 40, 5, 0))
	tmp.SetCollisionType(1)
	// 设置另外一个碰撞体
	bd := sp.AddBody(cp.NewBody(1, cp.MomentForBox(1, 20, 5)))
	bd.SetPosition(cp.Vector{X: 0, Y: 15})
	tmp = sp.AddShape(cp.NewBox(bd, 20, 5, 0))
	tmp.SetCollisionType(2)
	// 设置碰撞解析器，只处理 1 与 2 的碰撞，位运算
	handler := sp.NewCollisionHandler(1, 2)
	handler.PreSolveFunc = func(arb *cp.Arbiter, space *cp.Space, userData any) bool {
		fmt.Println("preSolveFunc")
		return true // 返回 false 忽略本次碰撞
	}
	return &CollisionHandler{Space: sp, Drawer: NewDrawer()}
}

func (o *CollisionHandler) Update() error {
	o.Space.Step(1 / float64(ebiten.TPS()))
	return nil
}

func (o *CollisionHandler) Draw(screen *ebiten.Image) {
	o.Drawer.DrawSpace(screen, o.Space)
}

func (o *CollisionHandler) Layout(w, h int) (int, int) {
	return w, h
}
