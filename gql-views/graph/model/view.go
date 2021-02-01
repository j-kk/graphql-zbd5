package model

type View struct {
	ID        int    `json:"id"`
	AdID      int    `json:"ad"`
	UserID    int    `json:"user"`
	Timestamp string `json:"timestamp"`
}
