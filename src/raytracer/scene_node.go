package raytracer

type Node struct {
	geometry *Geometry
	shader   *Shader
}

func NewNode(geometry *Geometry, shader *Shader) Node {
	return Node{geometry, shader}
}

func (n *Node) GetGeometry() *Geometry {
	return n.geometry
}

func (n *Node) GetShader() *Shader {
	return n.shader
}
