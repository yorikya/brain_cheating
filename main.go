package main

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/yorikya/brain_cheating/circle"
	"github.com/yorikya/brain_cheating/timeline"
	"golang.org/x/image/colornames"
)

func main() {

	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
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

	last := time.Now()
	for !win.Closed() {
		cam := pixel.IM.Moved(win.Bounds().Center().Sub(pixel.ZV))
		win.SetMatrix(cam)

		if win.JustPressed(pixelgl.MouseButtonLeft) {

			mouse := cam.Unproject(win.MousePosition())
			if circle.InRange(mouse) {
				circle.RandXY()
			}
			fmt.Printf("%+v in range: %s\n", mouse, circle.InRange(mouse))
		}

		if time.Since(last) > time.Second {
			// fmt.Println("Pass second")
			timeline.Dec()
			last = time.Now()
		}

		win.Clear(colornames.Aliceblue)
		circle.Draw(win)
		timeline.Draw(win)
		win.Update()
		time.Sleep(50 * time.Millisecond)
	}
}
