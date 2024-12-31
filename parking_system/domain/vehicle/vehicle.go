package vehicle

type VehicleType string

const (
	BIKE  VehicleType = "BIKE"
	CAR   VehicleType = "CAR"
	TRUCK VehicleType = "TRUCK"
)

type Vehicle interface {
	GetId() uint32
	GetVehicleType() VehicleType
}

type vehicle struct {
	id          uint32
	vehicleType VehicleType
}

func (v *vehicle) GetId() uint32 {
	return v.id
}

func (v *vehicle) GetVehicleType() VehicleType {
	return v.vehicleType
}

func NewVehicle(id uint32, vType VehicleType) Vehicle {
	return &vehicle{
		id:          id,
		vehicleType: vType,
	}
}
