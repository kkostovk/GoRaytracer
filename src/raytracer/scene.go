package raytracer

import "GoRaytracer/src/utils"

type Scene struct {
	sceneNodes   []Node
	lights       []Light
	ambientLight utils.Color
}

func NewScene() Scene {
	return Scene{make([]Node, 0), make([]Light, 0), utils.Color{0.5, 0.5, 0.5}}
}

func (s *Scene) SetAmbientLight(color utils.Color) {
	s.ambientLight = color
}

func (s *Scene) AddLight(light Light) {
	s.lights = append(s.lights, light)
}

func (s *Scene) AddNode(geometry *Geometry, shader *Shader) {
	node := NewNode(geometry, shader)
	s.sceneNodes = append(s.sceneNodes, node)
}
