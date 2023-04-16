package main

import (
	"fmt"
	"os"

	"github.com/XiXiangFiles/gitcralwer/internal/cmd"
	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("print", "Prints provided string to stdout")

	path := parser.String("p", "path", &argparse.Options{Required: true, Help: "git comments doc path"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	fmt.Println(cmd.ParseMessage(*path))
}
