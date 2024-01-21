package aredotna

import (
	"time"
)

type Image struct {
	Url          string `json:"url"`
	FileSize     int    `json:"file_size,omitempty"`
	FileSizeDisp string `json:"file_size_display,omitempty"`
}

type Block struct {
	Id              int       `json:"id"`
	Title           string    `json:"title,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	State           string    `json:"state"`
	CommentCount    int       `json:"comment_count"`
	GeneratedTitle  string    `json:"generated_title"`
	Class           string    `json:"class"`
	BaseClass       string    `json:"base_class"`
	Content         string    `json:"content,omitempty"`
	ContentHtml     string    `json:"content_html,omitempty"`
	Description     string    `json:"description,omitempty"`
	DescriptionHtml string    `json:"description_html,omitempty"`
	Source          *struct {
		Url      string `json:"url"`
		Title    string `json:"title"`
		Provider struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"provider"`
	} `json:"source,omitempty"`
	Image *struct {
		Filename    string    `json:"filename"`
		ContentType string    `json:"content_type"`
		UpdatedAt   time.Time `json:"updated_at"`
		Thumb       Image     `json:"thumb"`
		Square      Image     `json:"square"`
		Display     Image     `json:"display"`
		Large       Image     `json:"large"`
		Original    Image     `json:"original"`
	} `json:"image,omitempty"`
	User struct {
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
		Metadata       struct {
			Description interface{} `json:"description"`
		} `json:"metadata"`
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
	} `json:"user"`
	Connections []struct {
		Id            int       `json:"id"`
		Title         string    `json:"title"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		AddedToAt     time.Time `json:"added_to_at"`
		Published     bool      `json:"published"`
		Open          bool      `json:"open"`
		Collaboration bool      `json:"collaboration"`
		Slug          string    `json:"slug"`
		Length        int       `json:"length"`
		Kind          string    `json:"kind"`
		Status        string    `json:"status"`
		UserId        int       `json:"user_id"`
		Metadata      struct {
			Description string `json:"description"`
		} `json:"metadata"`
		ShareLink interface{} `json:"share_link"`
		BaseClass string      `json:"base_class"`
	} `json:"connections"`
}

type Channel struct {
	Id                string    `json:"id"`
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
	Manifest          map[string]string
	Contents          []Block `json:"contents"`
	//Items []string
}

type User struct {
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
	//	Description any `json:"description"`
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

type Group struct {
	Class       string    `json:"class"`
	BaseClass   string    `json:"base_class"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	User        struct {
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
		//	Description any `json:"description"`
		//} `json:"metadata"`
		IsPremium                   bool `json:"is_premium"`
		IsLifetimePremium           bool `json:"is_lifetime_premium"`
		IsSupporter                 bool `json:"is_supporter"`
		IsExceedingConnectionsLimit bool `json:"is_exceeding_connections_limit"`
		IsConfirmed                 bool `json:"is_confirmed"`
		IsPendingReconfirmation     bool `json:"is_pending_reconfirmation"`
		IsPendingConfirmation       bool `json:"is_pending_confirmation"`
		//Badge                       any    `json:"badge"`
		Id        int    `json:"id"`
		BaseClass string `json:"base_class"`
		Class     string `json:"class"`
	} `json:"user"`
	Users           []any  `json:"users"`
	MemberIds       []int  `json:"member_ids"`
	AccessibleByIds []int  `json:"accessible_by_ids"`
	Published       bool   `json:"published"`
	Title           string `json:"title"`
	Id              int    `json:"id"`
}
