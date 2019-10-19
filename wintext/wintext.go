package wintext

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type Wintext struct {
	txt *text.Text
}

func NewWintext() *Wintext {
	w := &Wintext{
		txt: text.New(pixel.V(-300, 0), text.NewAtlas(basicfont.Face7x13, text.ASCII)),
	}
	return w
}

func NewWinnerText() *Wintext {
	w := NewWintext()
	w.Update("You are WIN!!!")
	return w
}

func NewLooseText() *Wintext {
	w := NewWintext()
	w.Update("You are LOOSE :-(")
	return w
}

func (w *Wintext) Update(str string) {
	w.txt.Clear()
	w.txt.Color = colornames.Burlywood
	fmt.Fprintf(w.txt, "%s", str)
}

func (w *Wintext) Draw(t pixel.Target) {
	w.txt.Draw(t, pixel.IM.Scaled(w.txt.Orig, 6))
}
