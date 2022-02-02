package main

import "fmt"

// class definiton
type car struct {
	gas_pedal      uint16  //min: 0,      max: 65535    16bit
	brake_pedal    uint16  //min: 0,      max: 65535
	steering_wheel int16   //min: -32768  max: 32768
	top_speed_kmh  float64 //what's our top speed?
}

// constant variables
const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

// class method
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / kmh_multiple / usixteenbitmax)
}

// class pointer receivers
func (c car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func (c *car) new_top_speed_pointer(newspeed float64) {
	c.top_speed_kmh = newspeed
}

// Pass the class as agrument
func newer_top_speed(c car, speed float64) car {
	c.top_speed_kmh = speed
	return c
}

// main function
func main() {
	a_car := car{
		gas_pedal:      16535,
		brake_pedal:    0,
		steering_wheel: 12562,
		top_speed_kmh:  225.0,
	}
	// a_car := car{22314, 0, 12562, 225.0}
	fmt.Println("gas_pedal:", a_car.gas_pedal, "brake_pedal:", a_car.brake_pedal, "steering_wheel:", a_car.steering_wheel, "top_speed_kmh:", a_car.top_speed_kmh)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

	// update speed
	a_car.new_top_speed(500)
	fmt.Println("gas_pedal:", a_car.gas_pedal, "brake_pedal:", a_car.brake_pedal, "steering_wheel:", a_car.steering_wheel, "top_speed_kmh:", a_car.top_speed_kmh)

	// update speed by pointer
	a_car.new_top_speed_pointer(500)
	fmt.Println("gas_pedal:", a_car.gas_pedal, "brake_pedal:", a_car.brake_pedal, "steering_wheel:", a_car.steering_wheel, "top_speed_kmh:", a_car.top_speed_kmh)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

	// update speed by func
	a_car = newer_top_speed(a_car, 200)
	fmt.Println("gas_pedal:", a_car.gas_pedal, "brake_pedal:", a_car.brake_pedal, "steering_wheel:", a_car.steering_wheel, "top_speed_kmh:", a_car.top_speed_kmh)
}
