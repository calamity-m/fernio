package food

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Define units for imperial/metric
type Mass float32
type Energy float32

const (
	Gram      Mass = 0
	Milligram Mass = 1
	Kilogram  Mass = 2
	Pound     Mass = 3
	Ounce     Mass = 4
)
const (
	Kilojoule   Energy = 0
	Kilocalorie Energy = 1
)

// Basic food, meant to capture a food item such as "Chicken Breast" or "Maccas Cheeseburger"
type Food struct {
	Id          uuid.UUID `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
}

// Represents a consumption/recording of Food
type FoodRecord struct {
	Id     uuid.UUID `json:"id,omitempty"`
	UserID uuid.UUID `json:"user_id,omitempty"`
	KJ     Energy    `json:"kj,omitempty"`
	Grams  Mass      `json:"grams,omitempty"`
	Date   time.Time `json:"date,omitempty"`
	Food   Food      `json:"food,omitempty"`
}

func (cf *FoodRecord) GetEnergy(unit Energy) Energy {
	if cf.KJ == 0 {
		return 0
	}

	switch unit {
	case Kilocalorie:
		return cf.KJ / 4.184
	default:
		return cf.KJ
	}
}

func (cf *FoodRecord) GetWeight(unit Mass) Mass {
	if cf.Grams == 0 {
		return 0
	}

	switch unit {
	case Kilogram:
		return cf.Grams / 1000
	case Milligram:
		return cf.Grams * 1000
	case Pound:
		return cf.Grams / 453.6
	case Ounce:
		return cf.Grams / 28.35
	default:
		return cf.Grams
	}
}

func Fuzz() {

	f1 := &FoodRecord{}

	f1.Food.Name = "h"

	cals := f1.GetEnergy(Kilocalorie)

	fmt.Println(cals)

	fmt.Println(f1.Id.Time())

}
