package model

type User struct {
	ID        int      `json:"id"`
	Gender    *Gender  `json:"gender"`
	BirthYear *int     `json:"birth_year"`
	Income    *int64   `json:"income"`
	GeoPosID  *int     `json:"geo_pos"`
	Interests []string `json:"interests"`
}
