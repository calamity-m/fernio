package food

type MongoFoodRepository struct {
}

func (repo *MongoFoodRepository) GetById(id string) (Food, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *MongoFoodRepository) GetByFilter(filter Filter) ([]Food, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *MongoFoodRepository) PutOne(data FoodRecord) (FoodRecord, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *MongoFoodRepository) PutMany(data []FoodRecord) ([]FoodRecord, error) {
	panic("not implemented") // TODO: Implement
}
func (repo *MongoFoodRepository) DeleteById(id string) error {
	panic("not implemented") // TODO: Implement
}

func (repo *MongoFoodRepository) DeleteByIds(id []string) error {
	panic("not implemented") // TODO: Implement
}

func (repo *MongoFoodRepository) DeleteByFilter(filter Filter) error {
	panic("not implemented") // TODO: Implement
}

type PostgresSqlFoodRepository struct {
}

func (repo *PostgresSqlFoodRepository) GetById(id string) (Food, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *PostgresSqlFoodRepository) GetByFilter(filter Filter) ([]Food, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *PostgresSqlFoodRepository) PutOne(data FoodRecord) (FoodRecord, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *PostgresSqlFoodRepository) PutMany(data []FoodRecord) ([]FoodRecord, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *PostgresSqlFoodRepository) DeleteById(id string) error {
	panic("not implemented") // TODO: Implement
}

func (repo *PostgresSqlFoodRepository) DeleteByIds(id []string) error {
	panic("not implemented") // TODO: Implement
}

func (repo *PostgresSqlFoodRepository) DeleteByFilter(filter Filter) error {
	panic("not implemented") // TODO: Implement
}

// Test for impl of interfaces
func TestIntfImpl(repo FoodRepository) {

}

func Go() {
	sql := &PostgresSqlFoodRepository{}
	mongo := &MongoFoodRepository{}

	TestIntfImpl(sql)
	TestIntfImpl(mongo)

}
