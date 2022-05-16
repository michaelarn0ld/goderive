package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inputEquation(order *int, args []int) string {
	var buffer strings.Builder
	for i := 0; i <= *order; i++ {
		if *order-i == 0 {
			buffer.WriteString(fmt.Sprintf("%d", args[i]))
		} else if *order-i == 1 {
			buffer.WriteString(fmt.Sprintf("%dx + ", args[i]))
		} else {
			buffer.WriteString(fmt.Sprintf("%dx^%d + ", args[i], *order-i))
		}
	}
	return buffer.String()
}

func outputEquation(args []int) string {
	var buffer strings.Builder
	var d int
	order := len(args) - 1
	for i := order; i > 0; i-- {
		d = args[len(args)-1-i] * i
		if i > 2 {
			buffer.WriteString(fmt.Sprintf("%dx^%d + ", d, order-1))
		} else if i == 2 {
			buffer.WriteString(fmt.Sprintf("%dx + ", d))
		} else {
			buffer.WriteString(fmt.Sprintf("%d ", d))
		}
	}
	return buffer.String()
}

func main() {
	order := flag.Int("order", -1, "The order of the equation to be derived")
	flag.Parse()

	args := flag.Args()

	if len(args) != *order+1 {
		os.Exit(1)
	}

	intArgs := make([]int, len(args))
	for i, val := range args {
		if j, err := strconv.Atoi(val); err != nil {
			os.Exit(1)
		} else {
			intArgs[i] = j
		}
	}

	fmt.Println(inputEquation(order, intArgs))
	fmt.Println("Deriving...")
	fmt.Println(outputEquation(intArgs))

}
