package model

type Pet struct {
	Breeds  []PetBreed `json:"breeds"`
	Height  int        `json:"height"`
	Id      string     `json:"id"`
	Url     string     `json:"url"`
	Width   string     `json:"width"`
	PetType string     `json:"pet_type"`
}
