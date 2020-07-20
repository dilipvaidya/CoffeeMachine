package ingredientslots

// Slots holds id of the single coffee machine slot
type Slots struct {
	id          int8
	isAllocated bool
}

// NewIngredientSlots return new instance for CoffeeMachineOutlets
func NewIngredientSlots(id int8) *Slots {
	return &Slots{
		id:          id,
		isAllocated: false,
	}
}

// GetSlotID returns the slot Id of current slot
func (s *Slots) GetSlotID() int8 {
	return s.id
}

// SetSlotID returns the slot Id of current slot
func (s *Slots) SetSlotID(id int8) {
	s.id = id
}

// IsAllocated returns the slot is occupied
func (s *Slots) IsAllocated() bool {
	return s.isAllocated
}

// SetIsAllocated returns the slot Id of current slot
func (s *Slots) SetIsAllocated(allocated bool) {
	s.isAllocated = allocated
}
