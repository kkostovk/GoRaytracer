// Package sdlwrapper provides a simple wrapper for SDL.
package sdlwrapper

import (
	"GoRaytracer/src/utils"
	"errors"

	"github.com/veandco/go-sdl2/sdl"
)

// Defines a wrapper for sdls window and renderer.
type DisplayWrapper struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

// Create and initialize a new DisplayWrapper with the given width, height and name.
// Return a pointer to the DisplayWrapper and nil error if everything is successful.
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

// Destroy the sdl window and renderer, and call sdl.Quit().
func (d *DisplayWrapper) Destroy() {
	d.renderer.Destroy()
	d.window.Destroy()
	sdl.Quit()
}

// Draw a pixel with the given color at the given coordinates.
func (d *DisplayWrapper) DrawPixel(x, y int, color utils.Color) {
	r, g, b := color.ToRGB()
	d.renderer.SetDrawColor(r, g, b, 255)
	d.renderer.DrawPoint(int32(x), int32(y))
}

// Display the pixels that have been set.
func (d *DisplayWrapper) Display() {
	d.renderer.Present()
}

// Wait for an exit event.
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
