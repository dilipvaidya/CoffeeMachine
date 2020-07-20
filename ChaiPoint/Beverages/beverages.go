package beverages

const (
	// IngredientWaterRequirementML is water requirement for receipe in ML
	IngredientWaterRequirementML = 50.0
	// IngredientMilkRequirementML milk requirement for receipe in ML
	IngredientMilkRequirementML = 10.0
	// IngredientTeaRequirementML teasyrup requirement for receipe in ML
	IngredientTeaRequirementML = 10.0
	// IngredientGingerSyrupRequirementML gingersyrup requirement for receipe in ML
	IngredientGingerSyrupRequirementML = 5.0
	// IngredientElaichiSyrupRequirementML elaichisyrup requirement for receipe in ML
	IngredientElaichiSyrupRequirementML = 5.0
	// IngredientCoffeeRequirementML coffee requirement for receipe in ML
	IngredientCoffeeRequirementML = 10.0
	// IngredientSugarSyrupRequirementML sugar syrup requirement for receipe in ML
	IngredientSugarSyrupRequirementML = 10.0
)

// Beverages interface as direction to implement all beverages
type Beverages interface {
	ProcessBeverageRequest() (bool, error)
	GetBeverageName() string
	SetBeverageName(name string)
}
