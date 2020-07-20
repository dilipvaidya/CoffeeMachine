package beverages

import (
	"errors"

	ingredients "github.com/CoffeeMachine/ChaiPoint/Ingredients"
)

// Coffee beverage class
type Coffee struct {
	name        string
	water       *ingredients.Ingredient
	milk        *ingredients.Ingredient
	coffeeSyrup *ingredients.Ingredient
	sugarSyrup  *ingredients.Ingredient
	Beverages
}

// NewCoffee return new instance of Coffee Beverage
func NewCoffee(beverageName string, waterIngredient *ingredients.Ingredient, milkIngredient *ingredients.Ingredient,
	coffeeSyrupIngredient *ingredients.Ingredient, sugarSyrupIngredient *ingredients.Ingredient) *Coffee {

	return &Coffee{
		name:        beverageName,
		water:       waterIngredient,
		milk:        milkIngredient,
		coffeeSyrup: coffeeSyrupIngredient,
		sugarSyrup:  sugarSyrupIngredient,
	}
}

// ProcessBeverageRequest always process one request to prepare the beverages
// check if all ingredients are available in required quantity
// if either of them is not available in required quantity,
// then return false along with error
func (c *Coffee) ProcessBeverageRequest() (bool, error) {

	if !c.GetWater().ProcessIngredientRequest(IngredientWaterRequirementML) {
		return false, errors.New(c.GetBeverageName() + " cannot be prepared because : " + c.GetWater().GetName() + " is not available")
	}
	if !c.GetMilk().ProcessIngredientRequest(IngredientMilkRequirementML) {
		return false, errors.New(c.GetBeverageName() + " cannot be prepared because : " + c.GetMilk().GetName() + " is not available")
	}
	if !c.GetCoffeeSyrup().ProcessIngredientRequest(IngredientTeaRequirementML) {
		return false, errors.New(c.GetBeverageName() + " cannot be prepared because : " + c.GetCoffeeSyrup().GetName() + " is not available")
	}
	if !c.GetSugarSyrup().ProcessIngredientRequest(IngredientSugarSyrupRequirementML) {
		return false, errors.New(c.GetBeverageName() + " cannot be prepared because : " + c.GetSugarSyrup().GetName() + " is not available")
	}

	return true, errors.New(c.GetBeverageName() + " is prepared")
}

// GetBeverageName returns the beverage name
func (c *Coffee) GetBeverageName() string {
	return c.name
}

// SetBeverageName sets the beverage name
func (c *Coffee) SetBeverageName(name string) {
	c.name = name
}

// GetMilk returns the Milk Ingredient
func (c *Coffee) GetMilk() *ingredients.Ingredient {
	return c.milk
}

// GetWater returns Water Ingredient
func (c *Coffee) GetWater() *ingredients.Ingredient {
	return c.water
}

// GetCoffeeSyrup returns CoffeeSyrup Ingredient
func (c *Coffee) GetCoffeeSyrup() *ingredients.Ingredient {
	return c.coffeeSyrup
}

// GetSugarSyrup returns SugarSyrup Ingredient
func (c *Coffee) GetSugarSyrup() *ingredients.Ingredient {
	return c.sugarSyrup
}
