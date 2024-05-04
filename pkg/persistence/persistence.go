package persistence

import (
	"fmt"
)

type Filter struct {
	Term  interface{}
	Value interface{}
}

// High level interface for a persistent store. Abstracted this way as while it's pointless, the intent is to play with
// and test subbing in/out different databases, allowing a service to not care about its storage mechanisms.
type StorageDriver interface {
	// Put(model interface{})
	// GetOne(filter interface{}) interface{}
	// GetAll(filter interface{}) interface{}
	// Delete(filter interface{}) interface{}
}

type Dri struct {
	StorageDriver
}

type PersistentRepository[T interface{}] interface {
	Put(model *T) error
	GetOne(filter Filter) (T, error)
	GetMany(filter Filter) ([]T, error)
	GetAll() ([]*T, error)
	Delete(filter Filter) error
}

type Driver interface {
	Info()
}

type Repository[T interface{}] interface {
	GetOne(filter Filter) (T, error)
}

type BaseRepository[T interface{}] struct {
	Host     string
	URI      string
	Username string
	Password string
	Driver   StorageDriver
	Repository[T]
}

type RepositoryImplementation[T interface{}] struct {
	driver Driver
	Repository[T]
}

/// impl test

type FakeD struct {
}

func (d *FakeD) Info() {
	panic("not implemented") // TODO: Implement
}

type ModelA struct {
	id string
}

type ModelB struct {
	name string
}

type ModelARepository struct{}

func (r *ModelARepository) GetOne(filter Filter) (ModelA, error) {
	fmt.Println("Model A GetOne")
	return ModelA{}, nil
}

type Stuff struct {
	Id string
}

type ModelBRepository[T ModelB] struct {
	BaseRepository[T]
	Stuff
}

func (r *ModelBRepository[T]) GetOne(filter Filter) (T, error) {
	fmt.Println("Model A GetOne")
	return T{name: "h"}, nil
}

func DoRepo[T interface{}](repo Repository[T]) {
	fmt.Println(repo)
	fmt.Println(repo.GetOne(Filter{}))
}

func Testing() {

	modelA := &ModelARepository{}
	// Could force people to use something like NewRepository under fernio/recorder/internal/storage? idk
	modelB := &ModelBRepository[ModelB]{BaseRepository: BaseRepository[ModelB]{Driver: FakeD{}}}

	DoRepo(modelA)
	DoRepo(modelB)

}
