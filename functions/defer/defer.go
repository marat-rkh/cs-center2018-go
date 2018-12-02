package errors

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func writeData(data []string) (string, error) {
	filename := "./data" + strconv.Itoa(rand.Int()) + ".txt"
	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("writing file: %v", err)
	}
	defer file.Close()
	err = doWrite(file, data)
	if err != nil {
		return "", fmt.Errorf("writing file: %v", err)
	}
	return filename, nil
}

func doWrite(file *os.File, data []string) error {
	// ...
	return nil
}
