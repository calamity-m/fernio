package food

// This stuff should be taken and put into pkg

// Filter object
type Filter struct {
	Limit  uint
	Offset uint
	Terms  map[string]any
}

type FoodRecordRepository interface {
	GetById(id string) (FoodRecord, error)
	GetByFilter(filter Filter) ([]FoodRecord, error)
	PutOne(data FoodRecord) (FoodRecord, error)
	PutMany(data []FoodRecord) ([]FoodRecord, error)
	DeleteById(id string) error
	DeleteByIds(id []string) error
	DeleteByFilter(filter Filter) error
}

type FoodRepository interface {
	GetById(id string) (Food, error)
	GetByFilter(filter Filter) ([]Food, error)
	PutOne(data FoodRecord) (FoodRecord, error)
	PutMany(data []FoodRecord) ([]FoodRecord, error)
	DeleteById(id string) error
	DeleteByIds(id []string) error
	DeleteByFilter(filter Filter) error
}
