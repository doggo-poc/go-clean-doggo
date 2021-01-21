package usecases

import (
	"github.com/Fcmam5/go-clean-doggo/repositories"
)

type DoggoUseCase struct {
	doggoRepo repositories.DoggoService
}

func (d DoggoUseCase) GetDoggos(page int, limit int) ([]*repositories.DoggoDto, error) {
	return d.doggoRepo.GetDoggos(page, limit)
}
