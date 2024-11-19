package aredotna

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Parameters map[string]string

func (p Parameters) Values() url.Values {
	v := url.Values{}
	for k, v := range Parameters {
		v.Set(k, v)
	}
	return v
}

type Arena struct {
	Client  http.Client
	BaseURL string
	key     string
}

func New(key string) *Arena {
	return &Arena{
		BaseURL: "https://api.are.na/v2",
		Client: &http.Client{
			Timeout: a.Timeout,
		},
		key: key,
	}
}

func (a *Arena) get(path string, params Parameters, target any) error {

	reqURL, _ := url.JoinPath(a.BaseURL, path)
	reqURL = reqUrl + "?" + params.Encode()

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return errors.Wrapf(err, "invalid request %q", url)
	}

	req.Header.Add("Authorization", "Bearer "+a.key)
	req.Header.Add("Content-Type", "application/json")

	res, err := c.Do(req)
	if err != nil {
		return errors.Wrapf(err, "executing request for %q failed", reqURL)
	}

	if res.Body == nil {
		return fmt.Errorf("no data returned for %q", reqURL)
	}

	x, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	err = json.Unmarshal(b, target)
	if err != nil {
		return errors.Wrapf(err, "json unmarshal error for %q", reqURL)
	}

	return

}

/*
func (a *Arena) GetBlock(id string) (*Block, error) {
	b, err := a.get("blocks/", id)
	fmt.Printf("%s", string(b))
	if err != nil {
		return nil, err
	}

	block := Block{}
	err = json.Unmarshal(b, &block)
	if err != nil {
		return nil, err
	}

	return &block, nil
}*/

func (a *Arena) GetUser(id string) (*ApiUser, error) {
	b, err := a.get("users/", id)
	if err != nil {
		return nil, err
	}

	u := ApiUser{}
	err = json.Unmarshal(b, &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

/*
func (a *Arena) GetGroup(slug string) (*Group, error) {
	b, err := a.get("groups/", slug)
	if err != nil {
		return nil, err
	}

	g := Group{}
	err = json.Unmarshal(b, &g)
	if err != nil {
		return nil, err
	}

	return &g, nil
}*/
