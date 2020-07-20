package ingredients

import (
	"sync"
)

// Ingredient class
//
// This need not to be singleton class as there may be possibility that
// more than one ingredient slot holding same ingredient.
// This class's few methods need to be thread safe
type Ingredient struct {
	name             string
	quantity         float32
	minQuantityLevel float32
	isFinished       bool
	isBelowMinLevel  bool
	lock             *sync.Mutex
}

const (
	minQuantityPercentage = 0.2
)

// NewIngredient instantiates the ingredient - Ingredient
func NewIngredient(name string, quantity float32) *Ingredient {
	return &Ingredient{
		name:             name,
		quantity:         quantity,
		minQuantityLevel: quantity * minQuantityPercentage,
		isFinished:       false,
		isBelowMinLevel:  false,
		lock:             &sync.Mutex{},
	}
}

// GetName returns the name of the ingredient
func (i *Ingredient) GetName() string {
	return i.name
}

// SetName sets the name of the ingredient
func (i *Ingredient) SetName(name string) {
	i.name = name
}

// GetQuantity returns the name of the ingredient
func (i *Ingredient) GetQuantity() float32 {
	return i.quantity
}

// SetQuantity sets the name of the ingredient
func (i *Ingredient) SetQuantity(quantity float32) {
	i.quantity = quantity
}

// TopUpQuantity sets the name of the ingredient
func (i *Ingredient) TopUpQuantity(quantity float32) float32 {
	// function can be modified to see how much topup is possible
	// for current implementation, user is responsible for over filling
	i.quantity += quantity
	return quantity
}

// IsFinished returns if ingredient is finished.
func (i *Ingredient) IsFinished() bool {
	return i.isFinished
}

// IsBelowMinLevel returns if ingredient is finished.
func (i *Ingredient) IsBelowMinLevel() bool {
	return i.isBelowMinLevel
}


// ProcessIngredientRequest process the request for dispencing ingredient of mentioned volume.
// if not possible to serve, it returns false otherwise true
// function is thread safe to make sure requests get processed in sequential manner
// to preserve the quantity is right volume
func (i *Ingredient) ProcessIngredientRequest(quantity float32) bool {

	//fmt.Println(i.name, i.quantity, quantity)
	if i.IsFinished() {
		return false
	}

	// Note: Making it threaad-unsaafe. take care of it at application.
	//i.lock.Lock()
	//defer i.lock.Unlock()

	if (i.quantity - quantity) >= 0 {
		i.quantity -= quantity

		if i.quantity <= 0 {
			i.isFinished = true
		}

		if !i.isBelowMinLevel && i.quantity <= i.minQuantityLevel {
			i.isBelowMinLevel = true
		}
		return true
	}

	return false
}
