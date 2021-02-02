package model

type CatDto struct {
	Breeds []CatBreedDto `json:"breeds"`
	Height int           `json:"height"`
	Id     string        `json:"id"`
	Url    string        `json:"url"`
	Width  string        `json:"width"`
}
