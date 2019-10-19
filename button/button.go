package button

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type Button struct {
	x, y, width float64
	strtxt      string
	imd         *imdraw.IMDraw
	txt         *text.Text
}

func (b *Button) Draw(t pixel.Target) {
	b.imd.Draw(t)
	b.txt.Draw(t, pixel.IM.Scaled(b.txt.Orig, 4))
}
func NewButton(str string, x, y float64) *Button {
	imd := imdraw.New(nil)
	imd.Color = colornames.White
	imd.EndShape = imdraw.RoundEndShape
	b := &Button{
		x:      x,
		y:      y,
		width:  70,
		imd:    imd,
		strtxt: str,
		txt:    text.New(pixel.V(x-15, y-15), text.NewAtlas(basicfont.Face7x13, text.ASCII)),
	}

	b.Update()
	return b
}

func (b *Button) Update() {
	b.imd.Clear()
	b.imd.Push(pixel.V(b.x, b.y), pixel.V(b.x+100, b.y)) //start(x,y) stop(x, y)   x-> , y ^
	b.imd.Line(b.width)

	b.txt.Clear()
	b.txt.Color = colornames.Black
	fmt.Fprintf(b.txt, "%s", b.strtxt)
}

func (b *Button) Click(v pixel.Vec) bool {
	x, y := v.Sub(pixel.V(b.x, b.y)).XY()
	return math.Abs(x) <= b.width && math.Abs(y) <= b.width

}
