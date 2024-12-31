package parking_service

import (
	"fmt"
	"parking-system/domain/parking_directory"
	"parking-system/domain/plot"
	"parking-system/domain/slot"
	"parking-system/domain/vehicle"
)

type ParkingService interface {
	DisplayArea() error
	ParkVehicle(vehicle.Vehicle) error
	UnparkVehicle(vehicle.Vehicle) error
}

type parkingService struct {
	parkingPlot      plot.Plot
	parkingDirectory *parking_directory.ParkingDirectory
}

func (ps *parkingService) DisplayArea() error {
	rowArrangement, _ := ps.parkingPlot.Show()

	fmt.Println("parking areas as of now : ")
	for _, row := range rowArrangement {
		fmt.Println(row)
	}
	return nil
}

func (ps *parkingService) ParkVehicle(v vehicle.Vehicle) error {

	rowReq := uint32(1)
	colReq := uint32(0)

	var slotType slot.SlotType
	switch v.GetVehicleType() {
	case vehicle.BIKE:
		colReq = 1
		slotType = slot.SMALL
	case vehicle.CAR:
		colReq = 1
		slotType = slot.MEDIUM
	case vehicle.TRUCK:
		colReq = 5
		slotType = slot.LARGE
	}

	available, _ := ps.parkingPlot.CheckOpenArea(rowReq, colReq, slotType) // TODO : check slot type also
	if !available {
		return fmt.Errorf("no required space available")
	}

	slotIds, _ := ps.parkingPlot.OccupyArea(rowReq, colReq, slotType)

	ps.parkingDirectory.AddVehicleEntry(v.GetId(), parking_directory.ParkingAddress{
		PlotId: ps.parkingPlot.GetPlotId(),
		SlotId: slotIds,
	})

	return nil
}

func (ps *parkingService) UnparkVehicle(vh vehicle.Vehicle) error {
	vehicleId := vh.GetId()
	directory := ps.parkingDirectory.GetVehicleAddress(vehicleId)

	_ = ps.parkingPlot.FreeArea(directory.SlotId)
	ps.parkingDirectory.RemoveVehicleEntry(vehicleId)

	return nil
}

func NewParkingService(parkingPlot plot.Plot, parkingDirectory *parking_directory.ParkingDirectory) ParkingService {
	return &parkingService{
		parkingPlot:      parkingPlot,
		parkingDirectory: parkingDirectory,
	}
}
