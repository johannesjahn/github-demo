package joke

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
}

func (s *jokeService) GetJoke() Joke {

	var result = jokeDb[0]

	return result
}
