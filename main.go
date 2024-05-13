package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	elements := make([]IElement, 0, 10)
	visitor := ShoppingCartVisitor{}

	for i := 0; i < 10; i++ {
		if rand.Int()%2 == 0 {
			newBook := Book{rand.Intn(100) + 100, "Book " + strconv.Itoa(i)}
			elements = append(elements, &newBook)
		} else {
			newFruit := Fruit{rand.Intn(150) + 50, (rand.Intn(1) + 2), "Fruit " + strconv.Itoa(i)}
			elements = append(elements, &newFruit)
		}
	}

	totalCost := 0
	fmt.Println("Cal total cost")
	for _, element := range elements {
		totalCost += element.accept(&visitor)
	}
	fmt.Println("Total Cost = ", totalCost)
}
