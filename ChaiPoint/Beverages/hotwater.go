package beverages

import (
	"errors"

	ingredients "github.com/CoffeeMachine/ChaiPoint/Ingredients"
)

// HotWater beverage class
type HotWater struct {
	name  string
	water *ingredients.Ingredient
	Beverages
}

// NewHotWater return new instance of HotWater Beverage
func NewHotWater(beverageName string, waterIngredient *ingredients.Ingredient) *HotWater {
	return &HotWater{
		name:  beverageName,
		water: waterIngredient,
	}
}

// ProcessBeverageRequest always process one request to prepare the beverages
// check if all ingredients are available in required quantity
// if either of them is not available in required quantity,
// then return false along with error
func (w *HotWater) ProcessBeverageRequest() (bool, error) {

	if WaterAvailable := w.GetWater().ProcessIngredientRequest(IngredientWaterRequirementML); !WaterAvailable {
		return false, errors.New(w.GetBeverageName() + " cannot be prepared because : " + w.GetWater().GetName() + " is not available")
	}
	return true, errors.New(w.GetBeverageName() + " is prepared")
}

// GetBeverageName returns the beverage name
func (w *HotWater) GetBeverageName() string {
	return w.name
}

// SetBeverageName sets the beverage name
func (w *HotWater) SetBeverageName(name string) {
	w.name = name
}

// GetWater returns the Water Ingredient
func (w *HotWater) GetWater() *ingredients.Ingredient {
	return w.water
}
