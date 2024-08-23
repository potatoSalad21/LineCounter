package main

import (
	"fmt"
	"log"
	"os"

	"github.com/potatoSalad21/linecounter/pkg/lines"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: lc [proj-path]")
		return
	}

	var (
		path = os.Args[1]
		exts = []string{"go", "c", "js", "cpp", "html", "py", "ts", "rs", "cs"}
	)
	c := lines.NewCounter()

	err := c.CountAll(path, exts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d lines in the project\n", c.Count)
}
