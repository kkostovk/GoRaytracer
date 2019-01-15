package mathutils

import "math"

type Matrix [3][3]float64

func NewMatrix(diagonalElement float64) Matrix {
	var result Matrix

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				result[i][j] = diagonalElement
			} else {
				result[i][j] = 0.0
			}
		}
	}

	return result
}

func RotationAroundX(angle float64) Matrix {
	result := NewMatrix(1)
	sin := math.Sin(angle)
	cos := math.Cos(angle)

	result[1][1] = cos
	result[1][2] = -sin
	result[2][1] = sin
	result[2][2] = cos

	return result
}

func RotationAroundY(angle float64) Matrix {
	result := NewMatrix(1)
	sin := math.Sin(angle)
	cos := math.Cos(angle)

	result[0][0] = cos
	result[0][2] = sin
	result[2][0] = -sin
	result[2][2] = cos

	return result
}

func RotationAroundZ(angle float64) Matrix {
	result := NewMatrix(1)
	sin := math.Sin(angle)
	cos := math.Cos(angle)

	result[0][0] = cos
	result[0][1] = -sin
	result[1][0] = sin
	result[1][1] = cos

	return result
}

func (m *Matrix) Determinant() float64 {
	positive := m[0][0]*m[1][1]*m[2][2] + m[0][1]*m[1][2]*m[2][0] + m[0][2]*m[1][0]*m[2][1]
	negative := m[0][0]*m[1][2]*m[2][1] + m[0][1]*m[1][0]*m[2][2] + m[0][2]*m[1][1]*m[2][0]

	return positive - negative
}

func MatrixMultiplication(lhs, rhs Matrix) Matrix {
	result := NewMatrix(0)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				result[i][j] += lhs[i][k] * rhs[k][j]
			}
		}
	}

	return result
}

func MultiplyVectorMatrix(v Vector, m Matrix) Vector {
	return Vector{v.X*m[0][0] + v.Y*m[1][0] + v.Z*m[2][0], v.X*m[0][1] + v.Y*m[1][1] + v.Z*m[2][1], v.X*m[0][2] + v.Y*m[1][2] + v.Z*m[2][2]}
}
