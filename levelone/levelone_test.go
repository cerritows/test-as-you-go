package levelone

import "testing"

func Test_Jump(t *testing.T) {
	start := 2
	height := 5
	expect := 7

	land := Jump(start, height)

	if land != expect {
		t.Errorf("error jumping, expected %d, got %d", expect, land)
	}
}

func Test_JumpMore(t *testing.T) {
	tests := []struct {
		name      string
		argStart  int
		argHeight int
		expect    int
	}{
		{
			name:      "test 1",
			argStart:  2,
			argHeight: 3,
			expect:    5,
		},
		{
			name:      "test 2",
			argStart:  -22,
			argHeight: 10,
			expect:    -12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Jump(tt.argStart, tt.argHeight)
			if res != tt.expect {
				t.Errorf("error jumping, expected %d, got %d", tt.expect, res)
			}
		})
	}
}
