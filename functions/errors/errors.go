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

func main() {
	todo := []string{"TODO", "1. Keep calm", "2. And code some Go"}
	filename, err := writeData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote to file %s", filename)
}
