package mathutils

import (
	"math"
	"testing"
)

func TestNewVector(t *testing.T) {
	vec1 := NewVector(0, 0, 0)
	vec2 := NewVector(1, 1, 1)
	vec3 := NewVector(0, 1, 2)

	if vec1.X != 0 || vec1.Y != 0 || vec1.Z != 0 {
		t.Errorf("NewVector failed!")
	}

	if vec2.X != 1 || vec2.Y != 1 || vec2.Z != 1 {
		t.Errorf("NewVector failed!")
	}

	if vec3.X != 0 || vec3.Y != 1 || vec3.Z != 2 {
		t.Errorf("NewVector failed!")
	}
}

func TestLength(t *testing.T) {
	vec1 := NewVector(0, 0, 0)
	if vec1.Length() != 0 {
		t.Errorf("Vector.Length() failed!")
	}

	vec2 := NewVector(3, 0, 0)
	if vec2.Length() != 3 {
		t.Errorf("Vector.Length() failed!")
	}

	vec3 := NewVector(4, 3, 0)
	if vec3.Length() != 5 {
		t.Errorf("Vector.Length() failed!")
	}
}

func TestLengthSqr(t *testing.T) {
	vec1 := NewVector(0, 0, 0)
	if vec1.LengthSqr() != 0 {
		t.Errorf("Vector.LengthSqr() failed!")
	}

	vec2 := NewVector(3, 0, 0)
	if vec2.LengthSqr() != 9 {
		t.Errorf("Vector.LengthSqr() failed!")
	}

	vec3 := NewVector(4, 3, 0)
	if vec3.LengthSqr() != 25 {
		t.Errorf("Vector.LengthSqr() failed!")
	}
}

func TestAdd(t *testing.T) {
	vec1 := NewVector(0, 0, 0)
	vec2 := NewVector(1, 2, 3)
	vec3 := NewVector(7, 0, -2)

	resVec := vec1

	resVec.Add(vec1)
	if resVec != vec1 {
		t.Errorf("Vector.Add() failed!")
	}

	resVec = vec2
	resVec.Add(vec2)
	if resVec != NewVector(2, 4, 6) {
		t.Errorf("Vector.Add() failed!")
	}

	resVec = vec2
	resVec.Add(vec3)
	if resVec != NewVector(8, 2, 1) {
		t.Errorf("Vector.Add() failed!")
	}
}

func TestMutiply(t *testing.T) {
	vec1 := NewVector(0, 0, 0)
	vec2 := NewVector(1, 1, 1)
	vec3 := NewVector(1, 2, 4)

	vec1.Multiply(4)
	if vec1 != NewVector(0, 0, 0) {
		t.Errorf("Vector.Mutiply() failed!")
	}

	vec2.Multiply(6)
	if vec2 != NewVector(6, 6, 6) {
		t.Errorf("Vector.Mutiply() failed!")
	}

	vec3.Multiply(2)
	if vec3 != NewVector(2, 4, 8) {
		t.Errorf("Vector.Mutiply() failed!")
	}
}

func TestUnaryMinus(t *testing.T) {
	vec1 := NewVector(0, 0, 0)
	vec2 := NewVector(1, 1, 1)
	vec3 := NewVector(1, 2, 4)

	vec1.UnaryMinus()
	if vec1 != NewVector(0, 0, 0) {
		t.Errorf("Vector.UnaryMinus() failed!")
	}

	vec2.UnaryMinus()
	if vec2 != NewVector(-1, -1, -1) {
		t.Errorf("Vector.UnaryMinus() failed!")
	}

	vec3.UnaryMinus()
	if vec3 != NewVector(-1, -2, -4) {
		t.Errorf("Vector.UnaryMinus() failed!")
	}
}

