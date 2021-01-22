package repositories

type BreedDto struct {
	Id          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	BreedGroup  string `json:"breed_group"`
	LifeSpan    string `json:"life_span" validate:"required"`
	Temperament string `json:"temperament"`
}

type DoggoDto struct {
	Breeds []BreedDto
	Height int    `json:"height" validate:"required"`
	Id     string `json:"id" validate:"required"`
	Url    string `json:"url" validate:"required"`
	Width  string `json:"width" validate:"required"`
}
