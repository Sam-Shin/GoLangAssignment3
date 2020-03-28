package main

import "fmt"

//Car structs and map
type car struct {
	price int
	brand carBrand
}

type carBrand struct{
	manufacturer string
	model string
}

func (i *carBrand) getCarName() string {
	return i.manufacturer + " " + i.model
}


func printCarMap(){
	for i,j := range carMap{
		fmt.Println(i, j.brand.getCarName(), "\t\t\tPrice: $",j.price)
	}
}

func sellCar(){
	var n int
	fmt.Printf("Enter the key (1-%d) for which car you would like to sell", len(carMap))
	fmt.Scan(&n)
	fmt.Println("You have sold the", carMap[n].brand.manufacturer, carMap[n].brand.model, "for $", carMap[n].price)
	delete(carMap, n)
}

func buyCar() {
	newInt := len(carMap) + 1
	var p int
	fmt.Print("Enter the price of the car: ")
	fmt.Scan(&p)
	var manu string
	fmt.Print("Enter the manufacturer of the car: ")
	fmt.Scan(&manu)
	var mod string
	fmt.Print("Enter the model of the car: ")
	fmt.Scan(&mod)
	carMap[newInt] = car{price: p, brand: carBrand{manufacturer: manu, model: mod}}
}

func deleteCar(n int) {
	delete(carMap, n)
}

func addCar(priceCar int, carManufacturer, carModel string){
	newInt := len(carMap) + 1
	carMap[newInt] = car{price: priceCar, brand: carBrand{manufacturer: carManufacturer, model: carModel}}
}

var carMap map[int]car = map[int]car{1: car1, 2: car2, 3: car3, 4: car4, 5: car5}
var car1 car = car{price: 25000, brand: carBrand{manufacturer: "Toyota", model: "Camry"}}
var car2 car = car{price: 30000, brand: carBrand{manufacturer: "Honda", model: "Pilot"}}
var car3 car = car{price: 46000, brand: carBrand{manufacturer: "Mercedes", model: "CLA"}}
var car4 car = car{price: 28000, brand: carBrand{manufacturer: "Dodge", model: "Challenger"}}
var car5 car = car{price: 35000, brand: carBrand{manufacturer: "Subaru", model: "WRX"}}


//Substring Section for part 2
var phrase string = "Hello, World"

func sliceSub(start , end int) []rune{
	var tempSubstring []rune = []rune(phrase)
	var tempSlice []rune = tempSubstring[2:9]
	var newSubstring []rune = tempSlice[start:end]
	return newSubstring
}

//Main
func main() {
	//Part 1
	//loop breakpoint, giving user to choose what happens
loop:
	for {
		fmt.Printf("Here are your %d cars:", len(carMap))
		fmt.Println()
		printCarMap()
		fmt.Println("Enter '1' to sell a car \nEnter '2' to add a car\nEnter '3' to buy a car" +
			"\nEnter '4' to delete a car\nEnter '5' to quit")
		var x int
		fmt.Scanf("%d", &x)
		switch x {
		case 1:
			sellCar()
			printCarMap()
			break
		case 2:
			var manufact, model string
			var price int
			fmt.Println("Enter the manufacturer for the car you would like to add: ")
			fmt.Scan(&manufact)
			fmt.Println("Enter the model: ")
			fmt.Scan(&model)
			fmt.Println("Enter the price: ")
			fmt.Scanf("%d", &price)
			addCar(price, manufact, model)
			printCarMap()
			break
		case 3:
			buyCar()
			break
		case 4:
			var y int
			printCarMap()
			fmt.Printf("Enter the key (1-%d) for which car you would like to delete", len(carMap))
			fmt.Scan(&y)
			deleteCar(y)
			printCarMap()
			break
		case 5:
			fmt.Println("Good bye")
			break loop
		default:
			fmt.Println("invalid input, try again")
		}
	}
	//part 2
	var substring []rune = sliceSub(1,4)
	fmt.Println(string(substring))
}
//not sure on how to re-initialize keys to be in the range of the amount of cars in map
//i.e. there would be 5 values in a map but the range would be 2-6 instead of 1-5

//can use addCar(int, string, string) to hardcode a new car in carMap()
//can also use deleteCar(int) to hardcode deletion of a car