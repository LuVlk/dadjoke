package controller

import (
	"fmt"
	"github.com/LuVlk/dadjoke/model"
	"io/ioutil"
	"net/http"
)

type jokeRESTController struct {
}

func (c *jokeRESTController) GetRandomJoke() (*model.Joke, error) {
	jokeData, err := c.getRandomJokeAsJson()
	if err != nil {
		return nil, fmt.Errorf("could not get joke data - %w", err)
	}

	joke, err := model.MakeJokeFromJson(jokeData)
	if err != nil {
		return nil, fmt.Errorf("could create joke from response - %w", err)
	}

	return joke, nil
}

func (c *jokeRESTController) getRandomJokeAsJson() ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, "https://icanhazdadjoke.com/", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Dadjoke CLI (github.com/LuVlk/dadjoke)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, err
}
