package main

import (
	"GoRaytracer/src/mathutils"
	"GoRaytracer/src/raytracer"
	"GoRaytracer/src/sdlwrapper"
	"GoRaytracer/src/utils"
	"fmt"
	"sync"
)

type Pixel struct {
	x, y  int
	color utils.Color
}

const frameHeight int = 480
const frameWidth int = 640

var vfb [frameHeight][frameWidth]utils.Color
var scene raytracer.Scene
var camera raytracer.ParallelCamera

var finishedRendering bool = false
var pixels chan Pixel

func setupScene() {
	// Setup camera
	camera = raytracer.NewParallelCamera(mathutils.NewVector(60, 60, -100), 0, 30, 0, 90, float64(frameWidth)/float64(frameHeight))
	scene = raytracer.NewScene()
	scene.SetAmbientLight(utils.NewColor(255, 255, 255))
	light1 := raytracer.NewLight(mathutils.NewVector(35, 180, -100), utils.NewColor(255, 255, 255), 25000)
	scene.AddLight(light1)
	// Add scene nodes

	// Add a plane
	plane := raytracer.NewPlane(mathutils.NewVector(0, 0, 0), 300, 1)
	planeColor := raytracer.NewSimpleColor(utils.NewColor(0, 0, 255))
	planeShader := raytracer.NewLambert(utils.Color{0, 0, 0}, &planeColor)
	scene.AddNode(&plane, &planeShader)

	// Add a sphere
	sphere := raytracer.NewSphere(mathutils.NewVector(0, 70, 0), 20)
	sphereColor := raytracer.NewSimpleColor(utils.NewColor(255, 0, 0))
	//sphereChecker := raytracer.NewChecker(utils.NewColor(0, 255, 0), utils.NewColor(255, 0, 0), 50)
	//sphereShader := raytracer.NewPhong(utils.Color{0, 255, 0}, &sphereColor, 5.3, 20)
	sphereShader := raytracer.NewLambert(utils.Color{0, 255, 0}, &sphereColor)
	scene.AddNode(&sphere, &sphereShader)
}

func raytrace(ray *raytracer.Ray) utils.Color {
	var info raytracer.IntersectionInfo
	closestDistance := 1e99
	var closestInfo raytracer.IntersectionInfo
	var closestNode raytracer.Node
	for _, node := range scene.SceneNodes {
		if (*node.GetGeometry()).Intersect(ray, &info) {
			if info.Distance < closestDistance {
				closestDistance = info.Distance
				closestInfo = info
				closestNode = node
			}
		}
	}

	if closestDistance < 1e99 {
		return (*closestNode.GetShader()).Shade(ray, &closestInfo, &scene)
	}

	return utils.NewColor(255, 255, 255)
}

func render(displayWrapper *sdlwrapper.DisplayWrapper) {
	var wg sync.WaitGroup
	wg.Add(frameWidth)
	for x := 0; x < frameWidth; x++ {
		go func(x int) {
			defer wg.Done()
			for y := 0; y < frameHeight; y++ {
				ray := camera.GetScreenRay(float64(x), float64(y))
				resColor := raytrace(&ray)
				vfb[y][x] = resColor
				pixels <- Pixel{x, y, resColor}
			}
		}(x)
	}
	wg.Wait()
	close(pixels)
	finishedRendering = true
}

func display(displayWrapper *sdlwrapper.DisplayWrapper) {
	for !finishedRendering {
		for i := 0; i < 1024; i++ {
			pixel := <-pixels
			displayWrapper.DrawPixel(pixel.x, pixel.y, pixel.color)
		}
		displayWrapper.Display()
	}

	for pixel := range pixels {
		displayWrapper.DrawPixel(pixel.x, pixel.y, pixel.color)
	}
	displayWrapper.Display()
}

func main() {
	displayWrapper, err := sdlwrapper.NewDisplayWrapper(frameWidth, frameHeight, "GoRaytracer")
	if err != nil {
		return
	}
	defer displayWrapper.Destroy()

	pixels = make(chan Pixel, 1024)

	setupScene()
	go render(displayWrapper)
	display(displayWrapper)
	fmt.Println("We got here!")
	saver := utils.NewPNGSaver(frameWidth, frameHeight, "D:\\Go\\image.png")
	saver.Open()

	for x := 0; x < frameWidth; x++ {
		for y := 0; y < frameHeight; y++ {
			saver.SetPixel(x, y, vfb[y][x])
		}
	}
	saver.Save()
	saver.Close()

	sdlwrapper.WaitExit()
}
