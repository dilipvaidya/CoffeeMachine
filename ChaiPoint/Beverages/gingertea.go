package beverages

import (
	"errors"

	ingredients "github.com/CoffeeMachine/ChaiPoint/Ingredients"
)

// GingerTea beverage class
type GingerTea struct {
	name        string
	water       *ingredients.Ingredient
	milk        *ingredients.Ingredient
	teaSyrup    *ingredients.Ingredient
	gingerSyrup *ingredients.Ingredient
	sugarSyrup  *ingredients.Ingredient
}

// NewGingerTea return new instance of GingerTea Beverage
func NewGingerTea(beverageName string, waterIngredient *ingredients.Ingredient, milkIngredient *ingredients.Ingredient,
	teaSyrupIngredient *ingredients.Ingredient, gingerSyrupIngredient *ingredients.Ingredient,
	sugarSyrupIngredient *ingredients.Ingredient) *GingerTea {

	return &GingerTea{
		name:        beverageName,
		water:       waterIngredient,
		milk:        milkIngredient,
		teaSyrup:    teaSyrupIngredient,
		gingerSyrup: gingerSyrupIngredient,
		sugarSyrup:  sugarSyrupIngredient,
	}
}

// ProcessBeverageRequest always process one request to prepare the beverages
// check if all ingredients are available in required quantity
// if either of them is not available in required quantity,
// then return false along with error
func (g *GingerTea) ProcessBeverageRequest() (bool, error) {

	if !g.GetWater().ProcessIngredientRequest(IngredientWaterRequirementML) {
		return false, errors.New(g.GetBeverageName() + " cannot be prepared because : " + g.GetWater().GetName() + " is not available")
	}
	if !g.GetMilk().ProcessIngredientRequest(IngredientMilkRequirementML) {
		return false, errors.New(g.GetBeverageName() + " cannot be prepared because : " + g.GetMilk().GetName() + " is not available")
	}
	if !g.GetTeaSyrup().ProcessIngredientRequest(IngredientTeaRequirementML) {
		return false, errors.New(g.GetBeverageName() + " cannot be prepared because : " + g.GetTeaSyrup().GetName() + " is not available")
	}
	if !g.GetGingerSyrup().ProcessIngredientRequest(IngredientGingerSyrupRequirementML) {
		return false, errors.New(g.GetBeverageName() + " cannot be prepared because : " + g.GetGingerSyrup().GetName() + " is not available")
	}
	if !g.GetSugarSyrup().ProcessIngredientRequest(IngredientSugarSyrupRequirementML) {
		return false, errors.New(g.GetBeverageName() + " cannot be prepared because : " + g.GetSugarSyrup().GetName() + " is not available")
	}

	return true, errors.New(g.GetBeverageName() + " is prepared")
}

// GetBeverageName returns the beverage name
func (g *GingerTea) GetBeverageName() string {
	return g.name
}

// SetBeverageName sets the beverage name
func (g *GingerTea) SetBeverageName(name string) {
	g.name = name
}

// GetMilk returns the Milk Ingredient
func (g *GingerTea) GetMilk() *ingredients.Ingredient {
	return g.milk
}

// GetWater returns the Water Ingredient
func (g *GingerTea) GetWater() *ingredients.Ingredient {
	return g.water
}

// GetTeaSyrup returns TeaSyrup Ingredient
func (g *GingerTea) GetTeaSyrup() *ingredients.Ingredient {
	return g.teaSyrup
}

// GetGingerSyrup returns GingerSyrup Ingredient
func (g *GingerTea) GetGingerSyrup() *ingredients.Ingredient {
	return g.gingerSyrup
}

// GetSugarSyrup returns SugarSyrup Ingredient
func (g *GingerTea) GetSugarSyrup() *ingredients.Ingredient {
	return g.sugarSyrup
}
