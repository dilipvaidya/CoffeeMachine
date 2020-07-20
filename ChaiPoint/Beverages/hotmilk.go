package beverages

import (
	"errors"

	ingredients "github.com/CoffeeMachine/ChaiPoint/Ingredients"
)

// HotMilk beverage class
type HotMilk struct {
	name string
	milk *ingredients.Ingredient
	Beverages
}

// NewHotMilk return new instance of HotMilk Beverage
func NewHotMilk(beverageName string, milkIngredient *ingredients.Ingredient) *HotMilk {
	return &HotMilk{
		name: beverageName,
		milk: milkIngredient,
	}
}

// ProcessBeverageRequest always process one request to prepare the beverages
// check if all ingredients are available in required quantity
// if either of them is not available in required quantity,
// then return false along with error
func (m *HotMilk) ProcessBeverageRequest() (bool, error) {

	if milkAvailable := m.GetMilk().ProcessIngredientRequest(IngredientMilkRequirementML); !milkAvailable {
		return false, errors.New(m.GetBeverageName() + " cannot be prepared because : " + m.GetMilk().GetName() + " is not available")
	}
	return true, errors.New(m.GetBeverageName() + " is prepared")
}

// GetBeverageName returns the beverage name
func (m *HotMilk) GetBeverageName() string {
	return m.name
}

// SetBeverageName sets the beverage name
func (m *HotMilk) SetBeverageName(name string) {
	m.name = name
}

// GetMilk returns the Milk Ingredient
func (m *HotMilk) GetMilk() *ingredients.Ingredient {
	return m.milk
}
