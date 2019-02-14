// Package utils provides some simple utilities for the raytracer.
package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// Defines a image saver interface.
type ImageSaver interface {
	Open() bool
	SetPixel(int, int, Color)
	Save() bool
	Close() bool
}

// Defines a .png saver that supports the ImageSaver interface.
type PNGSaver struct {
	width, height int
	fileName      string
	image         *image.RGBA
	file          *os.File
}

// Create a new PNGSaver with the given width, height and name, and return it.
func NewPNGSaver(width, height int, fileName string) PNGSaver {
	return PNGSaver{width, height, fileName, image.NewRGBA(image.Rect(0, 0, width, height)), nil}
}

// Open or create(if necessary) the file.
func (p *PNGSaver) Open() bool {
	file, err := os.Create(p.fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return false
	}

	p.file = file
	return true
}

// Set a pixel with the given color at the given coordinates to be saved.
func (p *PNGSaver) SetPixel(x, y int, col Color) {
	r, g, b := col.ToRGB()
	p.image.Set(x, y, color.RGBA{r, g, b, 255})
}

// Save the image in the open file.
func (p *PNGSaver) Save() bool {
	err := png.Encode(p.file, p.image)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

// Close the open file.
func (p *PNGSaver) Close() bool {
	err := p.file.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
