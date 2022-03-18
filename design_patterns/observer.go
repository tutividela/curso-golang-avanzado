package main

import "fmt"

type Topic interface {
	Registrar(observer Observer)
	Broadcast()
}

type Observer interface {
	GetId() string
	UpdateValue(string)
}

//Item -> No disponible
//Item (cuando este disponible) -> Notificar -> HAY ITEM

type Item struct {
	observers []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateValue() {
	i.available = true
	fmt.Printf("Item %s is now available\n",i.name)
	//Aqui hace de trigger
	i.Broadcast()
}

func (i *Item) Broadcast() {
	for _,o := range i.observers {
		o.UpdateValue(i.name)
	}
}

func (i *Item) Registrar(o Observer) {
	i.observers = append(i.observers, o)
}

type EmailClient struct {
	id string
}
func (e *EmailClient) GetId() string {
	return e.id
}
func (e*EmailClient) UpdateValue(item string) {
	fmt.Printf("Sendiga via Email - %s available from client %s\n",item,e.id)
}

func main() {
	nvidiaItem := NewItem("RTX 3080")
	firstObserver := &EmailClient{
		id: "12ab",
	}
	secondObserver := &EmailClient{
		id: "34cd",
	}
	nvidiaItem.Registrar(firstObserver)
	nvidiaItem.Registrar(secondObserver)
	nvidiaItem.UpdateValue()
}