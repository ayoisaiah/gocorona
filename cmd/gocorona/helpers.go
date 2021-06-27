package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func fetch(url string, target interface{}) error {
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(url)
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
