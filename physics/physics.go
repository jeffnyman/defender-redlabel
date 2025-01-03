package physics

import (
	"math"
	"math/rand"

	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/defs"
)

func ScreenX(x float64) float64 {
	ww := float64(defs.WorldWidth)
	sw := float64(defs.ScreenWidth)
	cx := defs.CameraX()
	over := sw - (ww - cx)

	if over > 0 && x < over {
		x += ww
	}

	sx := x - cx

	return sx
}

func OffScreen(x, y float64) bool {

	return (x < -100 || x > defs.ScreenWidth+100 || y < 0 || y > defs.ScreenHeight)
}

func RandChoiceF(lst []float64) float64 {
	return lst[rand.Intn(len(lst))]
}

func RandChoiceI(lst []int) int {
	return lst[rand.Intn(len(lst))]
}

func RandChoiceS(lst []string) string {
	return lst[rand.Intn(len(lst))]
}

func ComputeBullet(firepos, playpos *components.Pos, time float64) (float64, float64) {
	tt := defs.MaxTPS * time
	projected_x := playpos.X + (playpos.DX * tt)
	projected_y := playpos.Y
	dx := (projected_x - firepos.X) / tt
	dy := (projected_y - firepos.Y) / tt

	return dx, dy
}

func Collide(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {
	left := math.Max(x1, x2)
	top := math.Max(y1, y2)
	right := math.Min(x1+w1, x2+w2)
	bottom := math.Min(y1+h1, y2+h2)

	return left < right && top < bottom
}

func Clamp(v1, v2, v3 float64) float64 {
	if v1 <= v2 {
		return v2
	}

	if v1 >= v3 {
		return v3
	}

	return v1
}
