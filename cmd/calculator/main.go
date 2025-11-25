package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Dd1235/CLI-Calculator/internal/calc"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  calculator <op> <a> <b>")
	fmt.Println("Ops: add | sub | mul | div")
}

func main() {
	if len(os.Args) != 4 {
		usage()
		os.Exit(1)
	}

	op := os.Args[1]
	a, err1 := strconv.ParseFloat(os.Args[2], 64)
	b, err2 := strconv.ParseFloat(os.Args[3], 64)
	if err1 != nil || err2 != nil {
		log.Fatalf("a and b must be numbers")
	}

	switch op {
	case "add":
		fmt.Println(calc.Add(a, b))
	case "sub":
		fmt.Println(calc.Sub(a, b))
	case "mul":
		fmt.Println(calc.Mul(a, b))
	case "div":
		res, err := calc.Div(a, b)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	default:
		usage()
		os.Exit(1)
	}
}
