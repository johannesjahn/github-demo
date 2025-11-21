package joke

import (
	"math/rand"
	"time"
)

type Joke struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type Service interface {
	GetJoke() Joke
}

type jokeService struct{}

func NewService() Service {
	return &jokeService{}
}

var jokeDb = []Joke{
	Joke{
		ID:    "1",
		Value: "Why don't scientists trust atoms? Because they make up everything!",
	},
	Joke{
		ID:    "2",
		Value: "Why did the scarecrow win an award? Because he was outstanding in his field!",
	},
	Joke{
		ID:    "3",
		Value: "What do you call a bear with no teeth? A gummy bear!",
	},
}

func (s *jokeService) GetJoke() Joke {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(jokeDb))
	return jokeDb[randomIndex]
}
