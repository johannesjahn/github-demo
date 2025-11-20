package joke

import (
	"reflect"
	"testing"
)

func Test_jokeService_GetJoke(t *testing.T) {
	tests := []struct {
		name string
		s    Service
		want Joke
	}{
		{
			name: "default joke",
			s:    NewService(),
			want: Joke{
				ID:    "1",
				Value: "Why don't scientists trust atoms? Because they make up everything!",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetJoke(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jokeService.GetJoke() = %v, want %v", got, tt.want)
			}
		})
	}
}
