package main

import "fmt"

//creating empty structs to keep focus on interfaces

//Car car
type Car struct{}

//Boat boar
type Boat struct{}

//WashingMachine washingmachine
type WashingMachine struct{}

//Driller driller
type Driller struct{}

//Makeable makeable
//interface containing two methods.
//Each type which we want to implement these interfaces has to implement its functions.
type Makeable interface {
	Test() bool
	Personalize() bool
}

//Pluggable pluggable
//another interface
type Pluggable interface {
	CheckPlugConnection() bool
}

func main() {
	//creates  a simple slice of structs
	products := []Makeable{
		Car{},
		Boat{},
		WashingMachine{},
		Driller{},
	}

	for _, v := range products {
		Finish(v)
	}

	fmt.Println("completed!")
}

//Test simple implementation for Test in Makeable interface
func (c Car) Test() bool {
	fmt.Println("testing car")
	return true
}

//Personalize simple implementation for Test in Makeable interface
func (c Car) Personalize() bool {
	fmt.Println("personalizing car")
	return true
}

//Test simple implementation for Test in Makeable interface
func (c Boat) Test() bool {
	fmt.Println("testing boat")
	return true
}

//Personalize simple implementation for Test in Makeable interface
func (c Boat) Personalize() bool {
	fmt.Println("personalizing boat")
	return true
}

//Test simple implementation for Test in Makeable interface
func (c WashingMachine) Test() bool {
	fmt.Println("testing washing machine")
	return true
}

//Personalize simple implementation for Test in Makeable interface
func (c WashingMachine) Personalize() bool {
	fmt.Println("personalizing washing machine")
	return true
}

//CheckPlugConnection simple implementation for Test in Pluggable interface
func (c WashingMachine) CheckPlugConnection() bool {
	fmt.Println("checking plug connection in the washing machine")
	return true
}

//Test simple implementation for Test in Makeable interface
func (c Driller) Test() bool {
	fmt.Println("testing driller")
	return true
}

//Personalize simple implementation for Test in Makeable interface
func (c Driller) Personalize() bool {
	fmt.Println("personalizing driller")
	return true
}

//Finish to test all our makeable products
//it takes a variable type Makeable and returns it.
func Finish(product Makeable) Makeable {
	if ok := product.Test(); !ok {
		fmt.Println("test failed!")
	}

	if ok := product.Personalize(); !ok {
		fmt.Println("personalization failed!")
	}

	//checking if a product is of type!
	pluggable, is := product.(Pluggable)
	if is {
		if ok := pluggable.CheckPlugConnection(); !ok {
			fmt.Println("plug connection failed!")
		}
	}

	return product
}
