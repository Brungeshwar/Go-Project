package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// to get input
func input() string {
	inp := bufio.NewReader(os.Stdin)
	output, _ := inp.ReadString('\n')
	output = strings.TrimSpace(output)
	return output
	//fmt.Println(output)
}

func main() {
	fmt.Println("Print the Bill Name")
	name := input()
	myBill := createBill(name)
	myBill.options()

	// mybill := createBill("First Bill")
	// mybill.addItem("chocolate", 24.5)
	// mybill.printBill()
	// fmt.Print("Task Completed!")
}
