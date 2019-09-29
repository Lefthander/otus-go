// Sergey Olisov 2019 (c)
// Lesson 4
// Yet another implementation of Double Linked list :)

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
Remove(i Item) // удалить элемент​ Hmmm logically remove should work with v:interface{}
 As we made add new item by value...
Item // элемент списка
Value() interface{} // возвращает значение
Nex() *Item // следующий Item
Prev() *Item // предыдущий
*/

// Item element of Double Linked List
type Item struct {
	DataContainer interface{} // Container to store the data
	nextItem      *Item       // Pointer to next Item
	prevItem      *Item       // Pointer to previous Item
}

// DList structure to store the components of the double linked list
type DList struct {
	Head   *Item // Head of the list
	Tail   *Item // Tail of the list
	Length int   // Lenght of the list
}

// Len returns lenght of the List
func (d *DList) Len() int {
	return d.Length
}

// First returns the first Item from the List
func (d *DList) First() *Item {
	return d.Head
}

// Last returns the last Item from the List
func (d *DList) Last() *Item {
	return d.Tail
}

// PushFront adds a new item at the begining of the List
func (d *DList) PushFront(v interface{}) {

	newItem := &Item{DataContainer: v}
	if d.Length == 0 {
		d.Head = newItem
		d.Tail = newItem
	} else {
		d.Head.prevItem = newItem
		newItem.nextItem = d.Head
		d.Head = newItem
	}
	d.Length++
}

//PushBack add a new Item to the end of the List
func (d *DList) PushBack(v interface{}) {
	newItem := &Item{DataContainer: v}
	// Check do we have an empty list... If so new item become a Head
	if d.Length == 0 {
		d.Head = newItem
		d.Tail = newItem
	} else {
		newItem.prevItem = d.Tail
		d.Tail.nextItem = newItem
		d.Tail = newItem
	}
	d.Length++
}

// Remove deletes the specific element from the List by the element's value.
func (d *DList) Remove(i *Item) bool {
	// Check is the List is empty
	if d.Length == 0 {
		return false
	}
	current := d.Head
	// Has a match
	if current == i {
		current = current.nextItem
		d.Length--
		if current != nil {
			d.Head = current
			d.Head.prevItem = nil
		} else {
			d.Tail = nil
			d.Head = nil
		}
		return true
	}
	for current.nextItem != nil {
		if current.nextItem == i {
			current.nextItem = current.nextItem.nextItem
			if current.nextItem == nil { //In case of last Item
				d.Tail = current
			} else {
				current.nextItem.prevItem = current
			}
			d.Length--
			return true
		}
		current = current.nextItem
	}
	return false
}

// Print does printing of all elements in the list,
// Misceleanous function does not requiered in the HW.
func (d *DList) Print() {
	i := d.Head
	for i != nil {
		fmt.Println(">", i.DataContainer)
		i = i.nextItem
	}
}

// Value returns the value of DataContainer from the specific Item
func (i *Item) Value() interface{} {
	return i.DataContainer
}

// Next returns the next Item
func (i *Item) Next() *Item {
	return i.nextItem
}

// Prev returns previous Item
func (i *Item) Prev() *Item {
	return i.prevItem
}

func main() {
	mylist := new(DList)
	mylist.PushFront("one")
	mylist.PushFront("two")
	mylist.Print()
	fmt.Println(mylist.Len())
	mylist.Remove(mylist.Last())
	mylist.Print()
	fmt.Println(mylist.Len())
	mylist.Remove(mylist.First())
	fmt.Println(mylist.Len())

}
