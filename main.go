package main

import (
	"flag"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/yorikya/brain_cheating/circle"
	"github.com/yorikya/brain_cheating/score"
	"github.com/yorikya/brain_cheating/timeline"
	"golang.org/x/image/colornames"
)

var delay = flag.Bool("delay", false, "Run with click delay")

func main() {
	flag.Parse()
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Cheat The Brain",
		Bounds: pixel.R(0, 0, 1024, 768),
		// VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// time line
	timeline := timeline.NewTimeLine() //width

	//circle
	circle := circle.NewCircle(100, 200)

	//score
	score := score.NewScore()

	last := time.Now()
	for !win.Closed() {
		cam := pixel.IM.Moved(win.Bounds().Center().Sub(pixel.ZV))
		win.SetMatrix(cam)

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			if circle.InRange(cam.Unproject(win.MousePosition())) {
				if *delay {
					time.Sleep(100 * time.Millisecond)
				}

				circle.RandXY()
				timeline.Reset()
				score.IncSuccess()
			}
			// fmt.Printf("%+v in range: %s\n", mouse, circle.InRange(mouse))
		}

		if time.Since(last) > 300*time.Millisecond {
			if !timeline.Dec() {
				circle.RandXY()
				timeline.Reset()
				score.IncFail()
			}
			last = time.Now()
		}

		win.Clear(colornames.Black)

		circle.Draw(win)
		timeline.Draw(win)
		score.Draw(win)

		win.Update()
		time.Sleep(40 * time.Millisecond)
	}
}
