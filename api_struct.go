package aredotna

// based on api responses

import (
	"html/template"
	"time"
)

type ApiUser struct {
	CreatedAt   time.Time `json:"created_at"`
	Slug        string    `json:"slug"`
	Username    string    `json:"username"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	FullName    string    `json:"full_name"`
	Avatar      string    `json:"avatar"`
	AvatarImage struct {
		Thumb   string `json:"thumb"`
		Display string `json:"display"`
	} `json:"avatar_image"`
	ChannelCount   int    `json:"channel_count"`
	FollowingCount int    `json:"following_count"`
	ProfileId      int    `json:"profile_id"`
	FollowerCount  int    `json:"follower_count"`
	Initials       string `json:"initials"`
	CanIndex       bool   `json:"can_index"`
	//Metadata       struct {
	//      Description *string `json:"description,omitempty"`
	//} `json:"metadata"`
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

type ApiChannelUser struct {
	CreatedAt   time.Time `json:"created_at"`
	Slug        string    `json:"slug"`
	Username    string    `json:"username"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	FullName    string    `json:"full_name"`
	Avatar      string    `json:"avatar"`
	AvatarImage *struct {
		Thumb   string `json:"thumb"`
		Display string `json:"display"`
	} `json:"avatar_image,omitempty"`
	ChannelCount   int    `json:"channel_count"`
	FollowingCount int    `json:"following_count"`
	ProfileID      int    `json:"profile_id"`
	FollowerCount  int    `json:"follower_count"`
	Initials       string `json:"initials"`
	CanIndex       bool   `json:"can_index"`
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
	BaseClass                   string `json:"base_class"` // 'User'
	Class                       string `json:"class"`      // 'User'
}

type ApiBlockImage struct {
	Filename    string    `json:"filename"`
	ContentType string    `json:"content_type"`
	UpdatedAt   time.Time `json:"updated_at"`
	Thumb       struct {
		Url string `json:"url"`
	} `json:"thumb"`
	Square struct {
		Url string `json:"url"`
	} `json:"square"`
	Display struct {
		Url string `json:"url"`
	} `json:"display"`
	Large struct {
		Url string `json:"url"`
	} `json:"large"`
	Original struct {
		Url             string `json:"url"`
		FileSize        int    `json:"file_size"`
		FileSizeDisplay string `json:"file_size_display"`
	} `json:"original"`
}

// not full block + connection to channel info
type ApiChannelBlock struct {
	Id              int           `json:"id"`
	Title           string        `json:"title,omitempty"`
	UpdatedAt       time.Time     `json:"updated_at"`
	CreatedAt       time.Time     `json:"created_at"`
	State           string        `json:"state"`
	CommentCount    int           `json:"comment_count"`
	GeneratedTitle  string        `json:"generated_title"`
	Content         string        `json:"content,omitempty"`
	Description     string        `json:"description,omitempty"`
	ContentHTML     template.HTML `json:"content_html,omitempty"`
	DescriptionHTML template.HTML `json:"description_html,omitempty"`
	Visibility      string        `json:"visibility"`
	Source          *struct {
		Url      string `json:"url"`
		Title    string `json:"title"`
		Provider struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"provider"`
	} `json:"source,omitempty"`
	//Embed               interface{}     `json:"embed"` // TODO
	//Attachment          interface{}     `json:"attachment"` // TODO
	//Metadata            interface{}     `json:"metadata"` // TODO
	Image               *ApiBlockImage  `json:"image,omitempty"`
	BaseClass           string          `json:"base_class"`
	Class               string          `json:"class"`
	User                *ApiChannelUser `json:"user"`
	Position            int             `json:"position"`
	Selected            bool            `json:"selected"`
	ConnectionId        int             `json:"connection_id"`
	ConnectedAt         time.Time       `json:"connected_at"`
	ConnectedByUserId   int             `json:"connected_by_user_id"`
	ConnectedByUsername string          `json:"connected_by_username"`
	ConnectedByUserSlug string          `json:"connected_by_user_slug"`
}

type ApiChannelResp struct {
	Id                int       `json:"id"`
	Title             string    `json:"title"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	AddedToAt         time.Time `json:"added_to_at"`
	Published         bool      `json:"published"`
	Open              bool      `json:"open"`
	Collaboration     bool      `json:"collaboration"`
	CollaboratorCount int       `json:"collaborator_count"`
	Slug              string    `json:"slug"`
	Length            int       `json:"length"`
	Kind              string    `json:"kind"`
	Status            string    `json:"status"`
	UserId            int       `json:"user_id"`
	Class             string    `json:"class"`      // Channel
	ClassName         string    `json:"class_name"` // Channel
	BaseClass         string    `json:"base_class"` // Channel
	Page              int       `json:"page"`
	Per               int       `json:"per"`
	CanIndex          bool      `json:"can_index"`
	ShareLink         string    `json:"share_link,omitempty"`
	FollowerCount     int       `json"follower_count"`
	Nsfw              bool      `json:"nsfw?,omitempty"`
	Metadata          *struct {
		Description string `json:"description,omitempty"`
	} `json:"metadata,omitempty"`
	Owner    *ApiChannelUser   `json:"owner,omitempty"`
	User     *ApiChannelUser   `json:"user"`
	Contents []ApiChannelBlock `json:"contents,omitempty"`
	//Collaborators *ApiChannelCollaborators `json:"collaborators"`
}
