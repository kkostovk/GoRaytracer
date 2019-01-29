package mathutils

import (
	"math"
	"testing"
)

func TestToRadians(t *testing.T) {
	result := ToRadians(360.0)
	if math.Abs(result-math.Pi*2) > 1e-6 {
		t.Errorf("ToRadians failed!")
	}

	result = ToRadians(180.0)
	if math.Abs(result-math.Pi) > 1e-6 {
		t.Errorf("ToRadians failed!")
	}

	result = ToRadians(90.0)
	if math.Abs(result-math.Pi/2) > 1e-6 {
		t.Errorf("ToRadians failed!")
	}

	result = ToRadians(60.0)
	if math.Abs(result-math.Pi/3) > 1e-6 {
		t.Errorf("ToRadians failed!")
	}

	result = ToRadians(45.0)
	if math.Abs(result-math.Pi/4) > 1e-6 {
		t.Errorf("ToRadians failed!")
	}

	result = ToRadians(36.0)
	if math.Abs(result-math.Pi/5) > 1e-6 {
		t.Errorf("ToRadians failed!")
	}

	result = ToRadians(30.0)
	if math.Abs(result-math.Pi/6) > 1e-6 {
		t.Errorf("ToRadians failed!")
	}

	result = ToRadians(15.0)
	if math.Abs(result-math.Pi/12) > 1e-6 {
		t.Errorf("ToRadians failed!")
	}

	result = ToRadians(0.0)
	if result != 0.0 {
		t.Errorf("ToRadians failed!")
	}
}

func TestToDegrees(t *testing.T) {
	result := ToDegrees(math.Pi * 2)
	if math.Abs(result-360.0) > 1e-6 {
		t.Errorf("ToDegrees failed!")
	}

	result = ToDegrees(math.Pi)
	if math.Abs(result-180.0) > 1e-6 {
		t.Errorf("ToDegrees failed!")
	}

	result = ToDegrees(math.Pi / 2)
	if math.Abs(result-90.0) > 1e-6 {
		t.Errorf("ToDegrees failed!")
	}

	result = ToDegrees(math.Pi / 3)
	if math.Abs(result-60.0) > 1e-6 {
		t.Errorf("ToDegrees failed!")
	}

	result = ToDegrees(math.Pi / 4)
	if math.Abs(result-45.0) > 1e-6 {
		t.Errorf("ToDegrees failed!")
	}

	result = ToDegrees(math.Pi / 5)
	if math.Abs(result-36.0) > 1e-6 {
		t.Errorf("ToDegrees failed!")
	}

	result = ToDegrees(math.Pi / 6)
	if math.Abs(result-30.0) > 1e-6 {
		t.Errorf("ToDegrees failed!")
	}

	result = ToDegrees(math.Pi / 12)
	if math.Abs(result-15.0) > 1e-6 {
		t.Errorf("ToDegrees failed!")
	}

	result = ToDegrees(0.0)
	if result != 0.0 {
		t.Errorf("ToDegrees failed!")
	}
}
