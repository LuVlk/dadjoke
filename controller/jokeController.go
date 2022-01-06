package controller

import "github.com/LuVlk/dadjoke/model"

type JokeController interface {
	GetRandomJoke() (*model.Joke, error)
}

func MakeJokeController() JokeController {
	return &jokeRESTController{}
}
