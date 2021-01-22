package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DoggoRepository interface {
	GetDoggosByBreedId(page int, limit int, breedID string) ([]DoggoDto, error)
	GetDoggos(page int, limit int) ([]DoggoDto, error)
}

type doggoRepository struct {
}

func NewDoggoRepository() DoggoRepository {
	return &doggoRepository{}
}

func (doggoRepository *doggoRepository) GetDoggos(page int, limit int) ([]DoggoDto, error) {
	return fetchDoggos(page, limit, "")
}

func (doggoRepository *doggoRepository) GetDoggosByBreedId(page int, limit int, breedID string) ([]DoggoDto, error) {
	return fetchDoggos(page, limit, breedID)
}

func fetchDoggos(page int, limit int, breedID string) ([]DoggoDto, error) {
	var url string = fmt.Sprintf("https://api.thedogapi.com/v1/images/search?page=%d&limit=%d", page, limit)
	if breedID != "" {
		url += fmt.Sprintf("&breed_id=%s", breedID)
	}
	resp, err := http.Get(url)
	if err != nil {
		print(err)
		return nil, err
	}
	defer resp.Body.Close()
	var data []DoggoDto
	json.NewDecoder(resp.Body).Decode(&data)
	for r := range data {
		fmt.Printf("ID=%+v", r)
	}
	return data, nil
}
