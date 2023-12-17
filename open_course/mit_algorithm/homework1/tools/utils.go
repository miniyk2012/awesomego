package tools

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func ReadNumbers(path string) (arr []int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		num, err := strconv.Atoi(string(line))
		if err != nil {
			break
		}
		arr = append(arr, num)
	}
	return arr
}
