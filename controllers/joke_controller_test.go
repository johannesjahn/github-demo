package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"johannes-jahn.com/demo/joke"
)

type mockJokeService struct{}

func (m *mockJokeService) GetJoke() joke.Joke {
	return joke.Joke{
		ID:    "1",
		Value: "a joke",
	}
}

func TestJokeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/joke", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := jokeHandler(&mockJokeService{})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"id":"1","value":"a joke"}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}