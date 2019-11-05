package main

import "fmt"

type car struct {
	gas_pedal      uint16 //from 0 to 65535 (means unsigned integer)
	brake_pedal    uint16
	steering_wheel uint16
	top_speed_kmh  float64
}

func main() {
	car := car{gas_pedal: 293,
		brake_pedal:    0,
		steering_wheel: 657,
		top_speed_kmh:  466.3} // long way to write, it makes code more readable (tells us what each value corresponds to)
	fmt.Println(car.gas_pedal)

}
