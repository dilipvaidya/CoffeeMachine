package beverages

import (
	"errors"

	ingredients "github.com/CoffeeMachine/ChaiPoint/Ingredients"
)

// ElaichiTea beverage class
type ElaichiTea struct {
	name         string
	water        *ingredients.Ingredient
	milk         *ingredients.Ingredient
	teaSyrup     *ingredients.Ingredient
	elaichiSyrup *ingredients.Ingredient
	sugarSyrup   *ingredients.Ingredient
}

// NewElaichiTea return new instance of ElaichiTea Beverage
func NewElaichiTea(beverageName string, waterIngredient *ingredients.Ingredient, milkIngredient *ingredients.Ingredient,
	teaSyrupIngredient *ingredients.Ingredient, elaichiSyrupIngredient *ingredients.Ingredient,
	sugarSyrupIngredient *ingredients.Ingredient) *ElaichiTea {

	return &ElaichiTea{
		name:         beverageName,
		water:        waterIngredient,
		milk:         milkIngredient,
		teaSyrup:     teaSyrupIngredient,
		elaichiSyrup: elaichiSyrupIngredient,
		sugarSyrup:   sugarSyrupIngredient,
	}
}

// ProcessBeverageRequest always process one request to prepare the beverages
// check if all ingredients are available in required quantity
// if either of them is not available in required quantity,
// then return false along with error
func (e *ElaichiTea) ProcessBeverageRequest() (bool, error) {

	if !e.GetWater().ProcessIngredientRequest(IngredientWaterRequirementML) {
		return false, errors.New(e.GetBeverageName() + " cannot be prepared because : " + e.GetWater().GetName() + " is not available")
	}
	if !e.GetMilk().ProcessIngredientRequest(IngredientMilkRequirementML) {
		return false, errors.New(e.GetBeverageName() + " cannot be prepared because : " + e.GetMilk().GetName() + " is not available")
	}
	if !e.GetTeaSyrup().ProcessIngredientRequest(IngredientTeaRequirementML) {
		return false, errors.New(e.GetBeverageName() + " cannot be prepared because : " + e.GetTeaSyrup().GetName() + " is not available")
	}
	if !e.GetElaichiSyrup().ProcessIngredientRequest(IngredientElaichiSyrupRequirementML) {
		return false, errors.New(e.GetBeverageName() + " cannot be prepared because : " + e.GetElaichiSyrup().GetName() + " is not available")
	}
	if !e.GetSugarSyrup().ProcessIngredientRequest(IngredientSugarSyrupRequirementML) {
		return false, errors.New(e.GetBeverageName() + " cannot be prepared because : " + e.GetSugarSyrup().GetName() + " is not available")
	}

	return true, errors.New(e.GetBeverageName() + " is prepared")
}

// GetBeverageName returns the beverage name
func (e *ElaichiTea) GetBeverageName() string {
	return e.name
}

// SetBeverageName sets the beverage name
func (e *ElaichiTea) SetBeverageName(name string) {
	e.name = name
}

// GetMilk returns the Milk Ingredient
func (e *ElaichiTea) GetMilk() *ingredients.Ingredient {
	return e.milk
}

// GetWater returns the Water Ingredient
func (e *ElaichiTea) GetWater() *ingredients.Ingredient {
	return e.water
}

// GetTeaSyrup returns TeaSyrup Ingredient
func (e *ElaichiTea) GetTeaSyrup() *ingredients.Ingredient {
	return e.teaSyrup
}

// GetElaichiSyrup returns ElaichiSyrup Ingredient
func (e *ElaichiTea) GetElaichiSyrup() *ingredients.Ingredient {
	return e.elaichiSyrup
}

// GetSugarSyrup returns SugarSyrup Ingredient
func (e *ElaichiTea) GetSugarSyrup() *ingredients.Ingredient {
	return e.sugarSyrup
}
