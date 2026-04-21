package service

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{"divisível por 3 e 5", 15, "FizzBuzz"},
		{"divisível apenas por 3", 9, "Fizz"},
		{"divisível apenas por 5", 10, "Buzz"},
		{"não divisível por 3 nem 5", 1, "1"},
		{"múltiplo maior", 30, "FizzBuzz"},
		{"zero", 0, "FizzBuzz"}, // 0 é divisível por 3 e 5
		{"negativo não divisível", -1, "The number must be positive."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FizzBuzz(tt.num)
			if got != tt.want {
				t.Errorf("fizzBuzz(%d) = %q, want %q", tt.num, got, tt.want)
			}
		})
	}
}
