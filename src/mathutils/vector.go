package mathutils

import "math"

type Vector struct {
	X, Y, Z float64
}

func NewVector(x, y, z float64) Vector {
	return Vector{x, y, z}
}

func (v *Vector) Lenght() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector) Add(rhs Vector) {
	*v = Vector{v.X + rhs.X, v.Y + rhs.Y, v.Z + rhs.Z}
}

func (v *Vector) Multiply(mul float64) {
	v.X *= mul
	v.Y *= mul
	v.Z *= mul
}

func (v *Vector) UnaryMinus() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

func (v *Vector) Normalize() {
	mul := 1.0 / v.Lenght()
	v.Multiply(mul)
}

func DotProduct(lhs, rhs Vector) float64 {
	return lhs.X*rhs.X + lhs.Y*rhs.Y + lhs.Z*rhs.Z
}

func CrossProduct(lhs, rhs Vector) Vector {
	return Vector{lhs.Y*rhs.Z - lhs.Z*rhs.Y, lhs.Z*rhs.X - lhs.X*rhs.Z, lhs.X*rhs.Y - lhs.Y*rhs.X}
}

func VectorSubstraction(lhs, rhs Vector) Vector {
	return Vector{lhs.X - rhs.X, lhs.Y - rhs.Y, lhs.Z - rhs.Z}
}
