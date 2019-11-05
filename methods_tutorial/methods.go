package main

import "fmt"

const u16bitMax float64 = 65535
const kmhMultiple float64 = 1.60934

type car struct {
	gasPedal      uint16 //from 0 to 65535 (means unsigned integer)
	brakePedal    uint16
	steeringWheel uint16
	topSpeedKmh   float64
}

func (c *car) kmh() float64 { // making function pointer-reciever of car struct, we can permanently modify the instance of that struct
	c.topSpeedKmh = 500
	return float64(c.gasPedal) * (c.topSpeedKmh / u16bitMax)
}

func (c car) mph() float64 { // value-reciever does not use '*' for memory access; creates a copy of the var we pass before operating; impermanent changes
	// if car.kmh() had not been a pointer-reciever method, car.topSpeedKmh would have defaulted to the value we assigned on init
	return float64(c.gasPedal) * (c.topSpeedKmh / u16bitMax / kmhMultiple)
}

func (c *car) newTopSpeed(newSpeed float64) { // '*' delimits pointer-reciever; we are modifying the memory of the variable we pass, so change is permanent
	// without asterisk, changes made to the variable remain in local scope
	c.topSpeedKmh = newSpeed
}

/*
We want to use value-recievers with small structs because we will not waste too much memory when the reciever copies the struct
We want pointer-recievers when we are working with large structs and do not have the capacity to copy the entire object in memory
before operating. It is a matter of efficiency.
*/

// here is a function (rather than method/reciever) that does the same thing as above methods
func newSpeed(c car, speed float64) car {
	c.topSpeedKmh = speed
	return c
}

func main() {
	car := car{gasPedal: 293,
		brakePedal:    0,
		steeringWheel: 657,
		topSpeedKmh:   466.3} // long way to write, it makes code more readable (tells us what each value corresponds to)
	fmt.Println(car.gasPedal)
	fmt.Println(car.kmh())
	fmt.Println(car.mph())
	car.newTopSpeed(500)
	fmt.Println(car.kmh())
	fmt.Println(car.mph())
	car = newSpeed(car, 600)
	fmt.Println(car.kmh())
	fmt.Println(car.mph())
}
