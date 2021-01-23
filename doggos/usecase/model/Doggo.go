package model

type Doggo struct {
	Breeds []Breed `json:"breeds"`
	Height int     `json:"height"`
	Id     string  `json:"id"`
	Url    string  `json:"url"`
	Width  string  `json:"width"`
}
