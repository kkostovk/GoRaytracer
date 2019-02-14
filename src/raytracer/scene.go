// Package raytracer provides the raytracer logic.
package raytracer

import "GoRaytracer/src/utils"

// Scene defines a holder for the all the scene elements.
type Scene struct {
	SceneNodes   []Node      // Holds all the nodes in the scene.
	lights       []Light     // Holds all the lights in the scene.
	ambientLight utils.Color // Holds the ambient light of the scene.
}

// NewScene creates a new empty scene with a default ambient light.
func NewScene() Scene {
	return Scene{make([]Node, 0), make([]Light, 0), utils.Color{0.5, 0.5, 0.5}}
}

// SetAmbientLight sets the ambient light of the scene to the specified color.
func (s *Scene) SetAmbientLight(color utils.Color) {
	s.ambientLight = color
}

// AddLight adds a light to the scene.
func (s *Scene) AddLight(light Light) {
	s.lights = append(s.lights, light)
}

// AddNode adds a node with the specified geometry and shader to the scene.
func (s *Scene) AddNode(geometry Geometry, shader Shader) {
	s.SceneNodes = append(s.SceneNodes, NewNode(&geometry, &shader))
}
