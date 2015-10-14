package main

import (
	"fmt"
)

type node struct {
	next       *node
	value      int
	start_node bool
}
type LinkedList struct {
	head   *node
	tail   *node
	lenght int
}

func (ll *LinkedList) add_element(element int) {

	empty_node := &node{next: nil, start_node: true}
	if ll.head == nil {
		ll.head = empty_node
		ll.tail = ll.head
	}

	new_node := &node{next: nil, value: element}
	ll.tail.next = new_node
	ll.tail = ll.tail.next
	fmt.Println(ll)
	ll.lenght++
}

func (ll *LinkedList) print_ll() {
	current := ll.head.next
	for true {
		fmt.Print(current.value, " ")
		if current.next == nil {
			break
		}
		current = current.next
	}
}

func (ll *LinkedList) move_elements(start_range int, end_range int, query int) {
	// range 1 to n special case
	if start_range > ll.lenght || end_range > ll.lenght {
		return
	}

	start_pnt := ll.head
	start_count := 0
	for ; start_count < start_range-1; start_count++ {
		fmt.Println("start_range:", start_pnt)
		start_pnt = start_pnt.next
	}

	end_pnt := start_pnt
	for ; start_count < end_range; start_count++ {
		fmt.Println("end_range:", start_count, end_range)
		end_pnt = end_pnt.next
	}

	if query == 2 {
		ll.tail.next = start_pnt.next
		start_pnt.next = end_pnt.next
		end_pnt.next = nil
	} else if query == 1 {

	}

	fmt.Println(ll.head)
}

func define_query(ll LinkedList, query int, start_range int, end_range int) {
	if query == 1 {
		//move_elements_front
		ll.move_elements(start_range, end_range, query)
	} else if query == 2 {
		//move_elements_back
		ll.move_elements(start_range, end_range, query)
	}
}

func main() {
	ll := LinkedList{tail: nil, head: nil, lenght: 0}
	ll.add_element(1)
	ll.add_element(2)
	ll.add_element(3)
	ll.add_element(4)
	ll.add_element(5)
	ll.add_element(6)
	ll.add_element(7)
	ll.add_element(8)
	fmt.Print("Before: ")
	ll.print_ll()
	fmt.Println()
	define_query(ll, 2, 1, 4)

	fmt.Print("Aflter: ")
	ll.print_ll()
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// fmt.Println(strings.Split(scanner.Text(), " "))
	// fmt.Println(scanner.Text())
	// }
	// var a, b int
	// fmt.Scanf("%v %v", &a, &b)

	// fmt.Println(a, b)
	// fmt.Scanf("%v", &a)
	// fmt.Println(a)
}
