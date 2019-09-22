// Sergey Olisov 2019 (c)
// Lesson 4
// Yet another one implementation of Double Linked list

package main

import (
	"fmt"
)

/*  List // тип контейнер
Len() // длинна списка
First() // первый Item
Last() // последний Item
PushFront(v interface{}) // добавить значение в начало
PushBack(v interface{}) // добавить значение в конец
Remove(i Item) // удалить элемент​

Item // элемент списка
Value() interface{} // возвращает значение
Nex() *Item // следующий Item
Prev() *Item // предыдущий
*/

type Item struct {
	DataContainer interface{}
	dl            *DList // Doual Linked list which this element belongs to
	nextItem      *Item  // Pointer to next Item
	prevItem      *Item  // Pointer to previous Item
}

type DList struct {
	Head   *Item // Head of the list
	Tail   *Item // Tail of the list
	Length int   // Lenght of the list
}

func (d *DList) Initialize() *DList {
	d.Head.nextItem = d.Tail
	d.Head.prevItem = nil
	d.Tail.nextItem = nil
	d.Tail.prevItem = d.Head
	d.Length = 0
	return d
}
func createDList() *DList {
	return new(DList).Initialize()
}

func (d *DList) Len() int {
	return d.Length
}

func (d *DList) First() *Item {
	return d.Head
}

func (d *DList) Last() *Item {
	return d.Tail
}

func (d *DList) PushFront(v interface{}) {

}
func (d *DList) PushBack(v interface{}) {

}

func (d *DList) Remove(i *Item) {

}

func (i *Item) Value() interface{} {
	return i.DataContainer
}
func (i *Item) Next() *Item {
	return i.nextItem
}

func (i *Item) Prev() *Item {
	return i.prevItem
}

func main() {
	fmt.Println("Linked List")
}
