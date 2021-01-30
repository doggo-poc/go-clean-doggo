package breeds

import (
	"DoggosPkg/repositories"
	"DoggosPkg/repositories/breeds/model"
	"encoding/json"
	"net/http"
)

type BreedsRepository interface {
	GetBreeds() ([]model.BreedDto, error)
}

type breedsRepository struct {
	Client repositories.HttpClient
}

func NewBreedsRepository(client repositories.HttpClient) BreedsRepository {
	return &breedsRepository{
		Client: client,
	}
}

func (breedsRepository *breedsRepository) GetBreeds() ([]model.BreedDto, error) {
	r, error := http.NewRequest("GET", "https://api.thedogapi.com/v1/breeds", nil)
	if error != nil {
		print(error)
		return nil, error
	}
	resp, err := breedsRepository.Client.Execute(r)
	if err != nil {
		print(err)
		return nil, err
	}
	defer resp.Body.Close()
	var data []model.BreedDto
	json.NewDecoder(resp.Body).Decode(&data)
	return data, nil
}
