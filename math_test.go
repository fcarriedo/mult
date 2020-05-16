package mult

import "testing"

func TestMult(t *testing.T) {
	if res := Mult(5, 5); res != 25 {
		t.Errorf("Expected %d but got %d", 25, res)
	}
}
