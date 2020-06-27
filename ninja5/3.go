package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	/*t1 := truck{
		vehicle.doors = 4,
		vehicle.color = "red",
		fourWheel = true,
	}*/ //wrong

	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	/*s1 := sedan{
		vehicle.doors: 2,
		vehicle.color: "silver",
		luxury:        true,
	}*/
	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "silver",
		},
		luxury: true,
	}

	fmt.Println(t1, s1)
}
