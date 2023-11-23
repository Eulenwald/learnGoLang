package main

import "fmt"

// Definition eines Elements in der verketteten Liste
type Node struct {
	Data int
	Next *Node
}

// Definition der verketteten Liste
type LinkedList struct {
	Head *Node
}

// Funktion zum Hinzufügen eines Elements am Ende der Liste
func (ll *LinkedList) Add(data int) {
	newNode := &Node{Data: data, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

// Funktion zum Anzeigen der Liste
func (ll *LinkedList) Display() {
	current := ll.Head

	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}

	fmt.Println("nil")
}

func linkedList() {
	// Erstellen einer neuen verketteten Liste
	myList := LinkedList{}

	// Hinzufügen von Elementen zur Liste
	myList.Add(1)
	myList.Add(2)
	myList.Add(3)

	// Anzeigen der Liste
	myList.Display()
}
