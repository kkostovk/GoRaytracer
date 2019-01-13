package main

import "github.com/veandco/go-sdl2/sdl"

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("GoRaytracer", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 256, 256, sdl.WINDOW_SHOWN)
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

	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			renderer.SetDrawColor(uint8(i), uint8(j), 0, 0)
			renderer.DrawPoint(int32(i), int32(j))
		}
	}
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
