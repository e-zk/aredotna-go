package aredotna

// based on api responses

import (
	"time"
)

type UserAvatar struct {
	Thumb   string `json:"thumb"`
	Display string `json:"display"`
}

type User struct {
	CreatedAt      time.Time  `json:"created_at"`
	Slug           string     `json:"slug"`
	Username       string     `json:"username"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	FullName       string     `json:"full_name"`
	Avatar         string     `json:"avatar"`
	AvatarImage    UserAvatar `json:"avatar_image"`
	ChannelCount   int        `json:"channel_count"`
	FollowingCount int        `json:"following_count"`
	ProfileId      int        `json:"profile_id"`
	FollowerCount  int        `json:"follower_count"`
	Initials       string     `json:"initials"`
	CanIndex       bool       `json:"can_index"`
	Metadata       *struct {
		Description string `json:"description,omitempty"`
	} `json:"metadata,omitempty"`
	IsPremium                   bool   `json:"is_premium"`
	IsLifetimePremium           bool   `json:"is_lifetime_premium"`
	IsSupporter                 bool   `json:"is_supporter"`
	IsExceedingConnectionsLimit bool   `json:"is_exceeding_connections_limit"`
	IsConfirmed                 bool   `json:"is_confirmed"`
	IsPendingReconfirmation     bool   `json:"is_pending_reconfirmation"`
	IsPendingConfirmation       bool   `json:"is_pending_confirmation"`
	Badge                       string `json:"badge"`
	Id                          int    `json:"id"`
	BaseClass                   string `json:"base_class"`
	Class                       string `json:"class"`
}

type UserChannels struct {
	Id         int       `json:"id"`
	Length     int       `json:"length"`
	TotalPages int       `json:"total_pages"`
	Page       int       `json:"current_page"`
	Per        int       `json:"per"`
	BaseClass  string    `json:"base_class"`
	Class      string    `json:"class"`
	Channels   []Channel `json:"channels"`
}

// Slug or Id works here
func (a *Arena) GetUser(slug string) (u *User, err error) {
	path := "users/" + slug
	err = a.get(path, Parameters{}, &u)
	if err != nil {
		return nil, err
	}
	return
}

func (a *Arena) GetUserChannels(slug string, params Parameters) (uc *UserChannels, err error) {
	path := "users/" + slug + "/channels"
	err = a.get(path, Parameters{}, &uc)
	if err != nil {
		return nil, err
	}
	return
}
