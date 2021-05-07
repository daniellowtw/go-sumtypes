package testfixture

//go:generate go-sumtypes -name Vehicle -type Car,Ship

type Car struct{}
type Ship struct{}

func Functionality(vehicle Vehicle) int {
	var res int
	vehicle.Match(VehicleMatcher(
		func(Car *Car) {
			res = 10
		},
		func(Ship *Ship) {
			res = 100
		},
	))
	return res
}
