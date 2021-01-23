package repositories

type BreedDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	BreedGroup  string `json:"breed_group"`
	LifeSpan    string `json:"life_span"`
	Temperament string `json:"temperament"`
}

type DoggoDto struct {
	Breeds []BreedDto `json:"breeds"`
	Height int        `json:"height"`
	Id     string     `json:"id"`
	Url    string     `json:"url"`
	Width  string     `json:"width"`
}
