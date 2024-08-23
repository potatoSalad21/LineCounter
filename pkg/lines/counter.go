package lines

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strings"
)

type Counter struct {
	Count int
}

func (c *Counter) CountAll(path string, exts []string) error {
	if path[len(path)-1] != '/' {
		path += "/"
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			err := c.CountAll(path+f.Name(), exts)
			if err != nil {
				log.Println(err, "[[skipping]]")
				continue
			}
		}

		extension := strings.Split(f.Name(), ".")
		if slices.Contains(exts, extension[len(extension)-1]) {
			c.countLines(path + f.Name())
		}
	}

	return nil
}

func (c *Counter) countLines(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Println("Couldn't open file: ", path)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c.Count++
	}
}
