package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	operation := flag.String("op", "", "add|subtract|multiply|divide")
	flag.Parse()
	args := flag.Args()

	if *operation == "" || (*operation != "add" && *operation != "subtract" && *operation != "multiply" && *operation != "divide") {
		flag.PrintDefaults()
		os.Exit(1)
	} else if len(args) != 2 {
		fmt.Println("Must have exactly two arguments")
		os.Exit(1)
	}

	for _, val := range args {
		if _, err := strconv.ParseFloat(val, 32); err != nil {
			fmt.Println("Both arguments must be valid numbers")
			os.Exit(1)
		}
	}

	a, _ := strconv.ParseFloat(args[0], 32)
	b, _ := strconv.ParseFloat(args[1], 32)

	var c float64
	switch *operation {
	case "add":
		c = a + b
	case "subtract":
		c = a - b
	case "multiply":
		c = a * b
	case "divide":
		if b == 0 {
			fmt.Println("Cannot divide by 0")
			os.Exit(1)
		}
		c = a / b
	}
	fmt.Printf("%.2f\n", c)
}
