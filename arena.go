package aredotna

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Parameters map[string]string

func (p Parameters) Values() url.Values {
	v := url.Values{}
	for k, val := range p {
		v.Set(k, val)
	}
	return v
}

type Arena struct {
	Client  *http.Client
	BaseURL string
	key     string
}

func New(key string) *Arena {
	return &Arena{
		BaseURL: "https://api.are.na/v2",
		Client: &http.Client{
			Timeout: time.Second * 5,
		},
		key: key,
	}
}

func (a *Arena) get(path string, params Parameters, target any) error {

	reqURL, _ := url.JoinPath(a.BaseURL, path)
	reqURL = reqURL + "?" + params.Values().Encode()

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return fmt.Errorf("invalid request %q: %w", reqURL, err)
	}

	req.Header.Add("Authorization", "Bearer "+a.key)
	req.Header.Add("Content-Type", "application/json")

	res, err := a.Client.Do(req)
	if err != nil {
		return fmt.Errorf("executing request for %q failed: %w", reqURL, err)
	}

	if res.Body == nil {
		return fmt.Errorf("no data returned for %q", reqURL)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, target)
	if err != nil {
		return fmt.Errorf("json unmarshal error for %q: %w", reqURL, err)
	}

	return nil
}
