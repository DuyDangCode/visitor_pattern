package main

type IElement interface {
	accept(IVisitor) int
}

type Book struct {
	price      int
	isbnNumber string
}

func (b *Book) accept(v IVisitor) int {
	return v.visitForBook(b)
}

type Fruit struct {
	pricePerKg int
	weight     int
	name       string
}

func (f *Fruit) accept(v IVisitor) int {
	return v.visitForFruit(f)
}
