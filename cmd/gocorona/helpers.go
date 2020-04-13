package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

func fetch(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Something happened")
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
