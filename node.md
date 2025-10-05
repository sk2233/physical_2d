### 学习资料
> 参考文档：https://chipmunk-physics.net/release/ChipmunkLatest-Docs/
> 
> 学习case：https://github.com/jakecoffman/cp-examples
### 物体分类
静态物：从不或极少运动<br>
运动体：只受代码影响运动，不受重力影响且拥有无限质量<br>
动态体：默认的物理对象受各种物理量影响
### 力与冲量
力是瞬时的，冲量是力在一段时间的积累，因为计算机是离散的需要用冲量来模拟连续的力
### 分组
```go
sp3.SetFilter(cp.NewShapeFilter(1, cp.ALL_CATEGORIES, cp.ALL_CATEGORIES))
```
同一组内的对象不会互相碰撞