package game

import (
	"image/color"
	"math/rand"

	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/physics"

	"github.com/hajimehoshi/ebiten/v2"
)

var scrh = float64(defs.ScreenHeight)
var scrtop = float64(defs.ScreenTop)

type World struct {
	points  []float64
	img     *ebiten.Image
	ops     *ebiten.DrawImageOptions
	engine  *Engine
	explode bool
	counter int
	active  bool
}

func NewWorld(engine *Engine) *World {
	w := &World{
		engine: engine,
		active: true,
	}

	w.points = make([]float64, defs.WorldWidth+1)
	var y float64 = 50
	var dy float64 = 1

	for i := 0; i <= defs.WorldWidth; i++ {
		w.points[i] = y
		y += dy

		if i > 50 {
			if y < 50 || y > defs.ScreenHeight/4 || rand.Intn(10) == 1 {
				dy = -dy
			}
		} else {
			dy = 1
		}
	}

	y = 50
	dy = 1

	for i := defs.WorldWidth; i > 0; i-- {
		if y >= w.points[i] {
			break
		}

		w.points[i] = y
		dy := physics.RandChoiceF([]float64{0, 1, 1})
		y += dy
	}

	w.img = ebiten.NewImage(2, 2)
	w.ops = &ebiten.DrawImageOptions{}
	w.img.Fill(color.White)
	w.ops.ColorM.Scale(0.5, 0.3, 0, 1)

	return w
}

func (w *World) SetActive(b bool) {
	w.active = b
}

func (w *World) Explode() {
	w.explode = true
}

func (w *World) At(wx float64) float64 {
	if wx < 0 {
		wx = 0
	}

	if wx > defs.WorldWidth {
		wx = defs.WorldWidth
	}

	return w.points[int(wx)]
}

func (w *World) Update() {
	if !w.active {
		return
	}

	if w.explode {
		w.counter++

		if w.counter > defs.WorldExplodeTicks {
			return
		}

		ww := defs.WorldWidth
		i := int(defs.CameraX())

		for x := -2000; x < defs.ScreenWidth+2000; x++ {
			if i < 0 {
				i += ww
			} else if i > ww {
				i -= ww
			}

			w.points[i] += 20*rand.Float64() - 10

			if rand.Intn(50) < 1 {
				w.points[i] -= 10000
			}
			i++
		}
	}
}

func (w *World) Draw(scr *ebiten.Image) {
	if !w.active {
		return
	}

	if w.counter > defs.WorldExplodeTicks {
		return
	}

	ww := defs.WorldWidth
	i := int(defs.CameraX())

	for x := 0; x < defs.ScreenWidth; x++ {
		if i < 0 {
			i += ww
		} else if i > ww {
			i -= ww
		}

		h := w.points[i]
		w.ops.GeoM.Reset()
		xOffset := 0

		if w.explode {
			s := 4 * float64(w.counter) / float64(defs.WorldExplodeTicks)
			w.ops.GeoM.Scale(1+s, 1+s)
			xOffset = rand.Intn(30) - 15
		}

		w.ops.GeoM.Translate(float64(x+xOffset), float64(defs.ScreenHeight-h))
		scr.DrawImage(w.img, w.ops)
		i++
	}

	sw := float64(defs.ScreenWidth)
	cx := defs.CameraX() - float64(ww/2) + sw/2

	rs := sw / 4
	rw := sw / 2

	for j := 0; j < ww; j += 10 {
		ind := j + int(cx)

		if ind < 0 {
			ind += ww
		}

		if ind > ww-1 {
			ind -= ww
		}

		h := w.points[ind]
		sx := rs + rw*(float64(j)/float64(ww))
		w.ops.GeoM.Reset()
		w.ops.GeoM.Scale(0.5, 0.5)
		w.ops.GeoM.Translate(sx, float64(scrtop)-(float64(h*(scrtop/scrh))))

		scr.DrawImage(w.img, w.ops)
	}
}
