package controller

import (
	"fmt"
	"github.com/LuVlk/dadjoke/config"
	"github.com/LuVlk/dadjoke/model"
	"io/ioutil"
	"net/http"
	"strconv"
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
	req, err := http.NewRequest(http.MethodGet, config.BasePath(), nil)
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

func (c *jokeRESTController) GetJokesByTerm(term string, limit int) ([]model.Joke, error) {
	data, err := c.getJokesByTermAsJson(term, limit)
	if err != nil {
		return nil, fmt.Errorf("could not get joke data - %w", err)
	}

	jokes, err := model.MakeJokesFromJson(data)
	if err != nil {
		return nil, fmt.Errorf("could not create jokes from json - %w", err)
	}

	return jokes, nil
}

func (c *jokeRESTController) getJokesByTermAsJson(term string, limit int) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, config.BasePath()+"/search", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Dadjoke CLI (github.com/LuVlk/dadjoke)")

	query := req.URL.Query()
	query.Add("term", term)
	query.Add("limit", strconv.Itoa(limit))
	req.URL.RawQuery = query.Encode()

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
