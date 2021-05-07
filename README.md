# Go sumtypes

This is an **experimental** package to explore generating code for sum types, aka disjointed union or variant types.

Typically, disjointed unions are implemented as subtyping in object-oriented programs. But that can lead to cumbersome
changes to the subtypes when a new method. This presents an issue when we want to add functionality to the main type
because we need to modify all the subtypes and add the functionality there. If the subtypes are relatively stable, it
would be more convenient if we only need to add a new function when adding functionality.

Sum types can help with this. See the following example:

```go
// We want to be able to express the following:
type Vehicle = Car | Ship | Bicycle | Plane // pseudo syntax

func NewFuntionality(Vehicle v) Result {
    var res Result
    switch x := v {
        case Car: res = f(x) // Access car's field
        case Ship: res = g(x) // Access ship's field
    }
    return res
}
```

This package generates the code to make the above pattern easy.

```go

// The following expresses Vehicle = Car | Ship
//go:generate go-sumtype -name Vehicles -type Car,Ship
type Car struct {}
type Ship struct {}

// Then you can write code like
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

```

## Alternatives

The trouble with the above is the silent failure when a new subtype is introduced. 
The switch case code will simply return the empty Result.

Furthermore, due to the duck typing, we cannot be sure if the Vehicle type is limited, or sealed, to the 4 sub types.
Granted, this can be solved by adding an unexported function to the Vehicle interface so we can explicitly tag the sub types.