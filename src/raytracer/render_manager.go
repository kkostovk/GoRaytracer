// Package raytracer provides the raytracer logic.
package raytracer

import (
	"GoRaytracer/src/mathutils"
	"GoRaytracer/src/utils"
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
func (r *RenderManager) Setup(_fileName string) {
	r.dummySceneSetup()
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

func (r *RenderManager) dummySceneSetup() {
	r.frameWidth = 640
	r.frameHeight = 480
	r.vfb = make([][]utils.Color, r.frameHeight)
	for row := range r.vfb {
		r.vfb[row] = make([]utils.Color, r.frameWidth)
	}
	// Setup camera
	camera := NewParallelCamera(mathutils.NewVector(60, 60, -100), 0, 30, 0, 90, float64(r.frameWidth)/float64(r.frameHeight))
	r.camera = &camera
	scene := NewScene()
	r.scene = &scene
	scene.SetAmbientLight(utils.NewColor(255, 255, 255))
	light1 := NewLight(mathutils.NewVector(35, 180, -100), utils.NewColor(255, 255, 255), 25000)
	scene.AddLight(light1)
	// Add scene nodes

	// Add a plane
	plane := NewPlane(mathutils.NewVector(0, 0, 0), 300, 1)
	planeColor := NewSimpleColor(utils.NewColor(0, 0, 255))
	planeShader := NewLambert(utils.Color{0, 0, 0}, &planeColor)
	scene.AddNode(&plane, &planeShader)

	// Add a sphere
	sphere := NewSphere(mathutils.NewVector(0, 70, 0), 20)
	sphereColor := NewSimpleColor(utils.NewColor(255, 255, 0))
	//sphereChecker := raytracer.NewChecker(utils.NewColor(0, 255, 0), utils.NewColor(255, 0, 0), 50)
	//sphereShader := raytracer.NewPhong(utils.Color{0, 255, 0}, &sphereColor, 5.3, 20)
	sphereShader := NewLambert(utils.Color{0, 255, 0}, &sphereColor)
	scene.AddNode(&sphere, &sphereShader)
}
