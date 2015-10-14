package main

import (
	"fmt"
)

type node struct {
	next  *node
	value int
}
type LinkedList struct {
	head   *node
	tail   *node
	lenght int
}

func (ll *LinkedList) add_element(element int) {
	new_node := &node{next: nil, value: element}
	if ll.head == nil {
		ll.head = new_node
		ll.tail = ll.head
	} else {
		ll.tail.next = new_node
		ll.tail = ll.tail.next
	}
	fmt.Println(ll)
	ll.lenght++
}

func (ll *LinkedList) move_elements_front(start_range int, end_range int) {
	if start_range > ll.lenght || end_range > ll.lenght {
		return
	}
	var start_range_count int = 0
	for true {
		start_range_count++
		if start_range_count == start_range {
			break
		}
	}
}

func (ll *LinkedList) move_elements_back(start_range int, end_range int) {
	if start_range > ll.lenght || end_range > ll.lenght {
		return
	}
	var start_range_count int = 0
	for true {
		start_range_count++
		if start_range_count == start_range {
			break
		}
	}

}

func define_query(ll LinkedList, query int, start_range int, end_range int) {
	if query == 1 {
		ll.move_elements_front(start_range, end_range)
	} else if query == 2 {
		ll.move_elements_back(start_range, end_range)
	}
}

func (ll *LinkedList) print_ll() {
	current := *ll.head
	for true {
		fmt.Print(current.value, " ")
		if current.next == nil {
			break
		}
		current = *current.next
	}
}

func main() {
	ll := LinkedList{tail: nil, head: nil, lenght: 0}
	ll.add_element(1)
	ll.add_element(2)
	ll.add_element(3)
	ll.add_element(4)
	ll.add_element(5)
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
