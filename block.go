package aredotna

// based on api responses

import (
	"html/template"
	"strconv"
	"time"
)

type BlockImage struct {
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

type BlockSource struct {
	Url      string `json:"url"`
	Title    string `json:"title"`
	Provider struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"provider"`
}

// not full block + connection to channel info
type Block struct {
	Id                  int             `json:"id"`
	Title               string          `json:"title,omitempty"`
	UpdatedAt           time.Time       `json:"updated_at"`
	CreatedAt           time.Time       `json:"created_at"`
	State               string          `json:"state"`
	CommentCount        int             `json:"comment_count"`
	GeneratedTitle      string          `json:"generated_title"`
	Content             string          `json:"content,omitempty"`
	Description         string          `json:"description,omitempty"`
	ContentHTML         template.HTML   `json:"content_html,omitempty"`
	DescriptionHTML     template.HTML   `json:"description_html,omitempty"`
	Visibility          string          `json:"visibility"`
	Source              *BlockSource    `json:"source,omitempty"`
	Image               *ApiBlockImage  `json:"image,omitempty"`
	BaseClass           string          `json:"base_class"`
	Class               string          `json:"class"`
	User                *ApiChannelUser `json:"user"`
	Slug                string          `json:"slug,omitempty"`
	Position            int             `json:"position"`
	Selected            bool            `json:"selected"`
	ConnectionId        int             `json:"connection_id"`
	ConnectedAt         time.Time       `json:"connected_at"`
	ConnectedByUserId   int             `json:"connected_by_user_id"`
	ConnectedByUsername string          `json:"connected_by_username"`
	ConnectedByUserSlug string          `json:"connected_by_user_slug"`
	//Embed               interface{}     `json:"embed"` // TODO
	//Attachment          interface{}     `json:"attachment"` // TODO
	//Metadata            interface{}     `json:"metadata"` // TODO
}

func (a *Arena) GetBlock(id int) (b *Block, err error) {
	path := "blocks/" + strconv.Itoa(id)
	err = a.get(path, Parameters{}, &b)
	if err != nil {
		return nil, err
	}
	return
}
