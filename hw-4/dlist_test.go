package main
// Sergey Olisov (c) 2019
// Lesson 4 Double Linked List tests


import (
	"testing"
	"github.com/stretchr/testify/require"
)

/* The Following test cases will be covered:
- Create empty list - ensure that length is zero
- PushFront to empty list 
- PushFront to non empty list
- PushBack to empty list
- PushBack to non empty list
- PushFront & PushBack ??? (Looks like a dublicated test)
- Remove from the list of three elements an unknown element.
- Remove from the list of three elements an element in the middle.
- Remove from the list the last element.
- Remove from the list the first element.
- Remove form the empty list the last element.
- Check Item method - Prev()
- Check Item Method - Next()
- Check Item method - Value()
*/

func TestEmptyNewList(t *testing.T) {
	dl:=DList{}
	require.Equal(t,0,dl.Len(),"New empty List should has a length 0")
}

func TestPushFront(t *testing.T){
	dl:=DList{}
	dl.PushFront("front1") // Verify addition to front of empty list
	require.Equal(t,1,dl.Len(),"Lenght of list should be 1")
	require.Equal(t,"front1",dl.First().Value(),"The first value should be - front1")

	dl.PushFront("front2") // Verify addition to front of non empty list.
	
	require.Equal(t,2,dl.Len(),"Lenght of list should be 2")
	require.Equal(t,"front2",dl.First().Value(),"The first value should be - front2")
}

func TestPushBack(t *testing.T){
	dl:=DList{}
	dl.PushBack("back1") // verify addition to the end of empty list
	require.Equal(t,1,dl.Len(),"The Length of list should be 1")
	require.Equal(t,"back1",dl.Last().Value(),"The last value should be - back1")

	dl.PushBack("back2") // verify addition to the end of non empty list
	require.Equal(t,2,dl.Len(),"Lenght of list should be 2")
	require.Equal(t,"back2",dl.Last().Value(),"The latest value in the list should be - back2")
}
func TestCombinedPush(t *testing.T){
	dl:=DList{}
	dl.PushFront("front1") //Add from the beginin and from the end 
	dl.PushBack("back1")
	require.Equal(t,2,dl.Len(),"Lenght of List should be 2")
	require.Equal(t,"front1",dl.First().Value(),"Value of the fist element should be - front1")
	require.Equal(t,"back1",dl.Last().Value(),"Value of the last element should be - back1")
}
// Test Item methods ( Prev,Next,Value)
func TestItemBasicTest(t *testing.T){
	dl:=DList{}
	dl.PushFront("One")
	dl.PushFront("Two")
	require.Equal(t,dl.Last(),dl.First().Next(),"Check Item.Next(),should  point to the last item")
	require.Equal(t,dl.First(),dl.Last().Prev(),"Check that Item.Prev() should point to first item")
	require.Equal(t,"One",dl.Last().Value(),"Check the Item.Value() returns the correct value")
	require.Equal(t,"Two",dl.First().Value(),"Check the Item.Value() returns the correct value")
}
func TestRemove(t *testing.T){
	dl:=DList{}
	dl.PushFront("one")
	dl.PushFront("two")
	dl.PushFront("three")
	// Now we have a list like {three}-{two}-{one}
	// Let's try to remove the unknown element from the list
	unknownItem:=Item{DataContainer:"four"}
	require.Equal(t,false,dl.Remove(&unknownItem),"Removal of an unknown element shoud return false")
	require.Equal(t,3,dl.Len(),"Lenght of List should be 3")
	require.Equal(t,"one",dl.Last().Value(),"The last item must be one")
	require.Equal(t,"two",dl.First().nextItem.Value(),"The middle item must be two")
	require.Equal(t,"three",dl.First().Value(),"The first item must be three")
	// Let's try to remove the element in the middle of list
	//result:=dl.Remove(dl.First().nextItem)
	
	require.Equal(t,true,dl.Remove(dl.First().nextItem),"The removal should be true")
	require.Equal(t,2,dl.Len(),"Lenght of List should be 2")
	require.Equal(t,"one",dl.Last().Value())
	require.Equal(t,"three",dl.First().Value())

	// So we have {three}-{one}
	// Let's try to delete the last one

	require.Equal(t,true,dl.Remove(dl.Last()),"The removal should be true")
	require.Equal(t,1,dl.Len(),"Lenght of List should be 1")
	require.Equal(t,dl.First().Value(),dl.Last().Value())
	require.Equal(t,"three",dl.First().Value())

	// And we have -{three}- 
	// Let's del'em all

	require.Equal(t,true,dl.Remove(dl.First()),"Revomal should be ok - true")
	require.Equal(t,0,dl.Len(),"Lenght of List should be 0")
	require.Equal(t,(*Item)(nil),dl.Head,"First should be <nil>")
	require.Equal(t,(*Item)(nil),dl.Tail,"Last should be <nil>")

	// Just to try to remove something from the empty list
	require.Equal(t,false,dl.Remove(dl.Last()),"Removal from the empty list should return - false")
}
