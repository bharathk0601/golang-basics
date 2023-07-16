package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) Bill {
	b := Bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

func (b *Bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *Bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *Bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v $%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v $%v \n", "tip:", (*b).tip)
	fs += fmt.Sprintf("%-25v $%v \n", "total:", total)
	return fs
}

func createNewBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new Bill name: ", reader)
	b := newBill(name)

	return b
}

func (b *Bill) saveToFile() {
	data := []byte(b.format())

	err := ioutil.WriteFile("bill/"+b.name+".txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	fmt.Println("bill saved successfully.")
}

func promptOptions(b *Bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose options (a - add item, s - save Bill, t - add tip): ", reader)

	switch opt {
	case "a":
		item, _ := getInput("Enter the item name: ", reader)
		price, _ := getInput("Enter the "+item+" price: ", reader)
		parsedPrice, _ := strconv.ParseFloat(price, 64)

		b.addItem(item, parsedPrice)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter the tip: ", reader)
		parsedTip, _ := strconv.ParseFloat(tip, 64)

		b.updateTip(parsedTip)
		promptOptions(b)
	case "s":
		b.saveToFile()
		break
	default:
		fmt.Println("Invalid option")
		promptOptions(b)
	}
}
