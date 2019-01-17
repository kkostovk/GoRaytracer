package mathutils

import (
	"math"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	matr1 := NewMatrix(1)
	resMatr1 := Matrix{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	matr2 := NewMatrix(0)
	resMatr2 := Matrix{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	if matr1 != resMatr1 {
		t.Errorf("NewMatrix() failed!")
	}

	if matr2 != resMatr2 {
		t.Errorf("NewMatrix() failed!")
	}
}

func TestRotationAroundX(t *testing.T) {
	vec := NewVector(0, 1, 0)
	res := NewVector(0, 0, -1)
	rotationMatrix := RotationAroundX((90.0 / 180.0) * math.Pi)

	vec = MultiplyVectorMatrix(vec, rotationMatrix)
	if math.Abs(vec.X-res.X) > 1e-10 || math.Abs(vec.Y-res.Y) > 1e-10 || math.Abs(vec.Z-res.Z) > 1e-10 {
		t.Errorf("RotationAroundX() failed!")
	}
}

func TestRotationAroundY(t *testing.T) {
	vec := NewVector(1, 0, 0)
	res := NewVector(0, 0, 1)
	rotationMatrix := RotationAroundY((90.0 / 180.0) * math.Pi)

	vec = MultiplyVectorMatrix(vec, rotationMatrix)
	if math.Abs(vec.X-res.X) > 1e-10 || math.Abs(vec.Y-res.Y) > 1e-10 || math.Abs(vec.Z-res.Z) > 1e-10 {
		t.Errorf("RotationAroundY() failed!")
	}
}

func TestRotationAroundZ(t *testing.T) {
	vec := NewVector(1, 0, 0)
	res := NewVector(0, -1, 0)
	rotationMatrix := RotationAroundZ((90.0 / 180.0) * math.Pi)

	vec = MultiplyVectorMatrix(vec, rotationMatrix)
	if math.Abs(vec.X-res.X) > 1e-10 || math.Abs(vec.Y-res.Y) > 1e-10 || math.Abs(vec.Z-res.Z) > 1e-10 {
		t.Errorf("RotationAroundZ() failed!")
	}
}

func TestDeterminant(t *testing.T) {
	matrix := NewMatrix(1)
	if matrix.Determinant() != 1 {
		t.Errorf("Determinant() failed!")
	}

	matrix = Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	if matrix.Determinant() != 0 {
		t.Errorf("Determinant() failed!")
	}
}

func TestMatrixMultiplication(t *testing.T) {
	lhs := NewMatrix(1)
	rhs := NewMatrix(1)
	res := MatrixMultiplication(lhs, rhs)
	if res != NewMatrix(1) {
		t.Errorf("MatrixMultiplication() failed!")
	}

	lhs = Matrix{{1, 1, 2}, {3, 5, 8}, {4, 2, 0}}
	rhs = Matrix{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	expected := Matrix{{9, 9, 9}, {37, 37, 37}, {8, 8, 8}}
	res = MatrixMultiplication(lhs, rhs)

	if expected != res {
		t.Errorf("MatrixMultiplication() failed!")
	}
}

func TestMultiplyVectorMatrix(t *testing.T) {
	vec := NewVector(1, 2, 4)
	matrix := NewMatrix(1)
	res := MultiplyVectorMatrix(vec, matrix)
	if res != NewVector(1, 2, 4) {
		t.Errorf("MultiplyVectorMatrix() failed!")
	}

	vec = NewVector(4, 2, 1)
	matrix = Matrix{{1, 1, 2}, {3, 5, 8}, {13, 21, 34}}
	res = MultiplyVectorMatrix(vec, matrix)
	if res != NewVector(23, 35, 58) {
		t.Errorf("MultiplyVectorMatrix() failed!")
	}
}
