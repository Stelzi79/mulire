package main

import (
	"flag"
	"fmt"
)

func main() {

	var replacerName string
	flag.StringVar(&replacerName, "n", "replace", "The Name of the replacer to use")

	// numbPtr := flag.Int("numb", 42, "an int")
	// forkPtr := flag.Bool("fork", false, "a bool")

	// var svar string
	// flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.CommandLine.Usage = func() {
		fmt.Println("asdfasdf asd fafd asfd:")
		flag.PrintDefaults()
	}
	flag.Parse()

	fmt.Println("word:", replacerName)
	// fmt.Println("numb:", *numbPtr)
	// fmt.Println("fork:", *forkPtr)
	// fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
	fmt.Println("1.: ", flag.Arg(0))
	fmt.Println("2.: ", flag.Arg(1))

	fmt.Println("Hello, World!")

}
