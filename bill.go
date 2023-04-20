package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 3.99, "cake": 5.99},
		tip:   0,
	}

	return b
}

// receive function to receive data from bill struct
// format the bill
func (b *bill) format()string{
	fs := "bill breakdown\n"

	var total float64 = 0
	//list of the items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v .....$%v\n", k+":", v)
		total += v
	}
	
	//add the tip
	fs += fmt.Sprintf("%-25v....$%v0.2f", "Tip :", b.tip)


	//total of the bill
	fs += fmt.Sprintf("\n%-25v0.2f....$%v0.2f", "Total :", total+b.tip)

	return fs
}
//update tip
func (b *bill) updateTip(tip float64)  {
	 b.tip = tip	
}
//add item to the bill
func (b *bill) addItems(name string , price float64	) {

	b.items[name] = price
	
}
//save the bill
func (b *bill) save() {
	data := []byte(b.format())
	
	err:=os.WriteFile("bills/"+b.name+".txt",data,0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("Bills Was Saved In the File")

}