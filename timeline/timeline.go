package timeline

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type TimeLine struct {
	xStart, xEnd, yStart, yEnd, pxDelta, width float64
	imd                                        *imdraw.IMDraw
}

func NewTimeLine() *TimeLine {
	line := imdraw.New(nil)
	line.Color = colornames.Blue
	line.EndShape = imdraw.RoundEndShape
	t := &TimeLine{
		xStart:  480,
		xEnd:    480,
		yStart:  300,
		yEnd:    0,
		pxDelta: 40,
		width:   30,
		imd:     line,
	}

	t.UpdateXY()
	return t
}

func (t *TimeLine) Reset() {
	t.SetXY(480, 480, 300, 0)
	t.UpdateXY()
}

func (t *TimeLine) SetXY(xStart, xEnd, yStart, yEnd float64) {
	t.xStart = xStart
	t.xEnd = xEnd
	t.yStart = yStart
	t.yEnd = yEnd
}

func (tm *TimeLine) Draw(t pixel.Target) {
	tm.imd.Draw(t)
}

func (t *TimeLine) UpdateXY() {
	t.imd.Clear()
	t.imd.Push(pixel.V(t.xStart, t.yStart), pixel.V(t.xEnd, t.yEnd)) //start(x,y) stop(x, y)   x-> , y ^
	t.imd.Line(t.width)
}
func (t *TimeLine) Dec() bool {
	if delta := t.yStart - t.pxDelta; delta < 0 {
		return false
	}
	t.yStart = t.yStart - t.pxDelta
	t.UpdateXY()
	return true
}
