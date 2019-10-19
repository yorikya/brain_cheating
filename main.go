package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/yorikya/brain_cheating/button"
	"github.com/yorikya/brain_cheating/circle"
	"github.com/yorikya/brain_cheating/score"
	"github.com/yorikya/brain_cheating/timeline"
	"github.com/yorikya/brain_cheating/wintext"
	"golang.org/x/image/colornames"
)

func main() {
	flag.Parse()
	pixelgl.Run(run)
}

func EndGame(winner bool, score *score.Score, win *pixelgl.Window) {
	win.Clear(colornames.Black)
	var w *wintext.Wintext
	if winner {
		w = wintext.NewWinnerText()
	} else {
		w = wintext.NewLooseText()
	}

	w.Draw(win)
	score.Draw(win)
	win.Update()
	time.Sleep(5 * time.Second)
}

func StartGame(win *pixelgl.Window) (withdelay bool) {
	startbutton := button.NewButton("Start", 100, 0)
	delaybutton := button.NewButton("Delay", -100, 0)

	for !win.Closed() {
		cam := pixel.IM.Moved(win.Bounds().Center().Sub(pixel.ZV))
		win.SetMatrix(cam)

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			if startbutton.Click(cam.Unproject(win.MousePosition())) {
				withdelay = false
				return
			}

			if delaybutton.Click(cam.Unproject(win.MousePosition())) {
				withdelay = true
				return
			}
		}

		win.Clear(colornames.Black)

		startbutton.Draw(win)
		delaybutton.Draw(win)

		win.Update()
	}
	return
}

func run() {
	rand.Seed(time.Now().UnixNano())
	cfg := pixelgl.WindowConfig{
		Title:  "Cheat The Brain",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

Loop:
	delay := StartGame(win)
	timeline := timeline.NewTimeLine() //width
	circle := circle.NewCircle(100, 200)
	score := score.NewScore()
	numMoves := 20

	last := time.Now()
	for !win.Closed() {
		cam := pixel.IM.Moved(win.Bounds().Center().Sub(pixel.ZV))
		win.SetMatrix(cam)

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			if circle.InRange(cam.Unproject(win.MousePosition())) {
				if delay {
					time.Sleep(100 * time.Millisecond)
				}

				circle.RandXY()
				timeline.Reset()
				if score.IncSuccess() == numMoves {
					EndGame(true, score, win)
					goto Loop
				}
			}
		}

		if time.Since(last) > 300*time.Millisecond {
			if !timeline.Dec() {
				circle.RandXY()
				timeline.Reset()
				if score.IncFail() == numMoves {
					EndGame(false, score, win)
					goto Loop
				}
			}
			last = time.Now()
		}

		win.Clear(colornames.Black)

		circle.Draw(win)
		timeline.Draw(win)
		score.Draw(win)

		win.Update()
	}
}
