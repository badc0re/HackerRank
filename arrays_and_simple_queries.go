package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"reflect"
	"strings"
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
	//fmt.Println(ll.lenght)
	ll.lenght += 1
}

func (ll *LinkedList) print_ll() {
	current := ll.head.next
	count := 0
	for true {
		fmt.Print(current.value, " ")
		if current.next == nil {
			break
		}
		current = current.next
		count++
	}
	fmt.Println()
}

func (ll *LinkedList) move_elements(start_range int, end_range int, query int) {
	// range 1 to n special case
	if start_range > ll.lenght || end_range > ll.lenght {
		return
	}
	if query == 1 && start_range == 1 {
		return
	}
	if query == 2 && end_range == ll.lenght {
		return
	}

	start_pnt := ll.head
	start_count := 0
	for ; start_count < start_range-1; start_count++ {
		// fmt.Println("start_range:", start_pnt)
		start_pnt = start_pnt.next
	}

	end_pnt := start_pnt
	for ; start_count < end_range; start_count++ {
		// fmt.Println("end_range:", start_count, end_range)
		end_pnt = end_pnt.next
	}

	if query == 1 {
		temp := ll.head.next
		ll.head.next = start_pnt.next
		start_pnt.next = end_pnt.next
		end_pnt.next = temp

		// fmt.Println("head:", ll.head.next, "start:", start_pnt.next, "end:", end_pnt.next, "tmp:", temp)
	} else if query == 2 {
		ll.tail.next = start_pnt.next
		start_pnt.next = end_pnt.next
		end_pnt.next = nil
	}

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
	/*
		ll := LinkedList{tail: nil, head: nil, lenght: 0}
		ll.add_element(1)
		ll.add_element(2)
		ll.add_element(3)
		ll.add_element(4)
		ll.add_element(5)
		fmt.Print("Before: ")
		ll.print_ll()
		fmt.Println()
		define_query(ll, 1, 1, 5)
		fmt.Print("Aflter: ")
		ll.print_ll()
	*/

	var list_lenght, ops int
	fmt.Scanf("%v %v", &list_lenght, &ops)
	scanner := bufio.NewScanner(os.Stdin)

	ll := LinkedList{tail: nil, head: nil, lenght: 0}
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), " ")
		if ll.lenght == 0 {
			for _, element := range tmp {
				temp_element, _ := strconv.Atoi(element)
				ll.add_element(temp_element)
			}
		} else {
			query, _ := strconv.Atoi(tmp[0])
			start_range, _ := strconv.Atoi(tmp[1])
			end_range, _ := strconv.Atoi(tmp[2])
			define_query(ll, query, start_range, end_range)
			fmt.Println(query, start_range, end_range)
			ll.print_ll()
		}
		// ll.print_ll()

		// kmt.Println(scanner.Text())
	}

	// fmt.Println(a, b)
	// fmt.Scanf("%v", &a)
	// fmt.Println(a)
}
