package breeds

import (
	"DoggosPkg/repositories/breeds/model"
	"encoding/json"
	"net/http"
)

type BreedsRepository interface {
	GetBreeds() ([]model.BreedDto, error)
}

type breedsRepository struct {
}

func NewBreedsRepository() BreedsRepository {
	return &breedsRepository{}
}

func (breedsRepository *breedsRepository) GetBreeds() ([]model.BreedDto, error) {
	resp, err := http.Get("https://api.thedogapi.com/v1/breeds")
	if err != nil {
		print(err)
		return nil, err
	}
	defer resp.Body.Close()
	var data []model.BreedDto
	json.NewDecoder(resp.Body).Decode(&data)
	return data, nil
}
