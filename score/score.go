package score

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type Score struct {
	txt           *text.Text
	success, fail int
}

func NewScore() *Score {
	s := &Score{
		txt: text.New(pixel.V(-480, 360), text.NewAtlas(basicfont.Face7x13, text.ASCII)),
	}
	s.Update()
	return s
}

func (s *Score) Update() {
	s.txt.Clear()
	fmt.Fprintf(s.txt, "Score: %d, Fail: %d", s.success, s.fail)
}

func (s *Score) Draw(t pixel.Target) {
	s.txt.Draw(t, pixel.IM.Scaled(s.txt.Orig, 2))
}

func (s *Score) IncSuccess() {
	s.success += 1
	s.Update()
}

func (s *Score) IncFail() {
	s.fail += 1
	s.Update()
}

