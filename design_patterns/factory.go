package main

import "fmt"

type IProduct interface {
	getName() string
	setName(name string)
	getStock() int
	setStock(stock int)
}
type Computer struct {
	stock int
	name  string
}
func (c *Computer) getName() string {
	return c.name
}
func (c *Computer) setName(name string) {
	c.name = name
}
func (c *Computer) getStock() int {
	return c.stock
}
func (c *Computer) setStock(stock int) {
	c.stock = stock
}
type Desktop struct {
	Computer
}
type Laptop struct {
	Computer
}
func newDesktop(name string, stock int) (*Desktop, error) {
	return &Desktop{
		Computer{
			name:  name,
			stock: stock,
		},
	}, nil
}
func newLaptop(name string, stock int) (*Laptop, error) {
	return &Laptop{
		Computer{
			name:  name,
			stock: stock,
		},
	}, nil
}
func newProduct(name string, stock int) (IProduct, error) {
	if name == "laptop" {
		return newLaptop(name, stock)
	}
	if name == "desktop" {
		return newDesktop(name, stock)
	}
	return nil, fmt.Errorf("Invalid Computer Type")
}
func printNameAndStock(p IProduct) {
	fmt.Printf("Product , name:%s stock:%d\n",p.getName(),p.getStock())
}
func main(){
	laptop,err := newProduct("laptop",25)
	if err != nil {
		fmt.Print(err)
	}
	desktop,err := newProduct("desktop",40)
	if err != nil {
		fmt.Print(err)
	}
	printNameAndStock(laptop)
	printNameAndStock(desktop)
}
