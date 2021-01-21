package repositories

type DoggoService interface {
	GetDoggosByBreedId(page int, limit int, breedID string) ([]DoggoDto, error)
	GetDoggos(page int, limit int) ([]DoggoDto, error)
}
