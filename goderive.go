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

	input := inputEquation(order, intArgs)
	fmt.Println(input)
}
