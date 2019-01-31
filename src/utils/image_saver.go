package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type ImageSaver interface {
	Open() bool
	SetPixel(int, int, Color)
	Save() bool
	Close() bool
}

type PNGSaver struct {
	width, height int
	fileName      string
	image         *image.RGBA
	file          *os.File
}

func NewPNGSaver(width, height int, fileName string) PNGSaver {
	return PNGSaver{width, height, fileName, image.NewRGBA(image.Rect(0, 0, width, height)), nil}
}

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

func (p *PNGSaver) SetPixel(x, y int, col Color) {
	r, g, b := col.ToRGB()
	p.image.Set(x, y, color.RGBA{r, g, b, 255})
}

func (p *PNGSaver) Save() bool {
	err := png.Encode(p.file, p.image)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func (p *PNGSaver) Close() bool {
	err := p.file.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
