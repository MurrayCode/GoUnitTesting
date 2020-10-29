package Calculator

import(
	"testing"
)

func TestAddition(t *testing.T){
	got := Addition(2,2)
	want := 4

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestMultiplication(t *testing.T){
	got := Multiplication(2,2)
	want := 4

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDivision(t *testing.T){
	got := Division(2,2)
	want := 1

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSubtraction(t *testing.T){
	got := Subtraction(2,2)
	want := 0

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}