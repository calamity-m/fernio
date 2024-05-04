package persistence

import (
	"fmt"
)

// This might be better to use?
type Filter struct {
	Terms map[string]any
}

func FiltPlay() {
	filter := Filter{Terms: map[string]any{"a": "b"}}

	for term := range filter.Terms {
		fmt.Printf("Term: %v, Value: %v\n", term, filter.Terms[term])
	}
}

// High level interface for a persistent store. Abstracted this way as while it's pointless, the intent is to play with
// and test subbing in/out different databases, allowing a service to not care about its storage mechanisms.

type StorageHandler[T any] interface {
	// Handle fetching and gathering the connection
	Connect(uri string) error
	// Handle disconnection when done
	Disconnect() error

	// Shortcut function for commonly used get by id
	GetById(id string) (T, error)
	// Used for getting one entity
	GetOne(filter Filter) (T, error)
	// Used for creating or updating one entity
	PutOne(data T) (T, error)
	// Used for getting many entities at once
	GetMany(filter Filter) ([]T, error)
	// Used for creating or updating many entities at once
	PutMany(datas []T) ([]T, error)
	// Used for deleting one
	DeleteOne(filter Filter) error
	// Used for deleting many
	DeleteMany(filter Filter) error
}

type MongoHandler[T any] struct {
}

// Handle fetching and gathering the connection
func (h *MongoHandler[T]) Connect(uri string) error {
	fmt.Println("i connected")
	return nil
}

// Handle disconnection when done
func (h *MongoHandler[T]) Disconnect() error {
	panic("not implemented") // TODO: Implement
}

// Shortcut function for commonly used get by id
func (h *MongoHandler[T]) GetById(id string) (T, error) {
	fmt.Println("Getting something from mongo, id")
	var out T
	return out, nil
}

// Used for getting one entity
func (h *MongoHandler[T]) GetOne(filter Filter) (T, error) {
	panic("not implemented") // TODO: Implement
}

// Used for creating or updating one entity
func (h *MongoHandler[T]) PutOne(data T) (T, error) {
	panic("not implemented") // TODO: Implement
}

// Used for getting many entities at once
func (h *MongoHandler[T]) GetMany(filter Filter) ([]T, error) {
	panic("not implemented") // TODO: Implement
}

// Used for creating or updating many entities at once
func (h *MongoHandler[T]) PutMany(datas []T) ([]T, error) {
	panic("not implemented") // TODO: Implement
}

// Used for deleting one
func (h *MongoHandler[T]) DeleteOne(filter Filter) error {
	panic("not implemented") // TODO: Implement
}

// Used for deleting many
func (h *MongoHandler[T]) DeleteMany(filter Filter) error {
	panic("not implemented") // TODO: Implement
}

type Repository[T any] struct {
	Host     string
	URI      string
	Username string
	Password string
	Handler  StorageHandler[T]
}
