package main

// No ; or , between package
import (
	"fmt"
	"math"
	"math/rand"
)

// Go default execute the main function
func main() {
	// To get the doc -> godoc fmt Println
	fmt.Println("Welcome to Go!")

	// Call another func
	get_square_root()
	get_rand()
	print_add()
	print_concat()
	print_pointer()
	loop()
	map_test()
	print_sth()
	reverse_print()
	print_method()
	array_test()
	slice_test()
	operator_test()
	condition_test()
	switch_test()
	struct_test()
}

// Capitalize the first letter of the imported function
func get_square_root() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func get_rand() {
	// Get random number (Default seed)
	fmt.Println("Random number =", rand.Intn(100))
}

////////////////////////////////////////////////////////////
// Function pass by variables
////////////////////////////////////////////////////////////

// func add(x float64, y float64) float64 {
func add(x, y float64) float64 {
	return x + y
}

func print_add() {
	// var num1 float64 = 1.1
	// var num2 float64 = 2.1
	// var num1, num2 float64 = 1.1, 2.1
	// num1, num2 := 1.1, 2.1

	var num0 float32 = 1.1
	var num1 float64 = float64(num0)
	var num2 float64 = 2.1

	fmt.Println(add(num1, num2))
}

func concat(a, b string) (string, string) {
	return a, b
}

func concat_named(a, b string) (result string) {
	result = a + b
	return
}

func print_concat() {
	a, b := "Hi", "Word"
	fmt.Println(concat(a, b))
	fmt.Println(concat_named(a, b))
}

////////////////////////////////////////////////////////////
// Function pass by pointers
////////////////////////////////////////////////////////////
func print_pointer() {
	x := 15
	p := &x

	fmt.Println("Pointer =", p)
	fmt.Println("Pointer =", *p)

	*p = 5
	fmt.Println("Pointer =", *p)
	*p = *p * *p
	fmt.Println("Pointer =", *p)
}

////////////////////////////////////////////////////////////
// Loop
////////////////////////////////////////////////////////////
func loop() {
	// limited loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	for x := 5; x < 25; x += 3 {
		fmt.Println("Do stuff", x)
	}

	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
		// continue
		// break
	}

	// infinite loop
	/*
		for {
			fmt.Println("Do stuff.")
		}
	*/

	array1 := [3]int{1, 2, 3}
	for i, v := range array1 {
		fmt.Println(i, v)
	}
}