func TestNormalize(t *testing.T) {
	vec1 := NewVector(4, 0, 0)
	vec1Res := NewVector(1, 0, 0)
	vec2 := NewVector(4, 4, 0)
	vec2Res := NewVector(1/math.Sqrt(2), 1/math.Sqrt(2), 0)
	vec3 := NewVector(6, 6, 6)
	vec3Res := NewVector(1/math.Sqrt(3), 1/math.Sqrt(3), 1/math.Sqrt(3))

	vec1.Normalize()
	if math.Abs(vec1.X-vec1Res.X) > 1e-10 || math.Abs(vec1.Y-vec1Res.Y) > 1e-10 || math.Abs(vec1.Z-vec1Res.Z) > 1e-10 {
		t.Errorf("Vector.Normalize() failed!")
	}

	vec2.Normalize()
	if math.Abs(vec2.X-vec2Res.X) > 1e-10 || math.Abs(vec2.Y-vec2Res.Y) > 1e-10 || math.Abs(vec2.Z-vec2Res.Z) > 1e-10 {
		t.Errorf("Vector.Normalize() failed!")
	}

	vec3.Normalize()
	if math.Abs(vec3.X-vec3Res.X) > 1e-10 || math.Abs(vec3.Y-vec3Res.Y) > 1e-10 || math.Abs(vec3.Z-vec3Res.Z) > 1e-10 {
		t.Errorf("Vector.Normalize() failed!")
	}

}

func TestDotProduct(t *testing.T) {
	vec1 := NewVector(0, 0, 0)
	vec2 := NewVector(1, 2, 4)
	vec3 := NewVector(10, 0, -3)
	vec4 := NewVector(0, 13, 0)

	if DotProduct(vec1, vec1) != 0 {
		t.Errorf("DotProduct() failed!")
	}

	if DotProduct(vec2, vec3) != -2 {
		t.Errorf("DotProduct() failed!")
	}

	if DotProduct(vec2, vec4) != 26 {
		t.Errorf("DotProduct() failed!")
	}

	if DotProduct(vec3, vec4) != 0 {
		t.Errorf("DotProduct() failed!")
	}
}

func TestCrossProduct(t *testing.T) {
	resVec := CrossProduct(NewVector(1, 1, 1), NewVector(1, 1, 1))
	if resVec != NewVector(0, 0, 0) {
		t.Errorf("CrossProduct() failed!")
	}

	resVec = CrossProduct(NewVector(0, 1, 0), NewVector(4, 2, 0))
	if resVec != NewVector(0, 0, -4) {
		t.Errorf("CrossProduct() failed!")
	}

	resVec = CrossProduct(NewVector(4, 2, 1), NewVector(1, 7, 3))
	if resVec != NewVector(-1, -11, 26) {
		t.Errorf("CrossProduct() failed!")
	}
}

func TestVectorSubstraction(t *testing.T) {
	resVec := VectorSubstraction(NewVector(1, 1, 1), NewVector(1, 1, 1))
	if resVec != NewVector(0, 0, 0) {
		t.Errorf("VectorSubstraction() failed!")
	}

	resVec = VectorSubstraction(NewVector(1, 2, 4), NewVector(-1, -1, -1))
	if resVec != NewVector(2, 3, 5) {
		t.Errorf("VectorSubstraction() failed!")
	}

	resVec = VectorSubstraction(NewVector(5, 8, 13), NewVector(21, 34, 55))
	if resVec != NewVector(-16, -26, -42) {
		t.Errorf("VectorSubstraction() failed!")
	}
}

func TestVectorAddition(t *testing.T) {
	resVec := VectorAddition(NewVector(0, 0, 0), NewVector(0, 0, 0))
	if resVec != NewVector(0, 0, 0) {
		t.Errorf("VectorAddition() failed!")
	}

	resVec = VectorAddition(NewVector(1, 2, 4), NewVector(-1, -1, -1))
	if resVec != NewVector(0, 1, 3) {
		t.Errorf("VectorAddition() failed!")
	}

	resVec = VectorAddition(NewVector(5, 8, 13), NewVector(21, 34, 55))
	if resVec != NewVector(26, 42, 68) {
		t.Errorf("VectorAddition() failed!")
	}
}

func TestVectorMultiply(t *testing.T) {
	resVec := VectorMultiply(NewVector(0, 0, 0), 42)
	if resVec != NewVector(0, 0, 0) {
		t.Errorf("VectorMultiply() failed!")
	}

	resVec = VectorMultiply(NewVector(1, 1, 1), 42)
	if resVec != NewVector(42, 42, 42) {
		t.Errorf("VectorMultiply() failed!")
	}

	resVec = VectorMultiply(NewVector(2, 4, 16), 0.5)
	if resVec != NewVector(1, 2, 8) {
		t.Errorf("VectorMultiply() failed!")
	}
}
