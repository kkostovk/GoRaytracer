package raytracer

import (
	"GoRaytracer/src/mathutils"
	"math"
)

//Plane orientation
const (
	XY = iota
	XZ
	YZ
)

type IntersectionInfo struct {
	Position mathutils.Vector
	Normal   mathutils.Vector
	Distance float64
	U, V     float64
}

type Geometry interface {
	Intersect(*Ray, *IntersectionInfo) bool
}

type Plane struct {
	center      mathutils.Vector
	limit       float64
	orientation uint8
}

func NewPlane(center mathutils.Vector, limit float64, orientation uint8) Plane {
	return Plane{center, limit, orientation}
}

func (p *Plane) Intersect(ray *Ray, info *IntersectionInfo) bool {

	var start, direction, plane float64

	if p.orientation == XY {
		start = ray.Start.Z
		direction = ray.Direction.Z
		plane = p.center.Z
	} else if p.orientation == XZ {
		start = ray.Start.Y
		direction = ray.Direction.Y
		plane = p.center.Y
	} else {
		start = ray.Start.X
		direction = ray.Direction.X
		plane = p.center.X
	}

	if direction >= 0.0 && start > plane || direction <= 0.0 && start < plane {
		return false
	}

	multiplayer := (start - plane) / -direction

	info.Position = ray.Direction
	info.Position.Multiply(multiplayer)
	info.Position = mathutils.VectorAddition(info.Position, ray.Start)
	info.Distance = multiplayer

	if p.orientation == XY {
		if math.Abs(p.center.X-info.Position.X) > p.limit/2 || math.Abs(p.center.Y-info.Position.Y) > p.limit/2 {
			return false
		}
		info.U = info.Position.X
		info.V = info.Position.Y
		info.Normal = mathutils.NewVector(0, 0, 1)
	} else if p.orientation == XZ {
		if math.Abs(p.center.X-info.Position.X) > p.limit/2 || math.Abs(p.center.Z-info.Position.Z) > p.limit/2 {
			return false
		}
		info.U = info.Position.X
		info.V = info.Position.Z
		info.Normal = mathutils.NewVector(0, 1, 0)
	} else {
		if math.Abs(p.center.Y-info.Position.Y) > p.limit/2 || math.Abs(p.center.Z-info.Position.Z) > p.limit/2 {
			return false
		}
		info.U = info.Position.Y
		info.V = info.Position.Z
		info.Normal = mathutils.NewVector(1, 0, 0)
	}

	if mathutils.DotProduct(ray.Direction, info.Normal) > 0 {
		info.Normal.UnaryMinus()
	}

	return true
}

type Sphere struct {
	center mathutils.Vector
	radius float64
}

func NewSphere(center mathutils.Vector, radius float64) Sphere {
	return Sphere{center, radius}
}

func (s *Sphere) Intersect(ray *Ray, info *IntersectionInfo) bool {
	H := mathutils.VectorSubstraction(ray.Start, s.center)

	A := ray.Direction.LengthSqr()
	B := 2 * mathutils.DotProduct(H, ray.Direction)
	C := H.LengthSqr() - s.radius*s.radius

	D := B*B - 4*A*C
	if D < 0 {
		return false
	}

	x1 := (-B + math.Sqrt(D)) / (2 * A)
	x2 := (-B - math.Sqrt(D)) / (2 * A)
	if x1 < 0 && x2 < 0 {
		return false
	}

	if x2 < 0 || (x1 >= 0 && x1 < x2) {
		info.Distance = x1
	} else {
		info.Distance = x2
	}

	info.Position = mathutils.VectorAddition(ray.Start, mathutils.VectorMultiply(ray.Direction, info.Distance))
	info.Normal = mathutils.VectorSubstraction(info.Position, s.center)
	info.Normal.Normalize()
	relativePosition := mathutils.VectorSubstraction(info.Position, s.center)
	info.U = math.Atan2(relativePosition.Z, relativePosition.X)
	info.V = math.Asin(relativePosition.Y / s.radius)
	info.U = (info.U + math.Pi) / (2 * math.Pi)
	info.V = -(info.V + math.Pi/2) / math.Pi
	return true
}

type Cube struct {
	center mathutils.Vector
	edge   float64
}

func NewCube(center mathutils.Vector, edge float64) Cube {
	return Cube{center, edge}
}

func (c *Cube) intersectSide(ray *Ray, normal mathutils.Vector, info *IntersectionInfo, level, start, direction float64) bool {
	if start > level && direction >= 0 {
		return false
	}
	if start < level && direction <= 0 {
		return false
	}

	scaleFactor := (level - start) / direction
	ip := mathutils.VectorAddition(ray.Start, mathutils.VectorMultiply(ray.Direction, scaleFactor))
	if ip.X > c.center.X+c.edge/2+1e-6 || ip.X < c.center.X-c.edge/2-1e-6 {
		return false
	}
	if ip.Y > c.center.Y+c.edge/2+1e-6 || ip.Y < c.center.Y-c.edge/2-1e-6 {
		return false
	}
	if ip.Z > c.center.Z+c.edge/2+1e-6 || ip.Z < c.center.Z-c.edge/2-1e-6 {
		return false
	}

	distance := scaleFactor
	if distance < info.Distance {
		info.Position = ip
		info.Distance = distance
		info.Normal = normal
		if normal.Y == 0 {
			info.U = ip.X + ip.Z
			info.V = ip.Y
		} else {
			info.U = ip.X
			info.V = ip.Z
		}

		return true
	}

	return false
}

func (c *Cube) Intersect(ray *Ray, info *IntersectionInfo) bool {
	info.Distance = 1e99
	c.intersectSide(ray, mathutils.NewVector(-1, 0, 0), info, c.center.X-c.edge/2, ray.Start.X, ray.Direction.X)
	c.intersectSide(ray, mathutils.NewVector(+1, 0, 0), info, c.center.X+c.edge/2, ray.Start.X, ray.Direction.X)
	c.intersectSide(ray, mathutils.NewVector(0, -1, 0), info, c.center.Y-c.edge/2, ray.Start.Y, ray.Direction.Y)
	c.intersectSide(ray, mathutils.NewVector(0, +1, 0), info, c.center.Y+c.edge/2, ray.Start.Y, ray.Direction.Y)
	c.intersectSide(ray, mathutils.NewVector(0, 0, -1), info, c.center.Z-c.edge/2, ray.Start.Z, ray.Direction.Z)
	c.intersectSide(ray, mathutils.NewVector(0, 0, +1), info, c.center.Z+c.edge/2, ray.Start.Z, ray.Direction.Z)

	return info.Distance < 1e99

}