////////////////////////////////////////////////////////////
// Map = dict
////////////////////////////////////////////////////////////
func map_test() {
	// var grades map[string]int  // nil map
	grades := make(map[string]float32)
	// var grades = map[string]int{"Timmy": 42, "Jess": 92, "Sam": 71}

	// Assign values
	grades["Timmy"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 71
	fmt.Println(grades)

	// Get dict values
	tims_grade := grades["Timmy"]
	fmt.Println(tims_grade)

	// Get dict values and if exists
	val1, ok1 := grades["Timmy"]
	fmt.Println(val1, ok1)

	val1, ok1 = grades["Matt"]
	fmt.Println(val1, ok1)

	// Delete dict
	delete(grades, "Timmy")
	fmt.Println(grades)

	// print all
	for k, v := range grades {
		fmt.Println(k, ":", v)
	}
}

////////////////////////////////////////////////////////////////
// Defer -> Run until the last func executed
////////////////////////////////////////////////////////////////
func print_sth() {
	defer fmt.Println("Done!")                       // last one
	defer fmt.Println("Are we done yet?")            // this is the second one
	fmt.Println("Doing some stuff, who knows what?") // this will be printed out first
}

func reverse_print() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

////////////////////////////////////////////////////////////////
// Data Type: int, uint, float32, float64, string, bool
// Print Format %v = value, %T = type, %#v = Go-syntax format, %b = binary
////////////////////////////////////////////////////////////////
func print_method() {
	var s string = "This"
	var b1 bool = true

	fmt.Println(b1)
	fmt.Print(s, "123", "\n")
	fmt.Printf("value = %v, type = %T, Go-syntax format = %#v \n", s, s, s)
}

////////////////////////////////////////////////////////////////
// Array
////////////////////////////////////////////////////////////////
func array_test() {
	// empty array
	var array0 = [3]string{}
	fmt.Println(array0)

	// Lengths and values defined
	var array1 = [3]int{1, 2, 3}
	fmt.Println(array1)

	for x := range array1 {
		fmt.Println(x, array1[x])
	}

	// length inferred
	var array2 = [...]float64{1.2, 2.3, 3.4}
	array2[2] = 4.5
	for x := range array2 {
		fmt.Println(x, array2[x])
	}

	// partial initialize
	array3 := [5]int{1: 10, 2: 40}
	fmt.Println(array3)
	fmt.Println(len(array3))
}

////////////////////////////////////////////////////////////////
// Slice -> similar to array but the length can grow and shrink
////////////////////////////////////////////////////////////////
func slice_test() {
	// empty slice
	slice1 := []int{}
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // length of the slice
	fmt.Println(cap(slice1)) // capacity of the slice

	// slice from array
	array1 := [5]int{1, 2, 3, 4, 5}
	slice2 := array1[2:4]
	fmt.Println(slice2)
	fmt.Println(len(slice2)) // 2
	fmt.Println(cap(slice2)) // 2 -> 5 = 3

	// make to create slice
	slice3 := make([]int, 5, 10) // second = len, third = capacity
	fmt.Println(slice3, len(slice3), cap(slice3))

	slice4 := make([]int, 5) // without capacity
	fmt.Println(slice4, len(slice4), cap(slice4))

	// append element to slice
	slice5 := append(slice1, 12, 23)
	fmt.Println(slice5, len(slice5), cap(slice5))

	// append 2 slices
	slice6 := append(slice2, slice5...)
	fmt.Println(slice6, len(slice6), cap(slice6))

	// copy slice to smaller slice
	slice7 := make([]int, len(slice2))
	copy(slice7, slice2)
	fmt.Println(slice7, len(slice7), cap(slice7))
}

////////////////////////////////////////////////////////////////
// Operator
////////////////////////////////////////////////////////////////

func operator_test() {
	var x int = 5
	fmt.Printf("x = %v (%b)\n", x, x)

	// and & / or |
	x |= 3
	fmt.Printf("x |= %v (%b)\n", 3, 3)
	fmt.Printf("x = %v (%b)\n", x, x)

	// ^ XOR
	x = 5
	x ^= 3
	fmt.Printf("x ^= %v (%b)\n", 3, 3)
	fmt.Printf("x = %v (%b)\n", x, x)

	// >>
	x = 5
	x >>= 2
	fmt.Printf("x = %v (%b)\n", x, x)

	// <<
	x = 5
	x <<= 2
	fmt.Printf("x = %v (%b)\n", x, x)

	// logical and/or && ||
	fmt.Println(x > 3 && x < 15)
	fmt.Println(x > 3 || x < 25)
	fmt.Println(!(x > 3 || x < 25))
}

////////////////////////////////////////////////////////////////
// Condition
////////////////////////////////////////////////////////////////
func condition_test() {
	x, y := 20, 20
	// if statement
	if x > y {
		fmt.Println("x is greater than y")
	} else if x < y {
		fmt.Println("x is smaller than y")
	} else {
		fmt.Println("x is equal to y")
	}
}

func switch_test() {
	day := 7

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a weekday")
	}

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}

////////////////////////////////////////////////////////////////
// struct
////////////////////////////////////////////////////////////////

type Employee struct {
	name string
	age  int
	job  string
}

func struct_test() {
	var p1 = Employee{name: "Peter", age: 20, job: "Teacher"}
	fmt.Println(p1, p1.name, p1.age, p1.job)

	var p2 = Employee{"Matt", 30, "Engineer"}
	fmt.Println(p2, p2.name, p2.age, p2.job)

	var p3 Employee
	p3.name, p3.age, p3.job = "Cindy", 25, "Nurse"
	fmt.Println(p3, p3.name, p3.age, p3.job)
}
