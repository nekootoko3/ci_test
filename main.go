package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./calculator"
)

func main() {
	fmt.Print("This is easy calculator\n")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("please input int\n")
		ls := scanner.Scan()
		if !ls {
			return
		}
		left, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Print("Please input int\n")
		}

		fmt.Printf("Please input operator ( + or -)\n")
		os := scanner.Scan()
		if !os {
			return
		}
		operator := scanner.Text()

		fmt.Printf("Please input int\n")
		rs := scanner.Scan()
		if !rs {
			return
		}

		right, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Print("Please input int\n")
		}

		switch operator {
		case "+":
			sum := calculator.Add(left, right)
			fmt.Printf("sum is %d\n", sum)
		case "-":
			diff := calculator.Substract(left, right)
			fmt.Printf("diff is %d\n", diff)
		default:
			fmt.Printf("invalid operator. got=%s\n", operator)
		}
		fmt.Print("----- next calculation -----\n")
	}
}
