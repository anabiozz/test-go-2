package builder

/*
	The Builder is a design pattern designed to provide a flexible solution to various object creation problems in object-oriented programming.
	The intent of the Builder design pattern is to separate the construction of a complex object from its representation
*/

// BuildProcess ...
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// ManufactoringDirector ...
type ManufactoringDirector struct {
	builder BuildProcess
}

// Construct ...
func (f *ManufactoringDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

// SetBuilder ...
func (f *ManufactoringDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// VehicleProduct ...
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

//************************* Car builder ***************************

// CarBuilder ...
type CarBuilder struct {
	v VehicleProduct
}

// SetWheels ...
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

// SetSeats ...
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

// SetStructure ...
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

// GetVehicle ...
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// **************** Bike Builder ***************************

// BikeBuilder ...
type BikeBuilder struct {
	v VehicleProduct
}

// SetWheels ...
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

// SetSeats ...
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

// SetStructure ...
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}

// GetVehicle ...
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

// **************** Bus Builder ***************************

// BusBuilder ...
type BusBuilder struct {
	v VehicleProduct
}

// SetWheels ...
func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 4
	return b
}

// SetSeats ...
func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 30
	return b
}

// SetStructure ...
func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}

// GetVehicle ...
func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}
