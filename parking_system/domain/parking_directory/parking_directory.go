package parking_directory

type ParkingAddress struct {
	PlotId uint32
	SlotId []uint32
}

type ParkingDirectory struct {
	VehicleToParkedAddress map[uint32]ParkingAddress
}

func NewParkingDirectory() *ParkingDirectory {
	return &ParkingDirectory{
		VehicleToParkedAddress: make(map[uint32]ParkingAddress),
	}
}

func (pd *ParkingDirectory) GetVehicleAddress(vehicleId uint32) ParkingAddress {
	return pd.VehicleToParkedAddress[vehicleId]
}

func (pd *ParkingDirectory) RemoveVehicleEntry(vehicleId uint32) error {
	delete(pd.VehicleToParkedAddress, vehicleId)
	return nil
}

func (pd *ParkingDirectory) AddVehicleEntry(vehicleId uint32, pAddress ParkingAddress) error {
	pd.VehicleToParkedAddress[vehicleId] = pAddress
	return nil
}
