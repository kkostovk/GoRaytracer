// Package raytracer provides the raytracer logic.
package raytracer

import (
	"GoRaytracer/src/mathutils"
	"GoRaytracer/src/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// SceneReader provides a way to parse a scene file.
type SceneReader struct {
	fileContent []string // Holds all the words of the scene.
	position    int      // Holds the current position.
}

// NewSceneReader creates and returns a new SceneReader
// If something goes wrong returns nil and error.
func NewSceneReader(filePath string) (*SceneReader, error) {
	content, err := scanWords(filePath)
	if err != nil {
		return nil, err
	}
	return &SceneReader{content, 0}, nil
}

//GetFrameSettings parses and returns the frame width and height.
func (s *SceneReader) GetFrameSettings() (width int, height int, err error) {
	err = check(s.fileContent[s.position], "FrameSettings")
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "{")
	if err != nil {
		return
	}

	for i := 0; i < 2; i++ {
		s.position++
		name := s.fileContent[s.position]
		s.position++
		var value int
		value, err = strconv.Atoi(s.fileContent[s.position])
		if err != nil {
			return
		}
		switch {
		case name == "frameWidth":
			width = value

		case name == "frameHeight":
			height = value
		}
	}
	s.position++
	err = check(s.fileContent[s.position], "}")
	if err != nil {
		return
	}

	s.position++
	return
}

// GetCamera parses and returns the camera from the scene file.
func (s *SceneReader) GetCamera() (camera ParallelCamera, err error) {
	err = check(s.fileContent[s.position], "Camera")
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "{")
	if err != nil {
		return
	}

	var (
		position                           mathutils.Vector
		yaw, pitch, roll, fov, aspectRatio float64
	)

	for i := 0; i < 6; i++ {
		s.position++
		name := s.fileContent[s.position]
		s.position++
		var value float64
		value, err = strconv.ParseFloat(s.fileContent[s.position], 64)
		if err != nil {
			return
		}
		switch {
		case name == "position":
			position, err = s.readVector()
			if err != nil {
				return
			}
		case name == "yaw":
			yaw = value
		case name == "pitch":
			pitch = value
		case name == "roll":
			roll = value
		case name == "fov":
			fov = value
		case name == "aspectRatio":
			aspectRatio = value
		}
	}

	s.position++
	err = check(s.fileContent[s.position], "}")
	if err != nil {
		return
	}
	s.position++

	parallelCamera := NewParallelCamera(position, yaw, pitch, roll, fov, aspectRatio)
	camera = parallelCamera
	return
}

// GetAmbientLight parses and returns the ambient light from the scene file.
func (s *SceneReader) GetAmbientLight() (ambientLight utils.Color, err error) {
	err = check(s.fileContent[s.position], "AmbientLight")
	if err != nil {
		return
	}

	s.position++
	ambientLight, err = s.readColor()
	s.position++
	return
}

// GetLights parses and returns all the lights from the scene file.
func (s *SceneReader) GetLights() (lights []Light, err error) {
	for {
		if s.fileContent[s.position] != "Light" {
			break
		}

		s.position++
		err = check(s.fileContent[s.position], "{")
		if err != nil {
			break
		}

		var light Light

		for i := 0; i < 3; i++ {
			s.position++
			name := s.fileContent[s.position]
			s.position++
			var value float64
			value, err = strconv.ParseFloat(s.fileContent[s.position], 64)
			if err != nil {
				return
			}
			switch {
			case name == "position":
				light.position, err = s.readVector()
				if err != nil {
					return
				}
			case name == "color":
				light.color, err = s.readColor()
				if err != nil {
					return
				}
			case name == "power":
				light.power = value
			}
		}

		s.position++
		err = check(s.fileContent[s.position], "}")
		if err != nil {
			break
		}

		lights = append(lights, light)

		s.position++
	}

	return
}

// GetSceneNodes parses and returns all the scene nodes from the scene file.
func (s *SceneReader) GetSceneNodes() (nodes []Node, err error) {
	for {
		if s.fileContent[s.position] != "Node" {
			break
		}

		s.position++
		err = check(s.fileContent[s.position], "{")
		if err != nil {
			break
		}
		var node Node

		// Read the geometry
		s.position++
		err = check(s.fileContent[s.position], "geometry")
		if err != nil {
			break
		}

		s.position++
		name := s.fileContent[s.position]
		switch {
		case name == "Sphere":
			var sphere Sphere
			sphere, err = s.readSphere()
			if err != nil {
				return
			}
			node.SetGeometry(&sphere)
		case name == "Plane":
			var plane Plane
			plane, err = s.readPlane()
			if err != nil {
				return
			}
			node.SetGeometry(&plane)

		case name == "Cube":
			var cube Cube
			cube, err = s.readCube()
			if err != nil {
				return
			}
			node.SetGeometry(&cube)
		}

		// Read the shader
		err = check(s.fileContent[s.position], "shader")
		if err != nil {
			break
		}

		s.position++
		name = s.fileContent[s.position]
		switch {
		case name == "Lambert":
			var lambert Lambert
			lambert, err = s.readLambert()
			if err != nil {
				return
			}
			node.SetShader(&lambert)

		case name == "Phong":
			var phong Phong
			phong, err = s.readPhong()
			if err != nil {
				return
			}
			node.SetShader(&phong)
		}

		nodes = append(nodes, node)
		err = check(s.fileContent[s.position], "}")
		if err != nil {
			break
		}
		s.position++
	}

	return
}

func scanWords(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil
}

func check(found, expected string) error {
	if found != expected {
		return fmt.Errorf("Incorrect format")
	}

	return nil
}

