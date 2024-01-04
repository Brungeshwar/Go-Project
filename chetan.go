package main

import (
	"fmt"
	"os"
	"strconv"
)

// options
func (b Bill) options() {

	fmt.Printf("%-25v--Choose the following options--- \n", " ")
	fmt.Printf("2. type %q to %q \n", "b", "Add item")
	fmt.Printf("3. type %q to %q \n", "c", "Add Tip")
	fmt.Printf("4. type %q to %q \n", "d", "Save Bill")
	output := input()

	switch output {
	case "b":
		fmt.Println("Type item and price")
		fmt.Println("Item:")
		item := input()
		fmt.Println("Price:")
		val, _ := strconv.ParseFloat(input(), 64)
		b.addItem(item, val)

	case "c":
		fmt.Println("Type tip amount")
		val, _ := strconv.ParseFloat(input(), 64)
		b.addTip(val)
	case "d":
		str := b.printBill()
		data := []byte(str)
		err := os.WriteFile("Bills"+b.billName+".txt", data, 0644)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Saved successfully")
		}

	default:
		fmt.Println("Invalid input")

	}
	b.options()

}

type Bill struct {
	billName  string
	items     map[string]float64
	tipAmount float64
}

// Bill constructor
func createBill(name string) Bill {

	bill := Bill{
		billName:  name,
		items:     map[string]float64{},
		tipAmount: 0.0,
	}
	return bill

}

// to print the bill
func (b Bill) printBill() string {
	str := fmt.Sprintf("%-25v%v\n", "", b.billName)
	Total := 0.0

	for k, v := range b.items {
		str += fmt.Sprintf("%-25v : %f\n", k, v)
		Total += v

	}
	str += fmt.Sprintf("TipAmount: %v \n", b.tipAmount)
	str += fmt.Sprintf("TotalAmount: %v \n", Total)
	str += "----------"
	//fmt.Println(str)
	return str
	//fmt.Println(b)
}

// Bill structures methods
func (b Bill) addItem(item string, val float64) {
	b.items[item] = val

}
func (b Bill) addTip(t float64) {
	b.tipAmount = t
}
