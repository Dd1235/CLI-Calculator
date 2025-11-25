package calc

import "testing"

func TestAdd(t *testing.T) {
	if Add(3.7, 1.2) != 4.9 {
		t.Fatalf("Add: unexpected output")
	}
}

func TestSub(t *testing.T) {
	if Sub(3.7, 1.2) != 2.5 {
		t.Fatalf("Sub: unexpected output")
	}
}

func TestMul(t *testing.T) {
	if Mul(4, 3) != 12 {
		t.Fatalf("expected 12")
	}
}

func TestDiv(t *testing.T) {
	got, err := Div(10, 2)
	if err != nil || got != 5 {
		t.Fatalf("expected 5, got %v, err %v", got, err)
	}
	if _, err := Div(1, 0); err == nil {
		t.Fatalf("expected division by zero error")
	}
}
