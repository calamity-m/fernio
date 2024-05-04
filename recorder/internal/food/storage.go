package food

import (
	"github.com/calamity-m/fernio/pkg/persistence"
	"github.com/google/uuid"
)

type FoodDao struct {
	Id   uuid.UUID
	Name string
}

type FoodRecordDao struct {
	Id   string
	Name string
}

type FoodRepo struct {
	persistence.Repository[FoodDao]
}

type FoodRecordRepo struct {
	persistence.Repository[FoodRecordDao]
}

func NewFoodRecordRepo() *FoodRecordRepo {

	repo := &FoodRecordRepo{}

	v := &persistence.MongoHandler[FoodRecordDao]{}

	repo.Handler = v

	repo.Handler.Connect("")

	return repo
}

func NewFoodRepo() *FoodRepo {

	repo := &FoodRepo{}

	v := &persistence.MongoHandler[FoodDao]{}

	repo.Handler = v

	repo.Handler.Connect("")

	return repo
}
