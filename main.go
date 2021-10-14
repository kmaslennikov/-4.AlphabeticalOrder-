package main

import "fmt"

func main() {
	fmt.Println("Введите первую строку")
	var first string
	fmt.Scan(&first)
	fmt.Println("Введите вторую строку")
	var second string
	fmt.Scan(&second)
	fmt.Println("Введите третью строку")
	var third string
	fmt.Scan(&third)

	if first < second {
		if first < third {
			fmt.Println(first)
			if second < third {
				fmt.Println(second)
				fmt.Println(third)
			} else {
				fmt.Println(third)
				fmt.Println(second)
			}
		} else {
			fmt.Println(third)
			fmt.Println(first)
			fmt.Println(second)
		}
	} else if second < third {
		fmt.Println(second)
		if first < third {
			fmt.Println(first)
			fmt.Println(third)
		} else {
			fmt.Println(third)
			fmt.Println(first)
		}
	} else {
		fmt.Println(third)
		fmt.Println(second)
		fmt.Println(first)
	}
}
