package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func fetch(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New(string(body))
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
