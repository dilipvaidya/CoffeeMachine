package outlets

// CoffeeMachineOutlets holds specific outvent information like id
type CoffeeMachineOutlets struct {
	id                  int8
	isInWorkingCodition bool
}

// NewCoffeeMachineOutlets return new instance for CoffeeMachineOutlets
func NewCoffeeMachineOutlets(id int8, workingCidition bool) *CoffeeMachineOutlets {
	return &CoffeeMachineOutlets{
		id:                  id,
		isInWorkingCodition: workingCidition,
	}
}

// GetID returns the slot Id of current slot
func (o *CoffeeMachineOutlets) GetID() int8 {
	return o.id
}

// SetID returns the slot Id of current slot
func (o *CoffeeMachineOutlets) SetID(id int8) {
	o.id = id
}

// IsInWorkingCodition returns the slot is occupied
func (o *CoffeeMachineOutlets) IsInWorkingCodition() bool {
	return o.isInWorkingCodition
}

// SetIsInWorkingCodition returns the slot Id of current slot
func (o *CoffeeMachineOutlets) SetIsInWorkingCodition(workingCidition bool) {
	o.isInWorkingCodition = workingCidition
}
