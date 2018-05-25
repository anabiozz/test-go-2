package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufactoringComplex := ManufactoringDirector{}

	carBuilder := &CarBuilder{}
	manufactoringComplex.SetBuilder(carBuilder)
	manufactoringComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on car must be 5 and they were %d\n", car.Seats)
	}

	bikeBuilder := &BikeBuilder{}
	manufactoringComplex.SetBuilder(bikeBuilder)
	manufactoringComplex.Construct()

	motorbike := bikeBuilder.GetVehicle()
	motorbike.Seats = 1

	if motorbike.Wheels != 2 {
		t.Errorf("Wheels on a car must be 2 and they were %d\n", car.Wheels)
	}

	if motorbike.Structure != "Motorbike" {
		t.Errorf("Structure on a car must be 'Motorbike' and was %s\n", car.Structure)
	}

	busBuilder := &BusBuilder{}

	manufactoringComplex.SetBuilder(busBuilder)
	manufactoringComplex.Construct()

	bus := busBuilder.GetVehicle()

	if bus.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}

	if bus.Structure != "Bus" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if bus.Seats != 30 {
		t.Errorf("Seats on car must be 30 and they were %d\n", car.Seats)
	}

}
