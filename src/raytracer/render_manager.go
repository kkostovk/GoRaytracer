// Package raytracer provides the raytracer logic.
package raytracer

import (
	"GoRaytracer/src/utils"
	"fmt"
	"sync"
)

// Renderer state
const (
	RenderingNotStarted = iota
	RenderingInProgress
	FinishedRendering
)

// Pixel defines a screen pixel.
type Pixel struct {
	X, Y  int
	Color utils.Color
}

// RenderManager hhh
type RenderManager struct {
	frameWidth, frameHeight int             // Frame width and height.
	camera                  *ParallelCamera // Pointer to the camera.
	scene                   *Scene          // Pointer to the scene.
	vfb                     [][]utils.Color // Display buffer.
	renderState             int             // The state of the renderer
}

// NewRenderManager creates and returns an empty RenderManager.
func NewRenderManager() RenderManager {
	return RenderManager{0, 0, nil, nil, nil, RenderingNotStarted}
}

// Setup sets up the current RenderManager from a scene file.
func (r *RenderManager) Setup(fileName string) {
	r.setupScene(fileName)
}

// Render prepares and starts the rendering.
// Returns a pixel channel.
func (r *RenderManager) Render() chan Pixel {
	pixels := make(chan Pixel, 1024)
	r.renderState = RenderingInProgress
	go r.render(pixels)
	return pixels
}

// RenderState returns the state of the renderer.
func (r *RenderManager) RenderState() int {
	return r.renderState
}

// GetVFB returns the display buffer.
func (r *RenderManager) GetVFB() [][]utils.Color {
	return r.vfb
}

// GetFrameHeight returns the frame height.
func (r *RenderManager) GetFrameHeight() int {
	return r.frameHeight
}

// GetFrameWidth returns the frame width.
func (r *RenderManager) GetFrameWidth() int {
	return r.frameWidth
}

func (r *RenderManager) render(pixels chan Pixel) {
	var wg sync.WaitGroup
	wg.Add(r.frameWidth)
	for x := 0; x < r.frameWidth; x++ {
		go func(x int) {
			defer wg.Done()
			for y := 0; y < r.frameHeight; y++ {
				ray := r.camera.GetScreenRay(float64(x), float64(y))
				resColor := r.raytrace(&ray)
				r.vfb[y][x] = resColor
				pixels <- Pixel{x, y, resColor}
			}
		}(x)
	}
	wg.Wait()
	close(pixels)
	r.renderState = FinishedRendering
}

func (r *RenderManager) raytrace(ray *Ray) utils.Color {
	var info IntersectionInfo
	closestDistance := 1e99
	var closestInfo IntersectionInfo
	var closestNode Node
	for _, node := range r.scene.SceneNodes {
		if (*node.GetGeometry()).Intersect(ray, &info) {
			if info.Distance < closestDistance {
				closestDistance = info.Distance
				closestInfo = info
				closestNode = node
			}
		}
	}

	if closestDistance < 1e99 {
		return (*closestNode.GetShader()).Shade(ray, &closestInfo, r.scene)
	}

	return utils.NewColor(255, 255, 255)
}

func (r *RenderManager) setupScene(fileName string) {
	sceneReader, err := NewSceneReader(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read frame width and height
	frameWidth, frameHeight, err := sceneReader.GetFrameSettings()
	if err != nil {
		fmt.Println(err)
		return
	}

	r.frameWidth = frameWidth
	r.frameHeight = frameHeight

	r.frameWidth = 640
	r.frameHeight = 480
	r.vfb = make([][]utils.Color, r.frameHeight)
	for row := range r.vfb {
		r.vfb[row] = make([]utils.Color, r.frameWidth)
	}

	scene := NewScene()
	r.scene = &scene

	// Read the camera
	camera, err := sceneReader.GetCamera()
	if err != nil {
		fmt.Println(err)
		return
	}
	r.camera = &(camera)

	// Read ambient light
	r.scene.ambientLight, err = sceneReader.GetAmbientLight()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read lights
	r.scene.lights, err = sceneReader.GetLights()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read scene nodes
	r.scene.SceneNodes, err = sceneReader.GetSceneNodes()
	if err != nil {
		fmt.Println(err)
		return
	}
}
