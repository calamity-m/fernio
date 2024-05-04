package food

import (
	"fmt"

	"github.com/calamity-m/fernio/pkg/persistence"
	"github.com/google/uuid"
)

type FoodDao struct {
	Id   uuid.UUID
	Name string
}

type FoodRepository struct{}

func (fr FoodRepository) Put(model *FoodDao) error {
	fmt.Println("put")
	return nil
}

func (fr FoodRepository) GetOne(filter persistence.Filter) (FoodDao, error) {
	fmt.Println("get")
	return FoodDao{}, nil
}

func (fr FoodRepository) GetMany(filter persistence.Filter) ([]FoodDao, error) {
	fmt.Println("getm")
	return nil, nil
}

func (fr FoodRepository) GetAll() ([]*FoodDao, error) {
	fmt.Println("getal")
	return nil, nil
}

func (fr FoodRepository) Delete(filter persistence.Filter) error {
	fmt.Println("del")
	return nil
}

func uhhhhh(repo persistence.PersistentRepository[FoodDao]) {
	repo.GetAll()
}

func StorageTesting() {
	// filter := &FoodFilter{}

	foodrepo := &FoodRepository{}

	foodrepo.Put(&FoodDao{})

	uhhhhh(foodrepo)

}

type FoodRepo[T FoodDao] struct {
	persistence.BaseRepository[T]
}

func (r *FoodRepo[T]) GetOne(filter persistence.Filter) (T, error) {
	fmt.Println("fake")
	return T{Name: "FakeItem", Id: uuid.New()}, nil
}

func NewFoodRepo() *FoodRepo[FoodDao] {
	return &FoodRepo[FoodDao]{BaseRepository: persistence.BaseRepository[FoodDao]{Driver: persistence.FakeD{}}}
}
