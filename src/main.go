package main

import (
	"GoRaytracer/src/raytracer"
	"GoRaytracer/src/sdlwrapper"
	"GoRaytracer/src/utils"
	"fmt"
	"os"
)

const (
	help = `
	-s/-sceneFile 	"filePath" 	: Scene file path.\n
	-o/-outputFile 	"filePath" 	: Output file path.\n
	-d/-display 	T/F 		: Display the rendering. True by default.\n
	`
	cannotParseArgument = `Cannot parse argument :`
)

var displayRendering = true

func display(displayWrapper *sdlwrapper.DisplayWrapper, renderManager *raytracer.RenderManager, pixels chan raytracer.Pixel) {
	for renderManager.RenderState() == raytracer.RenderingInProgress {
		for i := 0; i < 1024; i++ {
			pixel := <-pixels
			displayWrapper.DrawPixel(pixel.X, pixel.Y, pixel.Color)
		}
		if displayRendering {
			displayWrapper.Display()
		}
	}

	for pixel := range pixels {
		displayWrapper.DrawPixel(pixel.X, pixel.Y, pixel.Color)
	}
	if displayRendering {
		displayWrapper.Display()
	}
}

func saveResult(renderManager *raytracer.RenderManager, filename string) {
	vfb := renderManager.GetVFB()
	saver := utils.NewPNGSaver(renderManager.GetFrameWidth(), renderManager.GetFrameHeight(), filename)
	saver.Open()

	for x := 0; x < renderManager.GetFrameWidth(); x++ {
		for y := 0; y < renderManager.GetFrameHeight(); y++ {
			saver.SetPixel(x, y, vfb[y][x])
		}
	}
	saver.Save()
	saver.Close()
}

func main() {
	sceneFile := ""
	outputFile := ""
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		switch {
		case arg == "-s" || arg == "-sceneFile":
			i++
			sceneFile = os.Args[i]

		case arg == "-o" || arg == "-ouputFile":
			i++
			outputFile = os.Args[i]

		case arg == "-d" || arg == "-display":
			i++
			value := os.Args[i]
			if value == "F" {
				displayRendering = false
			}

		case arg == "-h" || arg == "-help":
			fmt.Println(help)
		default:
			nextArg := i + 1
			value := os.Args[nextArg]
			if value[0] != '-' {
				fmt.Println(cannotParseArgument, arg, value)
				i++
			} else {
				fmt.Println(cannotParseArgument, arg)
			}
			fmt.Println(help)
		}
	}

	renderManager := raytracer.NewRenderManager()
	renderManager.Setup(sceneFile)

	displayWrapper, err := sdlwrapper.NewDisplayWrapper(renderManager.GetFrameWidth(), renderManager.GetFrameHeight(), "GoRaytracer")
	if err != nil {
		return
	}
	defer displayWrapper.Destroy()

	pixels := renderManager.Render()
	display(displayWrapper, &renderManager, pixels)
	if outputFile != "" {
		saveResult(&renderManager, outputFile)
	}

	sdlwrapper.WaitExit()
}
