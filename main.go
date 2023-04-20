package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func CreateBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a New Bill Name :", reader)

	b := newBill(name)
	fmt.Println("create the bill -", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose Option(a - add item, s - save bill, t - add tip ):", reader)
	
	switch opt {
	case "a":
		name,_:=getInput("Enter Item Name:",reader)
		price,_:=getInput("Enter Price:",reader)

		p,err:=strconv.ParseFloat(price,64)

		if err !=nil {
			fmt.Println("The Price Must be number")
			promptOptions(b)
		}
		b.addItems(name,p)
		fmt.Println("Item Name:",name,"Item Price:",price)
		promptOptions(b)
	case "t":
		tip,_ := getInput("Enter tip Amount ($):",reader)

		t,err:=strconv.ParseFloat(tip,64)

		if err !=nil {
			fmt.Println("The tip Must be number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip Amount:",tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You choose the Save Bill:",b.name)
	default:
		fmt.Println("option Not Matched")
		promptOptions(b)
		
	}
}

func main() {
	myBill := CreateBill()
	promptOptions(myBill)

}
