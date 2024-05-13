package main

import "fmt"

type IVisitor interface {
	visitForBook(*Book) int
	visitForFruit(*Fruit) int
}

type ShoppingCartVisitor struct{}

func (s *ShoppingCartVisitor) visitForBook(b *Book) int {
	cost := 0
	if b.price > 50 {
		cost = b.price - 5
	} else {
		cost = b.price
	}
	fmt.Println("Book ISBN::", b.isbnNumber, " cost = ", cost)
	return cost
}

func (s *ShoppingCartVisitor) visitForFruit(f *Fruit) int {
	cost := f.pricePerKg * f.weight
	fmt.Println(f.name, " cost = ", cost)
	return cost
}
