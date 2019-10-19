package score

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type Score struct {
	txt                  *text.Text
	success, fail, total int
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
	if s.success >= s.fail {
		s.txt.Color = colornames.Green
	} else {
		s.txt.Color = colornames.Red
	}
	fmt.Fprintf(s.txt, "Score: %d, Fail: %d, Total: %d", s.success, s.fail, s.total)
}

func (s *Score) Draw(t pixel.Target) {
	s.txt.Draw(t, pixel.IM.Scaled(s.txt.Orig, 2))
}

func (s *Score) IncSuccess() int {
	s.success += 1
	s.total += 1
	s.Update()
	return s.success
}

func (s *Score) IncFail() int {
	s.fail += 1
	s.total += 1
	s.Update()
	return s.fail
}
