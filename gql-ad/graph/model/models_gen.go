// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Ad struct {
	ID         int           `json:"id"`
	Dimensions *AdDimensions `json:"dimensions"`
	Texts      []string      `json:"texts"`
	MainColor  *string       `json:"mainColor"`
}

type AdDimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type NewAd struct {
	Dimensions *NewAdDimensions `json:"dimensions"`
	Texts      []string         `json:"texts"`
	MainColor  *string          `json:"mainColor"`
}

type NewAdDimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
