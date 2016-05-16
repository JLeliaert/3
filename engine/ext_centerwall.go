package engine

import (
	"fmt"
	"github.com/mumax/3/data"
)

var (
	DWPos = NewGetScalar("ext_dwpos", "m", "Position of the simulation window while following a domain wall", GetShiftPos) // TODO: make more accurate
	DWPMAXPos = NewGetScalar("ext_PMAdwxpos", "m", "Position of the simulation window while following a domain wall", GetPMAPos)
	DWXPos = NewGetScalar("ext_dwxpos", "m", "Position of the simulation window while following a domain wall", GetXPos)
)

func init() {
	DeclFunc("ext_centerWall", CenterWall, "centerWall(c) shifts m after each step to keep m_c close to zero")
}

func GetJonathanPos(c int) float64 {
	mc := M.Comp(c).Average()
	n := Mesh().Size()
	cs := Mesh().CellSize()
	pos := (float64(n[0])*cs[0])/2.*(mc+1)
	pos += GetShiftPos()
return float64(pos)
}

func GetPMAPos() float64 {
	mc := M.Comp(2).Average()
	n := Mesh().Size()
	cs := Mesh().CellSize()
	pos := (float64(n[0])*cs[0])/2.*(mc+1)
	pos += GetShiftPos()
return float64(pos)
}

func GetXPos() float64 {
	c :=0
	M := &M
	mc := sAverageUniverse(M.Buffer().Comp(c))[0]
	n := Mesh().Size()
	cs := Mesh().CellSize()
	pos := (float64(n[0])*cs[0])/2.*(mc+1)
	pos += GetShiftPos()
return float64(pos)
}

func centerWall(c int) {
	component =c
	M := &M
	mc := sAverageUniverse(M.Buffer().Comp(c))[0]
	n := Mesh().Size()
	tolerance := 4 / float64(n[X]) // x*2 * expected <m> change for 1 cell shift

	zero := data.Vector{0, 0, 0}
	if ShiftMagL == zero || ShiftMagR == zero {
		sign := magsign(M.GetCell(0, n[Y]/2, n[Z]/2)[c])
		ShiftMagL[c] = float64(sign)
		ShiftMagR[c] = -float64(sign)
	}

	sign := magsign(ShiftMagL[c])

	//log.Println("mc", mc, "tol", tolerance)

	if mc < -tolerance {
		Shift(sign)
	} else if mc > tolerance {
		Shift(-sign)
	}
}

// This post-step function centers the simulation window on a domain wall
// between up-down (or down-up) domains (like in perpendicular media). E.g.:
// 	PostStep(CenterPMAWall)
func CenterWall(magComp int) {
	PostStep(func() { centerWall(magComp) })
}

func magsign(x float64) int {
	if x > 0.1 {
		return 1
	}
	if x < -0.1 {
		return -1
	}
	panic(fmt.Errorf("center wall: unclear in which direction to shift: magnetization at border=%v. Set ShiftMagL, ShiftMagR", x))
}

// used for speed
var (
	lastShift float64 // shift the last time we queried speed
	lastT     float64 // time the last time we queried speed
	lastV     float64 // speed the last time we queried speed
	DWSpeed   = NewGetScalar("ext_dwspeed", "m/s", "Speed of the simulation window while following a domain wall", getShiftSpeed)
	component  int //the component used in centerwall
)

func getShiftSpeed() float64 {
	if lastShift != GetJonathanPos(component) {
		lastV = (GetJonathanPos(component) - lastShift) / (Time - lastT)
		lastShift = GetJonathanPos(component)
		lastT = Time
	}
	return lastV
}

