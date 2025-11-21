package joke

import (
	"testing"
)

func Test_jokeService_GetJoke(t *testing.T) {
	tests := []struct {
		name string
		s    Service
	}{
		{
			name: "returns one of the jokes from database",
			s:    NewService(),
		},
	}
	
	validJokes := map[string]bool{
		"1": true,
		"2": true,
		"3": true,
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GetJoke()
			if !validJokes[got.ID] {
				t.Errorf("jokeService.GetJoke() returned joke with invalid ID = %v", got.ID)
			}
			if got.Value == "" {
				t.Errorf("jokeService.GetJoke() returned joke with empty value")
			}
		})
	}
}
