package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const frameHeight int = 480
const frameWidth int = 640

func render(r *sdl.Renderer) {
	for x := 0; x < frameWidth; x++ {
		for y := 0; y < frameHeight; y++ {
			r.SetDrawColor(uint8(float64(x)/float64(frameWidth)*255.0), uint8(float64(y)/float64(frameHeight)*255), 0, 0)
			r.DrawPoint(int32(x), int32(y))
		}
	}
}

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("GoRaytracer", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(frameWidth), int32(frameHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		return
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		return
	}
	surface.FillRect(nil, 0)

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return
	}
	defer renderer.Destroy()

	render(renderer)
	renderer.Present()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}
	}
}
