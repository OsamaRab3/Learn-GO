package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Head struct {
	lines int
	Files []string
}

func (h *Head) read_file() {
	for _, filename := range h.Files {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Head: %v \n", err)
			continue
		}
		fmt.Println("========| ", filename, " |=======")
		scanner := bufio.NewScanner(file)
		count := 0

		for scanner.Scan() && count < h.lines {
			fmt.Println(scanner.Text())
			count++
		}
		file.Close()

	}
}
func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: head <lines> <file1> [file2...]")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid number of lines")
		return
	}
	h := Head{
		lines: n,
		Files: os.Args[2:],
	}
	h.read_file()

}
