package repositories

type BreedDto struct {
	id          string `json:"id" validate:"required"`
	name        string `json:"name" validate:"required"`
	breedGroup  string `json:"breed_group"`
	lifeSpan    string `json:"life_span" validate:"required"`
	temperament string `json:"temperament"`
}

type DoggoDto struct {
	breeds []BreedDto
	height int    `json:"height" validate:"required"`
	id     string `json:"id" validate:"required"`
	url    string `json:"url" validate:"required"`
	width  string `json:"width" validate:"required"`
}
