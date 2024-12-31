package plot

import (
	"fmt"
	"parking-system/domain/slot"
)

type Plot interface {
	Show() ([]string, error)
	GetPlotId() uint32
	CheckOpenArea(rowSize uint32, columnSize uint32, slotType slot.SlotType) (bool, error)
	OccupyArea(rowSize uint32, columnSize uint32, slotType slot.SlotType) ([]uint32, error) // return list of occupied slot
	FreeArea(slotIds []uint32) error
}

type plot struct {
	id      uint32
	rows    uint32
	columns uint32
	area    map[uint32][]slot.Slot
}

func (p *plot) Show() ([]string, error) {
	var areaData []string
	for _, slots := range p.area {
		var rowData string
		for _, slot := range slots {
			isOccupied, _ := slot.IsOccupied()
			if isOccupied {
				rowData += "X "
			} else {
				rowData += "O "
			}
		}
		areaData = append(areaData, rowData)
	}

	return areaData, nil
}

func (p *plot) GetPlotId() uint32 {
	return p.id
}

func (p *plot) CheckOpenArea(rowSize uint32, columnSize uint32, slotType slot.SlotType) (bool, error) {

	for _, row := range p.area {
		for _, slot := range row {
			occupied, _ := slot.IsOccupied()
			currentSlotType, _ := slot.GetType()
			if columnSize > 0 && !occupied && slotType == currentSlotType {
				columnSize--
			}
			if columnSize == 0 {
				return true, nil
			}
		}
	}

	return false, nil
}

// return list of occupied slot
func (p *plot) OccupyArea(rowSize uint32, columnSize uint32, slotType slot.SlotType) ([]uint32, error) {
	for _, row := range p.area {

		slotNeeded := columnSize
		var occupiedIds []uint32

		for _, slot := range row {
			occupied, _ := slot.IsOccupied()
			currentSlotType, _ := slot.GetType()
			if !occupied && slotNeeded > 0 && slotType == currentSlotType {
				slotNeeded--
				slot.OccupySlot()
				occupiedIds = append(occupiedIds, slot.GetSlotId())
			}
		}
		if slotNeeded == 0 {
			return occupiedIds, nil
		}
	}

	return []uint32{}, fmt.Errorf("error while occupying area")
}

func (p *plot) FreeArea(slotIds []uint32) error {
	for _, row := range p.area {
		for _, slot := range row {

			slotId := slot.GetSlotId()

			for _, givenSlotId := range slotIds {
				if slotId == givenSlotId {
					slot.VaccantSlot()
				}
			}

		}
	}
	return nil
}

func NewPlot(id, row, column uint32, area map[uint32][]string) Plot {

	areaPloted := make(map[uint32][]slot.Slot)

	slotId := uint32(0)
	for rowId, slotType := range area {
		var rowSlots []slot.Slot
		for _, t := range slotType {
			switch t {
			case "S":
				rowSlots = append(rowSlots, slot.NewSlot(slotId, slot.SMALL))
			case "M":
				rowSlots = append(rowSlots, slot.NewSlot(slotId, slot.MEDIUM))
			case "L":
				rowSlots = append(rowSlots, slot.NewSlot(slotId, slot.LARGE))
			}
			slotId += 1
		}
		areaPloted[rowId] = rowSlots
	}

	return &plot{
		id:      id,
		rows:    row,
		columns: column,
		area:    areaPloted,
	}
}
