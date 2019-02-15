// Package raytracer provides the raytracer logic.
package raytracer

// Node defines a scene node.
type Node struct {
	geometry *Geometry // A pointer to the geometry of the node.
	shader   *Shader   // A pointer to the shader of the node.
}

// NewNode creates and return a new scene node.
func NewNode(geometry *Geometry, shader *Shader) Node {
	return Node{geometry, shader}
}

// GetGeometry returns the geometry associated with the scene node.
func (n *Node) GetGeometry() *Geometry {
	return n.geometry
}

// GetShader return the shader associated with the scene node.
func (n *Node) GetShader() *Shader {
	return n.shader
}

// SetGeometry sets the geometry for the current scene node.
func (n *Node) SetGeometry(geometry Geometry) {
	n.geometry = &geometry
}

// SetShader sets the shader for the current scene node.
func (n *Node) SetShader(shader Shader) {
	n.shader = &shader
}
