package search

import (
	"fmt"
	"strings"
)

type Image struct {
	ID         int      `json:"id"`
	Slug       string   `json:"slug"`
	Width      int      `json:"width"`
	Height     int      `json:"height"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
	PromotedAt string   `json:"promoted_at"`
	URLs       []string `json:"urls"`
	Links      []string `json:"links"`
}

func (img *Image) Response() string {
	return fmt.Sprintf(
		"%d, %s, %d, %d, %s, %s, %s, %s, %s",
		img.ID,
		img.Slug,
		img.Width,
		img.Height,
		img.CreatedAt,
		img.UpdatedAt,
		img.PromotedAt,
		strings.Join(img.URLs, ","),
		strings.Join(img.Links, ","),
	)
}
