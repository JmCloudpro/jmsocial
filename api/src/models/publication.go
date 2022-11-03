package models

import "time"

// Publication representes  users Publications  structure
type Publication struct {
	ID          uint64    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Content     string    `json:"content,omitempty"`
	CreatorID   uint64    `json:"creatorID,omitempty"`
	CreatorNick uint64    `json:"creatorNick,omitempty"`
	Likes       uint64    `json:"jikes"`
	Createdin   time.Time `json:"createdin,omitempty"`
}
