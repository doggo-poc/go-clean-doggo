package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DoggoRepository interface {
	GetDoggos(page int, limit int, breedID string) ([]DoggoDto, error)
}

type doggoRepository struct {
}

func NewDoggoRepository() DoggoRepository {
	return &doggoRepository{}
}

func (doggoRepository *doggoRepository) GetDoggos(page int, limit int, breedID string) ([]DoggoDto, error) {
	return fetchDoggos(page, limit, breedID)
}

func fetchDoggos(page int, limit int, breedID string) ([]DoggoDto, error) {
	var url string = fmt.Sprintf("https://api.thedogapi.com/v1/images/search?page=%d&limit=%d&mime_types=image/jpeg", page, limit)
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
	return data, nil
}
