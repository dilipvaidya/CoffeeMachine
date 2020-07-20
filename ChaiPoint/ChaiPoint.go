package coffeemachine

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	outlets "github.com/CoffeeMachine/ChaiPoint/CoffeeMachineOutlets"
	ingredients "github.com/CoffeeMachine/ChaiPoint/Ingredients"
)

// constants for JSON string
const (
	jsonMachine                 = "machine"
	jsonOutlet                  = "outlets"
	jsonCntOutlets              = "count_n"
	jsonTotalIngredientQuantity = "total_items_quantity"
	jsonBeverages               = "beverages"
)

// once to keep tthis class singlton
var once sync.Once

// singleton instant
var coffeemachineInst *CoffeeMachine

// CoffeeMachine structure
type CoffeeMachine struct {
	ingredients  map[string]*ingredients.Ingredient // ingredient name to ingredient object
	numOfOutlets int8
	outlets      []*outlets.CoffeeMachineOutlets
	queue        chan int
	lock         *sync.Mutex

	//outletToBeverageMap map[int8]*outlets.CoffeeMachineOutlets
	//slots               []slots.Slots
	//beverages           []beverages.Beverages              // all possible beverages
	//slotToIngredientMap map[int8]*ingredients.Ingredient
}

// InitCoffeeMachine will initiate the coffee machine by number of outlets and then adding the ingredients
// This is the singlot class as we don't want even accidental more than one object creations per machine.
func InitCoffeeMachine(coffeeMachineInputJSON []byte) *CoffeeMachine {

	once.Do(func() { // this will execute only once so no mor instanciation will happen

		var result map[string]interface{}
		json.Unmarshal([]byte(coffeeMachineInputJSON), &result)
		numberOfOutlets := int8(result[jsonMachine].(map[string]interface{})[jsonOutlet].(map[string]interface{})[jsonCntOutlets].(float64))
		coffeemachineInst = &CoffeeMachine{
			ingredients:  make(map[string]*ingredients.Ingredient),
			numOfOutlets: numberOfOutlets,
			outlets:      make([]*outlets.CoffeeMachineOutlets, numberOfOutlets),
			queue:        make(chan int, numberOfOutlets),
			lock:         &sync.Mutex{},
		}

		totalItemsQuantity := result[jsonMachine].(map[string]interface{})[jsonTotalIngredientQuantity].(map[string]interface{})
		for ingredientName, initialQuantity := range totalItemsQuantity {
			coffeemachineInst.AddIngredients(ingredientName, float32(initialQuantity.(float64)))
		}

		for i := int8(0); i < numberOfOutlets; i++ {
			coffeemachineInst.outlets[i] = outlets.NewCoffeeMachineOutlets(i, true)
		}
	})
	return coffeemachineInst
}

// AddIngredients will add ingredients into the CoffeeMachine
func (c *CoffeeMachine) AddIngredients(name string, quantity float32) error {
	if _, ok := c.ingredients[name]; ok {
		return errors.New("Ingredient " + name + " already present. Not adding again")
	}

	c.ingredients[name] = ingredients.NewIngredient(name, quantity)
	return nil
}

// TopUpIngredient topup the existing ingredient by provided quantity.
// on success, it return how much quantity added - 0: none added; quantity == returnQuantity: all added
// -1 along with error to indecate ingredient not present already
func (c *CoffeeMachine) TopUpIngredient(name string, quantity float32) (float32, error) {
	if _, ok := c.ingredients[name]; !ok {
		return -1, errors.New("Ingredient " + name + " not present. No top up possible")
	}

	return c.ingredients[name].TopUpQuantity(quantity), nil
}

// ProcessBeverageRequest processes inout json with beverage requests
// processing happens in parallel for numOfOutlets beverages
func (c *CoffeeMachine) ProcessBeverageRequest(beverageJSON []byte) {

	var wg sync.WaitGroup // to wait to let all the go routines finish there work
	var result map[string]interface{}
	json.Unmarshal([]byte(beverageJSON), &result)

	// get the list of each beverage and process them in parallel
	beverages := result[jsonMachine].(map[string]interface{})[jsonBeverages].(map[string]interface{})
	for key, value := range beverages {

		c.queue <- 1
		wg.Add(1)
		go TransactForBeverage(c, key, value.(map[string]interface{}), c.queue, &wg)
	}

	wg.Wait()
}

// TransactForBeverage transacts for beveraged by allocating ingredients from the whole availability
func TransactForBeverage(c *CoffeeMachine, beverageName string, ingredientInterface map[string]interface{}, syncChan chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	c.lock.Lock()
	defer c.lock.Unlock()

	// rollback mechanism
	rollBackMap := make(map[string]float32)

	// for given beverage lookout id all ingredients are available in enough quantity
	for ingredient, quantity := range ingredientInterface {

		// if ingredient is not listed, return
		if _, ok := c.ingredients[ingredient]; !ok {
			fmt.Println(beverageName + " cannot be prepared because " + ingredient + " is not available")
			c.RollBack(rollBackMap)
			<-syncChan
			return
		}

		// if listed ingredient is not available in enough quantity, return
		success := c.ingredients[ingredient].ProcessIngredientRequest(float32(quantity.(float64)))
		if !success {
			fmt.Println(beverageName + " cannot be prepared because item " + ingredient + " is 0")
			c.RollBack(rollBackMap)
			<-syncChan
			return
		}

		rollBackMap[ingredient] = float32(quantity.(float64))
	}

	fmt.Println(beverageName + " is prepared")
	<-syncChan
	return
}

// RollBack function rolls back those transacttons which were
func (c *CoffeeMachine) RollBack(occupiedIngredients map[string]float32) {
	for ingredient, quantity := range occupiedIngredients {
		obj := c.ingredients[ingredient]
		newQuantity := obj.GetQuantity() + quantity
		obj.SetQuantity(newQuantity)
	}
}
