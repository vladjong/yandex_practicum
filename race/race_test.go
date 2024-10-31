package race

import "testing"

func TestRace(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "race",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Race()
		})
	}
}
