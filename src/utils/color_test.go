package utils

import (
	"math"
	"testing"
)

func TestNewColor(t *testing.T) {
	color := NewColor(1, 128, 255)

	if math.Abs(color[0]-1.0/255.0) > 1e-10 || math.Abs(color[1]-128.0/255.0) > 1e-10 || math.Abs(color[2]-255.0/255.0) > 1e-10 {
		t.Errorf("NewColor() failed!")
	}
}

func TestToRGB(t *testing.T) {
	color := NewColor(0, 128, 255)
	r, g, b := color.ToRGB()
	if r != 0 || g != 128 || b != 255 {
		t.Errorf("ToRGB() failed!")
	}
}

func TestColorAddition(t *testing.T) {
	color1 := NewColor(64, 8, 13)
	color2 := NewColor(14, 11, 32)
	expected := NewColor(78, 19, 45)

	res := ColorAddition(color1, color2)
	if math.Abs(res[0]-expected[0]) > 1e-10 || math.Abs(res[1]-expected[1]) > 1e-10 || math.Abs(res[2]-expected[2]) > 1e-10 {
		t.Errorf("ColorAddition() failed!")
	}
}

func TestColorMultiplication(t *testing.T) {
	color1 := NewColor(6, 6, 6)
	color2 := NewColor(255, 255, 255)
	expected := NewColor(6, 6, 6)

	res := ColorMultiplication(color1, color2)
	if math.Abs(res[0]-expected[0]) > 1e-10 || math.Abs(res[1]-expected[1]) > 1e-10 || math.Abs(res[2]-expected[2]) > 1e-10 {
		t.Errorf("ColorMultiplication() failed!")
	}
}

func TestMultiplyColorFloat(t *testing.T) {
	color := NewColor(8, 47, 64)
	color = MultiplyColorFloat(color, 3)
	r, g, b := color.ToRGB()
	if r != 24 || g != 141 || b != 192 {
		t.Errorf("MultiplyColorFloat() failed!")
	}
}

func TestDivideColorFloat(t *testing.T) {
	color := NewColor(39, 126, 207)
	color = DivideColorFloat(color, 3)
	r, g, b := color.ToRGB()
	if r != 13 || g != 42 || b != 69 {
		t.Errorf("DivideColorFloat() failed!")
	}
}
