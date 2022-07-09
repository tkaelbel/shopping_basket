package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	"golang.org/x/exp/slices"
)

var INVENTORY []product = []product{{name: "A0001", price: 12.99}, {name: "A0002", price: 3.99}, {name: "A0003", price: 5.17}}

// Would also be possible to add an user struct
// And adding a basket to the user struct and so on...

type product struct {
	name  string
	price float32
}

type basket struct {
	items        []product
	buy1Get1Free product
	tenPercent   product
}

func (b *basket) total() float32 {
	var total float32 = 0.0

	var buy1Get1FreeProduct product

	for _, productItem := range b.items {

		if b.tenPercent.name == productItem.name {
			total = total + (productItem.price - (productItem.price * 0.1))
			continue
		}

		if b.buy1Get1Free.name == productItem.name {
			if buy1Get1FreeProduct.name == "" {
				buy1Get1FreeProduct = productItem
			} else {
				buy1Get1FreeProduct.name = ""
				continue
			}
		}

		total = total + productItem.price
	}

	return total
}

func new(buy1Get1Free product, tenPercent product) basket {
	return basket{items: []product{}, buy1Get1Free: buy1Get1Free, tenPercent: tenPercent}
}

func (b *basket) scan(name string) (product, error) {
	var p product
	idx := slices.IndexFunc(INVENTORY, func(p product) bool { return p.name == name })
	if idx == -1 {
		return p, errors.New("Not found")
	}
	p = INVENTORY[idx]
	b.items = append(b.items, p)
	return p, nil
}

func main() {
	fmt.Println("--- Welcome to the shop ---")
	fmt.Printf("--- The following items are available %v ---\n", INVENTORY)

	userBasket := new(product{name: "A0002"}, product{name: "A0001"})

	reader := bufio.NewReader(os.Stdin)

	printCommands()

	for true {

		error, input := getInput(reader)

		if error != nil {
			fmt.Println("--- Something wrong happened during input... Please try again ---")
			continue
		}

		if input == "total" {
			fmt.Printf("--- Current total price %v ---\n", userBasket.total())
			continue
		}

		if input == "h" {
			printCommands()
			continue
		}

		if input == "done" {
			fmt.Printf("--- Current total price %v ---\n", userBasket.total())
			fmt.Println("--- Exit. Bye ---")
			break
		}

		p, err := userBasket.scan(input)
		if err != nil {
			fmt.Printf("--- Could not find item %v ---\n", input)
			continue
		}

		fmt.Printf("--- Added product %v to basket ---\n", p.name)
		fmt.Printf("--- Current total price %v ---\n", userBasket.total())

	}

}

func getInput(reader *bufio.Reader) (error, string) {
	inputReturn, error := reader.ReadString('\n')
	var input string
	if runtime.GOOS == "windows" {
		input = strings.TrimRight(inputReturn, "\r\n")
	} else {
		input = strings.TrimRight(inputReturn, "\n")
	}
	return error, input
}

func printCommands() {
	fmt.Println("--- To add product type in the name -> A0001 ---")
	fmt.Println("--- To show the basket total price -> total ---")
	fmt.Println("--- To show commands -> h ---")
	fmt.Println("--- To finish buying -> done ---")
}
