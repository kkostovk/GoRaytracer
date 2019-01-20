package main

import (
	"GoRaytracer/src/mathutils"
	"GoRaytracer/src/raytracer"
	"GoRaytracer/src/utils"

	"github.com/veandco/go-sdl2/sdl"
)

const frameHeight int = 480
const frameWidth int = 640

var Lights []raytracer.Light
var nodes []*raytracer.Node
var camera raytracer.ParallelCamera

func setupScene() {
	// Setup camera
	camera = raytracer.NewParallelCamera(mathutils.NewVector(40, -30, 50), 0, -30, 0, 90, float64(frameWidth)/float64(frameHeight))
	// Add scene nodes
	plane := raytracer.NewPlane(mathutils.NewVector(0, 0, 0), 500.0, 1)
	checker := raytracer.NewChecker(utils.NewColor(255, 0, 0), utils.NewColor(0, 0, 255))
	node := raytracer.NewNode(&plane, &checker)
	nodes = append(nodes, &node)
}

func raytrace(ray *raytracer.Ray) utils.Color {
	var info raytracer.IntersectionInfo

	var closestInfo *raytracer.IntersectionInfo
	var closestNode *raytracer.Node
	for _, node := range nodes {
		if (*node.GetGeometry()).Intersect(ray, &info) {
			if closestInfo == nil || info.Distance < closestInfo.Distance {
				closestInfo = &info
				closestNode = node
			}
		}
	}

	if closestNode != nil {
		return (*closestNode.GetShader()).Shade(ray, closestInfo)
	}

	return utils.NewColor(0, 0, 0)
}

func render(renderer *sdl.Renderer) {
	for x := 0; x < frameWidth; x++ {
		for y := 0; y < frameHeight; y++ {
			ray := camera.GetScreenRay(float64(x), float64(y))
			resColor := raytrace(&ray)
			r, g, b := resColor.ToRGB()
			renderer.SetDrawColor(r, g, b, 0)
			renderer.DrawPoint(int32(x), int32(y))
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

	setupScene()
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
