package cats

import (
	"DoggosPkg/repositories"
	"DoggosPkg/repositories/cats/model"

	"encoding/json"
	"fmt"
	"net/http"
)

type CatRepository interface {
	GetCats(page int, limit int, breedID string) ([]model.CatDto, error)
}

type catRepository struct {
	Client repositories.HttpClient
}

func NewCatRepository(client repositories.HttpClient) CatRepository {
	return &catRepository{
		Client: client,
	}
}

func NewUrl(page int, limit int, breedID string) string {
	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?page=%d&limit=%d&mime_types=image/jpeg", page, limit)
	if breedID != "" {
		url += fmt.Sprintf("&breed_id=%s", breedID)
	}
	return url
}

func (catRepository *catRepository) GetCats(page int, limit int, breedID string) ([]model.CatDto, error) {
	var url string = NewUrl(page, limit, breedID)
	r, error := http.NewRequest("GET", url, nil)
	if error != nil {
		print(error)
		return nil, error
	}
	resp, err := catRepository.Client.Execute(r)
	if err != nil {
		print(err)
		return nil, err
	}
	defer resp.Body.Close()
	var data []model.CatDto
	json.NewDecoder(resp.Body).Decode(&data)
	return data, nil
}
