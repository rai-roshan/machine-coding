package main

import (
	"parking-system/domain/parking_directory"
	"parking-system/domain/plot"
	"parking-system/domain/vehicle"
	parking_service "parking-system/service"
)

func main() {
	plotArrangement := make(map[uint32][]string)
	plotArrangement[0] = []string{"L", "L", "L", "L"}
	plotArrangement[1] = []string{"L", "L", "L", "L"}
	plotArrangement[2] = []string{"L", "L", "L", "L"}
	newPlot := plot.NewPlot(1, 3, 4, plotArrangement)

	newParkingDirectory := parking_directory.NewParkingDirectory()

	newParkingService := parking_service.NewParkingService(newPlot, newParkingDirectory)
	newParkingService.DisplayArea()
	newParkingService.ParkVehicle(vehicle.NewVehicle(1, vehicle.CAR))
	newParkingService.DisplayArea()
	newParkingService.ParkVehicle(vehicle.NewVehicle(1, vehicle.CAR))
	newParkingService.DisplayArea()
	newParkingService.ParkVehicle(vehicle.NewVehicle(1, vehicle.CAR))
	newParkingService.DisplayArea()
	newParkingService.ParkVehicle(vehicle.NewVehicle(1, vehicle.BIKE))
	newParkingService.DisplayArea()
	newParkingService.ParkVehicle(vehicle.NewVehicle(1, vehicle.BIKE))
	newParkingService.DisplayArea()
}
