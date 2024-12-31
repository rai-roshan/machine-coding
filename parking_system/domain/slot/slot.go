package slot

type SlotType string

const (
	SMALL  SlotType = "SMALL"
	MEDIUM SlotType = "MEDIUM"
	LARGE  SlotType = "LARGE"
)

type Slot interface {
	GetSlotId() (uint32)
	IsOccupied() (bool, error)
	GetType() (SlotType, error)
	OccupySlot() (error)
	VaccantSlot() (error)
}

type slot struct {
	id       uint32
	slotType SlotType
	occupied bool
}

func (s *slot) GetSlotId() (uint32) {
	return s.id
}

func (s *slot) IsOccupied() (bool, error) {
	return s.occupied, nil
}

func (s *slot) GetType() (SlotType, error) {
	return s.slotType, nil
}

func (s *slot) OccupySlot() error {
	s.occupied = true
	return nil 
}

func (s *slot) VaccantSlot() error {
	s.occupied = false
	return nil
}

// builder
func NewSlot(id uint32, slotType SlotType) Slot {
	return &slot{
		id:       id,
		slotType: slotType,
	}
}

