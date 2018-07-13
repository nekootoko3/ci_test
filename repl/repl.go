package repl

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"../aa"
	"../calculator"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print("----- Let's calculate -----\n")

		fmt.Print("Please input int\n")
		fmt.Printf(PROMPT)
		ls := scanner.Scan()
		if !ls {
			printErrorAa(out, "Error: NOT INT")
			continue
		}
		left, err := strconv.Atoi(scanner.Text())
		if err != nil {
			printErrorAa(out, "Error: NOT INT")
			continue
		}

		fmt.Printf("Please input operator ( + or -)\n")
		fmt.Printf(PROMPT)
		os := scanner.Scan()
		if !os {
			printErrorAa(out, "Error: NOT OPERATOR")
			continue
		}
		operator := scanner.Text()

		fmt.Printf("Please input int\n")
		fmt.Printf(PROMPT)
		rs := scanner.Scan()
		if !rs {
			printErrorAa(out, "Error: NOT INT")
			continue
		}

		right, err := strconv.Atoi(scanner.Text())
		if err != nil {
			printErrorAa(out, "Error: NOT INT")
			continue
		}

		switch operator {
		case "+":
			sum := calculator.Add(left, right)
			fmt.Printf("Sum is %d\n", sum)
		case "-":
			diff := calculator.Substract(left, right)
			fmt.Printf("Diff is %d\n", diff)
		default:
			msg := fmt.Sprintf("Error: Invalid operator. got=%s\n", operator)
			printErrorAa(out, msg)
		}

		fmt.Print("\n")
	}
}

func printErrorAa(out io.Writer, msg string) {
	aa := aa.New()
	io.WriteString(out, aa.PrintAa())
	io.WriteString(out, msg+"\n")
}
