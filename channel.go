package aredotna

import (
	"net/url"
	"sort"
	"time"
)

type Channel struct {
	ID                int       `json:"id"`
	Slug              string    `json:"slug"`
	Title             string    `json:"title"`
	Length            int       `json:"length"`
	Kind              string    `json:"kind"`
	Status            string    `json:"status"`
	Open              bool      `json:"open"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	AddedToAt         time.Time `json:"added_to_at"`
	Published         bool      `json:"published"`
	Collaboration     bool      `json:"collaboration"`
	CollaboratorCount int       `json:"collaborator_count"`
	UserId            int       `json:"user_id"`
	Contents          []Block   `json:"contents,omitempty"`
	ClassName         string    `json:"class_name"`      // Channel
	BaseClass         string    `json:"base_class"`      // Channel
	Class             string    `json:"class,omitempty"` // Channel
	Page              int       `json:"page"`
	Per               int       `json:"per"`
	CanIndex          bool      `json:"can_index"`
	ShareLink         string    `json:"share_link,omitempty"`
	FollowerCount     int       `json"follower_count"`
	Nsfw              bool      `json:"nsfw?,omitempty"`
	Metadata          *struct {
		Description string `json:"description,omitempty"`
	} `json:"metadata,omitempty"`
	User *User `json:"user"`
	//Owner     *User  `json:"owner,omitempty"`
	//OwnerType string `json:"owner_type,omitempty"`
	//OwnerId   int    `json:"owner_id,omitempty"`
	//OwnerSlug string `json:"owner_slug,omitempty"`
	//Group *Group `json:"group"`
	//Collaborators []User `json:"collaborators"`
}

type channelContents struct {
	Contents []Block `json:"contents"`
}

func (a *Arena) GetChannelThumb(slug string, params Parameters) (ch *Channel, err error) {
	path, _ := url.JoinPath("channels", slug, "thumb")
	err = a.get(path, params, &ch)
	if err != nil {
		return nil, err
	}

	return
}

func (a *Arena) GetChannel(slug string, params Parameters) (ch *Channel, err error) {
	err = a.get("channels/"+slug, params, &ch)
	if err != nil {
		return nil, err
	}

	sort.Slice(ch.Contents[:], func(i, j int) bool {
		return ch.Contents[i].Position > ch.Contents[j].Position
	})

	return
}

func (a *Arena) GetChannelContents(slug string, params Parameters) (b *[]Block, err error) {
	var ch channelContents
	err = a.get("channels/"+slug+"/contents", params, &ch)
	if err != nil {
		return nil, err
	}

	sort.Slice(ch.Contents[:], func(i, j int) bool {
		return ch.Contents[i].Position > ch.Contents[j].Position
	})

	return &ch.Contents, nil
}
