package main

import "fmt"

type Vehicle struct {
	doors int
	color string
}

type Truck struct {
	vehicle   Vehicle
	fourWheel bool
}

type Sedan struct {
	vehicle Vehicle
	luxury  bool
}

func main() {
	vehicle := Vehicle{doors: 4, color: "Red"}
	vehicle2 := Vehicle{doors: 2, color: "Green"}
	vehicle3 := Vehicle{doors: 4, color: "Black"}

	truck := Truck{
		vehicle:   vehicle,
		fourWheel: true,
	}

	truck2 := Truck{
		vehicle:   vehicle2,
		fourWheel: false,
	}

	sedan := Sedan{
		vehicle: vehicle3,
		luxury:  true,
	}

	fmt.Println(truck)
	fmt.Println(truck2)
	fmt.Println(sedan)

	fmt.Println(truck.fourWheel)
	fmt.Println(sedan.luxury)
}
