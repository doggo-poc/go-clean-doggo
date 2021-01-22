package usecases

import (
	"DoggosPkg/repositories"
)

type doggoUseCase struct {
	doggoRepo repositories.DoggoRepository
}

func NewDoggoUseCase(repo repositories.DoggoRepository) domain.ArticleUsecase {
	return &doggoUseCase{
		doggoRepo: repo,
	}
}

func (d *doggoUseCase) GetDoggos(page int, limit int) ([]repositories.DoggoDto, error) {
	return d.doggoRepo.GetDoggos(page, limit)
}
