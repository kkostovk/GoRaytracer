package sdlwrapper

import (
	"GoRaytracer/src/utils"
	"errors"

	"github.com/veandco/go-sdl2/sdl"
)

type DisplayWrapper struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

func NewDisplayWrapper(width, height int, name string) (*DisplayWrapper, error) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return nil, errors.New("Failed to initialize SDL.")
	}

	var display DisplayWrapper
	display.window, err = sdl.CreateWindow(name, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(width), int32(height), sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, errors.New("Failed to create SDL window.")
	}

	display.renderer, err = sdl.CreateRenderer(display.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		display.window.Destroy()
		return nil, errors.New("Failed to create SDL renderer.")
	}
	display.renderer.Clear()

	return &display, nil
}

func (d *DisplayWrapper) Destroy() {
	d.renderer.Destroy()
	d.window.Destroy()
	sdl.Quit()
}

func (d *DisplayWrapper) DrawPixel(x, y int, color utils.Color) {
	r, g, b := color.ToRGB()
	d.renderer.SetDrawColor(r, g, b, 255)
	d.renderer.DrawPoint(int32(x), int32(y))
}

func (d *DisplayWrapper) Display() {
	d.renderer.Present()
}

func WaitExit() {
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
