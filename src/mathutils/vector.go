// Package mathutils provides some mathematical utilities used in the raytracer.
package mathutils

import "math"

// Defines a 3-dimentional vector.
type Vector struct {
	X, Y, Z float64
}

// Create a new Vector with x, y, z coordinates and return it.
func NewVector(x, y, z float64) Vector {
	return Vector{x, y, z}
}

// Return the legth of v.
func (v *Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Return the squared length of v.
func (v *Vector) LengthSqr() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Add rhs to v.
func (v *Vector) Add(rhs Vector) {
	*v = Vector{v.X + rhs.X, v.Y + rhs.Y, v.Z + rhs.Z}
}

// Mutiply v by mul.
func (v *Vector) Multiply(mul float64) {
	v.X *= mul
	v.Y *= mul
	v.Z *= mul
}

// Multiply v by -1.
func (v *Vector) UnaryMinus() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

// Normalize v. (Its length will be equal to 1)
func (v *Vector) Normalize() {
	mul := 1.0 / v.Length()
	v.Multiply(mul)
}

// Return the dot product of lhs and rhs.
func DotProduct(lhs, rhs Vector) float64 {
	return lhs.X*rhs.X + lhs.Y*rhs.Y + lhs.Z*rhs.Z
}

// Return the cross product of lhs and rhs.
func CrossProduct(lhs, rhs Vector) Vector {
	return Vector{lhs.Y*rhs.Z - lhs.Z*rhs.Y, lhs.Z*rhs.X - lhs.X*rhs.Z, lhs.X*rhs.Y - lhs.Y*rhs.X}
}

// Substract rhs from lhs and return a vector with the result.
func VectorSubstraction(lhs, rhs Vector) Vector {
	return Vector{lhs.X - rhs.X, lhs.Y - rhs.Y, lhs.Z - rhs.Z}
}

// Add lhs and rhs and return a vector with the result.
func VectorAddition(lhs, rhs Vector) Vector {
	return Vector{lhs.X + rhs.X, lhs.Y + rhs.Y, lhs.Z + rhs.Z}
}

// Multiply lhs by rhs and return a vector with the result.
func VectorMultiply(lhs Vector, rhs float64) Vector {
	return Vector{lhs.X * rhs, lhs.Y * rhs, lhs.Z * rhs}
}

// Reflect the in vector across the normal and return the resulting vector.
func Reflect(in, normal Vector) Vector {
	in.Normalize()
	result := in
	normal.Multiply(2)
	in.UnaryMinus()
	result.Add(VectorMultiply(normal, DotProduct(normal, in)))
	result.Normalize()

	return result
}

// Change the orientation of normal so it points to the light source.
func Faceforward(ray, normal Vector) Vector {
	if DotProduct(ray, normal) < 0 {
		return normal
	}

	normal.UnaryMinus()
	return normal
}
