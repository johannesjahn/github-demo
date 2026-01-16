package fibonacci

import (
	"reflect"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "Zero",
			n:    0,
			want: []int{},
		},
		{
			name: "Negative",
			n:    -5,
			want: []int{},
		},
		{
			name: "One",
			n:    1,
			want: []int{0},
		},
		{
			name: "Two",
			n:    2,
			want: []int{0, 1},
		},
		{
			name: "Five",
			n:    5,
			want: []int{0, 1, 1, 2, 3},
		},
		{
			name: "Ten",
			n:    10,
			want: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
