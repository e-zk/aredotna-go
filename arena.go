package aredotna

import (
	"encoding/json"
	//	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const baseUrl = "https://api.are.na/v2/"

type Arena struct {
	key string
}

func New(key string) *Arena {
	return &Arena{key: key}
}

func (a *Arena) getPaginated(page int, per int, end ...string) ([]byte, error) {
	c := http.Client{
		Timeout: time.Second * 5,
	}

	reqUrl, _ := url.JoinPath(baseUrl, end...)

	v := url.Values{}
	if page != 0 {
		v.Set("page", string(page))
	}
	if per != 0 {
		v.Set("per", string(page))
	}
	reqUrl = reqUrl + "?" + v.Encode()

	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
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

func (a *Arena) get(end ...string) ([]byte, error) {
	return a.getPaginated(0, 0, end...)
}

func (a *Arena) GetChannel(slug string) (*ApiChannelResp, error) {
	b, err := a.get("channels", slug, "thumb")
	if err != nil {
		return nil, err
	}

	ch := ApiChannelResp{}
	err = json.Unmarshal(b, &ch)
	if err != nil {
		return nil, err
	}

	return &ch, nil
}

// get channel contents (per = blocks per page; page = page)
func (a *Arena) GetChannelContents(slug string, per int, page int) ([]ApiChannelBlock, error) {
	b, err := a.getPaginated(per, page, "channels", slug)
	if err != nil {
		return nil, err
	}

	contents := []ApiChannelBlock{}
	err = json.Unmarshal(b, &contents)
	if err != nil {
		return nil, err
	}

	return contents, nil
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
