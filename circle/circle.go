package circle

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Circle struct {
	x, y, diam float64
	imd        *imdraw.IMDraw
}

func (c *Circle) Draw(t pixel.Target) {
	c.imd.Draw(t)
}
func NewCircle(x, y float64) *Circle {
	imd := imdraw.New(nil)
	imd.Color = colornames.Lightgrey
	c := &Circle{
		x:    x,
		y:    y,
		imd:  imd,
		diam: 70,
	}
	c.SetXY(x, y)
	return c
}

func (c *Circle) SetXY(x, y float64) {
	c.imd.Clear()
	c.imd.Push(pixel.V(x, y))
	c.imd.Circle(c.diam, 0)
}

func (c *Circle) InRange(v pixel.Vec) bool {
	x, y := v.Sub(pixel.V(c.x, c.y)).XY()
	return math.Abs(x) <= c.diam && math.Abs(y) <= c.diam

}
