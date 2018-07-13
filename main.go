package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./calculator"
)

const PROMPT = ">> "

func main() {
	fmt.Print("This is easy calculator\n")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("----- Let's calculate -----\n")

		fmt.Print("Please input int\n")
        fmt.Printf(PROMPT)
		ls := scanner.Scan()
		if !ls {
			return
		}
		left, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Print("Please input int\n")
		}

		fmt.Printf("Please input operator ( + or -)\n")
        fmt.Printf(PROMPT)
		os := scanner.Scan()
		if !os {
			return
		}
		operator := scanner.Text()

		fmt.Printf("Please input int\n")
        fmt.Printf(PROMPT)
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
			fmt.Printf("Sum is %d\n", sum)
		case "-":
			diff := calculator.Substract(left, right)
			fmt.Printf("Diff is %d\n", diff)
		default:
			fmt.Printf("Invalid operator. got=%s\n", operator)
		}

        fmt.Print("\n");
	}
}