func (s *SceneReader) readVector() (vector mathutils.Vector, err error) {

	vector.X, err = strconv.ParseFloat(s.fileContent[s.position], 64)
	if err != nil {
		return
	}

	s.position++
	vector.Y, err = strconv.ParseFloat(s.fileContent[s.position], 64)
	if err != nil {
		return
	}

	s.position++
	vector.Z, err = strconv.ParseFloat(s.fileContent[s.position], 64)
	if err != nil {
		return
	}

	return
}

func (s *SceneReader) readColor() (color utils.Color, err error) {
	var r, g, b int

	r, err = strconv.Atoi(s.fileContent[s.position])
	if err != nil {
		return
	}

	s.position++
	g, err = strconv.Atoi(s.fileContent[s.position])
	if err != nil {
		return
	}

	s.position++
	b, err = strconv.Atoi(s.fileContent[s.position])
	if err != nil {
		return
	}
	color = utils.NewColor(uint8(r), uint8(g), uint8(b))
	return
}

func (s *SceneReader) readSphere() (sphere Sphere, err error) {
	s.position++
	err = check(s.fileContent[s.position], "{")
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "center")
	if err != nil {
		return
	}
	s.position++
	sphere.center, err = s.readVector()
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "radius")
	if err != nil {
		return
	}

	s.position++
	sphere.radius, err = strconv.ParseFloat(s.fileContent[s.position], 64)
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "}")
	if err != nil {
		return
	}

	s.position++
	return
}

func (s *SceneReader) readPlane() (plane Plane, err error) {
	s.position++
	err = check(s.fileContent[s.position], "{")
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "center")
	if err != nil {
		return
	}
	s.position++
	plane.center, err = s.readVector()
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "limit")
	if err != nil {
		return
	}

	s.position++
	plane.limit, err = strconv.ParseFloat(s.fileContent[s.position], 64)
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "orientation")
	if err != nil {
		return
	}

	s.position++
	value := s.fileContent[s.position]
	switch {
	case value == "XY":
		plane.orientation = 0
	case value == "XZ":
		plane.orientation = 1
	case value == "YZ":
		plane.orientation = 2
	}

	s.position++
	err = check(s.fileContent[s.position], "}")
	if err != nil {
		return
	}

	s.position++
	return
}

func (s *SceneReader) readCube() (cube Cube, err error) {
	s.position++
	err = check(s.fileContent[s.position], "{")
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "center")
	if err != nil {
		return
	}
	s.position++
	cube.center, err = s.readVector()
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "edge")
	if err != nil {
		return
	}

	s.position++
	cube.edge, err = strconv.ParseFloat(s.fileContent[s.position], 64)
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "}")
	if err != nil {
		return
	}

	s.position++
	return
}

func (s *SceneReader) readLambert() (lambert Lambert, err error) {
	s.position++
	err = check(s.fileContent[s.position], "{")
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "color")
	if err != nil {
		return
	}

	s.position++
	lambert.color, err = s.readColor()
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "texture")
	if err != nil {
		return
	}

	s.position++
	name := s.fileContent[s.position]
	switch {
	case name == "SimpleColor":
		var simpleColor SimpleColor
		simpleColor, err = s.readSimpleColor()
		if err != nil {
			return
		}
		lambert.SetTexture(&simpleColor)
	case name == "Checker":
		var checker Checker
		checker, err = s.readChecker()
		if err != nil {
			return
		}
		lambert.SetTexture(&checker)
	}

	err = check(s.fileContent[s.position], "}")
	if err != nil {
		return
	}
	s.position++

	return
}

func (s *SceneReader) readPhong() (phong Phong, err error) {

	s.position++
	err = check(s.fileContent[s.position], "{")
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "color")
	if err != nil {
		return
	}

	s.position++
	phong.color, err = s.readColor()
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "texture")
	if err != nil {
		return
	}

	s.position++
	name := s.fileContent[s.position]
	switch {
	case name == "SimpleColor":
		var simpleColor SimpleColor
		simpleColor, err = s.readSimpleColor()
		if err != nil {
			return
		}
		phong.SetTexture(&simpleColor)
	case name == "Checker":
		var checker Checker
		checker, err = s.readChecker()
		if err != nil {
			return
		}
		phong.SetTexture(&checker)
	}

	s.position++
	err = check(s.fileContent[s.position], "specularMultiplier")
	if err != nil {
		return
	}

	s.position++
	phong.specularMultiplier, err = strconv.ParseFloat(s.fileContent[s.position], 64)
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "specularExponent")
	if err != nil {
		return
	}

	s.position++
	phong.specularExponent, err = strconv.ParseFloat(s.fileContent[s.position], 64)
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "}")
	if err != nil {
		return
	}
	s.position++

	return
}

func (s *SceneReader) readSimpleColor() (simpleColor SimpleColor, err error) {
	s.position++
	err = check(s.fileContent[s.position], "{")
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "color")
	if err != nil {
		return
	}

	s.position++
	simpleColor.color, err = s.readColor()
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "}")
	if err != nil {
		return
	}
	s.position++
	return
}

func (s *SceneReader) readChecker() (checker Checker, err error) {
	s.position++
	err = check(s.fileContent[s.position], "{")
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "color1")
	if err != nil {
		return
	}

	s.position++
	checker.color1, err = s.readColor()
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "color2")
	if err != nil {
		return
	}

	s.position++
	checker.color2, err = s.readColor()
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "scale")
	if err != nil {
		return
	}

	s.position++
	checker.scale, err = strconv.ParseFloat(s.fileContent[s.position], 64)
	if err != nil {
		return
	}

	s.position++
	err = check(s.fileContent[s.position], "}")
	if err != nil {
		return
	}

	s.position++
	return
}
