package main

import "fmt"

type Color int

const (
	green Color = iota
	blue
)

type Size int

const (
	small Size = iota
	medium
	big
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return s.size == p.size
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

type Filter struct{}

func (f *Filter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, medium}
	house := Product{"House", blue, big}
	products := []Product{apple, tree, house}

	fmt.Printf("Green products: \n")
	greenSpec := ColorSpecification{green}
	filter := Filter{}

	for _, v := range filter.Filter(products, greenSpec) {
		fmt.Printf(" -%s is green!\n", v.name)
	}

	fmt.Printf("Big blue products: \n")
	largeSpec := SizeSpecification{big}
	blueSpec := ColorSpecification{blue}
	largeBlueSpec := AndSpecification{largeSpec, blueSpec}

	for _, v := range filter.Filter(products, largeBlueSpec) {
		fmt.Printf(" -%s is blue and big!\n", v.name)
	}
}
