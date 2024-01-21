package aredotna

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"time"
)

const baseUrl = "https://api.are.na/v2/"

type Arena struct {
	key string
}

func New(key string) *Arena {
	return &Arena{key: key}
}

func (a *Arena) get(end ...string) ([]byte, error) {
	c := http.Client{
		Timeout: time.Second * 5,
	}

	url := baseUrl + path.Join(end...)
	//	log.Printf("url=> %q", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Add("Authorization", "Bearer "+a.key)
	req.Header.Add("Content-Type", "application/json")

	res, err := c.Do(req)
	if err != nil {
		return []byte{}, err
	}

	if res.Body == nil {
		return []byte{}, err
	}

	x, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return x, nil

}

func (a *Arena) GetChannel(slug string) (*Channel, error) {
	b, err := a.get("channels/", slug)
	if err != nil {
		return nil, err
	}

	ch := Channel{}
	err = json.Unmarshal(b, ch)
	if err != nil {
		return nil, err
	}

	return &ch, nil
}

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
}

func (a *Arena) GetUser(id string) (*User, error) {
	b, err := a.get("users/", id)
	if err != nil {
		return nil, err
	}

	u := User{}
	err = json.Unmarshal(b, &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

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
}
