package raytracer

import (
	"GoRaytracer/src/mathutils"
	"math"
)

// Camera provides an interface for getting screen rays from cameras.
type Camera interface {
	GetScreenRay(x, y float64) Ray
}

// ParallelCamera defines a pinhole camera.
type ParallelCamera struct {
	position   mathutils.Vector // The position of the camera.
	topLeft    mathutils.Vector // The top left corner of the screen view.
	topRight   mathutils.Vector // The top right corner of the screen view.
	bottomLeft mathutils.Vector // The bottom left corner of the screen view.
}

// NewParallelCamera creates and returns a new pinhole camera.
func NewParallelCamera(position mathutils.Vector, yaw, pitch, roll, fov, aspectRatio float64) ParallelCamera {
	x2d := aspectRatio
	y2d := 1.0
	wantedAngle := mathutils.ToRadians(fov / 2.0)
	wantedLength := math.Tan(wantedAngle)
	hypotLength := math.Sqrt(aspectRatio*aspectRatio + 1.0)
	scaleFactor := wantedLength / hypotLength

	x2d *= scaleFactor * 1.5
	y2d *= scaleFactor * 1.5

	topLeft := mathutils.NewVector(-x2d, y2d, 1)
	topRight := mathutils.NewVector(x2d, y2d, 1)
	bottomLeft := mathutils.NewVector(-x2d, -y2d, 1)

	rotAroundX := mathutils.RotationAroundX(mathutils.ToRadians(roll))
	rotAroundY := mathutils.RotationAroundY(mathutils.ToRadians(pitch))
	rotAroundZ := mathutils.RotationAroundZ(mathutils.ToRadians(yaw))
	rotation := mathutils.MatrixMultiplication((mathutils.MatrixMultiplication(rotAroundX, rotAroundY)), rotAroundZ)

	topLeft = mathutils.MultiplyVectorMatrix(topLeft, rotation)
	topRight = mathutils.MultiplyVectorMatrix(topRight, rotation)
	bottomLeft = mathutils.MultiplyVectorMatrix(bottomLeft, rotation)

	topLeft.Add(position)
	topRight.Add(position)
	bottomLeft.Add(position)

	return ParallelCamera{position, topLeft, topRight, bottomLeft}
}

// GetScreenRay return the screen ray for the given coordinates.
func (c *ParallelCamera) GetScreenRay(x, y float64) Ray {
	direction := c.topLeft
	width := mathutils.VectorSubstraction(c.topRight, c.topLeft)
	height := mathutils.VectorSubstraction(c.bottomLeft, c.topLeft)
	width.Multiply(x / 640.0)
	height.Multiply(y / 480.0)

	direction.Add(width)
	direction.Add(height)
	direction = mathutils.VectorSubstraction(direction, c.position)
	direction.Normalize()

	return Ray{c.position, direction}
}
