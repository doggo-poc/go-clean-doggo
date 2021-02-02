package model

type Cat struct {
	Breeds []CatBreed `json:"breeds"`
	Height int        `json:"height"`
	Id     string     `json:"id"`
	Url    string     `json:"url"`
	Width  string     `json:"width"`
}
