package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func find_princess(up_or_down int, princess int, m int) {
	for i := 0; i < (m-1)/2; i++ {
		if up_or_down == 0 {
			fmt.Println("UP")
		} else {
			fmt.Println("DOWN")
		}

		if princess > 0 {
			fmt.Println("RIGHT")
		} else {
			fmt.Println("LEFT")
		}
	}
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	m := 0
	reader := bufio.NewReader(os.Stdin)
	fmt.Scanf("%v", &m)

	princess := 0
	up_or_down := 0
	for i := 0; i < m; i++ {
		input_text, _ := reader.ReadString('\n')
		char_idx := strings.Index(input_text, "p")
		if i == 0 && char_idx != -1 {
            		up_or_down = 0
			princess = char_idx
		} else if i > 0 && char_idx != -1 {
			up_or_down = m
			princess = char_idx
		}
	}
	find_princess(up_or_down, princess, m)
}
