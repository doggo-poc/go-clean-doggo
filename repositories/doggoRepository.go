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
	Client HttpClient
}

func NewDoggoRepository(c HttpClient) DoggoRepository {
	return &doggoRepository{
		Client: c,
	}
}

func NewUrl(page int, limit int, breedID string) string {
	var url string = fmt.Sprintf("https://api.thedogapi.com/v1/images/search?page=%d&limit=%d&mime_types=image/jpeg", page, limit)
	if breedID != "" {
		url += fmt.Sprintf("&breed_id=%s", breedID)
	}
	return url
}

func (doggoRepository *doggoRepository) GetDoggos(page int, limit int, breedID string) ([]DoggoDto, error) {
	url := NewUrl(page, limit, breedID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		print(err)
		return nil, err
	}

	resp, err := doggoRepository.Client.Execute(req)
	if err != nil {
		print(err)
		return nil, err
	}

	defer resp.Body.Close()
	var data []DoggoDto
	json.NewDecoder(resp.Body).Decode(&data)
	return data, nil
}
