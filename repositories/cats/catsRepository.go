package cats

import (
	"DoggosPkg/repositories/cats/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type CatRepository interface {
	GetCats(page int, limit int, breedID string) ([]model.CatDto, error)
}

type catRepository struct {
}

func NewCatRepository() CatRepository {
	return &catRepository{}
}

func (catRepository *catRepository) GetCats(page int, limit int, breedID string) ([]model.CatDto, error) {
	return fetchCats(page, limit, breedID)
}

func fetchCats(page int, limit int, breedID string) ([]model.CatDto, error) {
	var url string = fmt.Sprintf("https://api.thecatapi.com/v1/images/search?page=%d&limit=%d&mime_types=image/jpeg", page, limit)
	if breedID != "" {
		url += fmt.Sprintf("&breed_id=%s", breedID)
	}
	resp, err := http.Get(url)
	if err != nil {
		print(err)
		return nil, err
	}
	defer resp.Body.Close()
	var data []model.CatDto
	json.NewDecoder(resp.Body).Decode(&data)
	return data, nil
}
