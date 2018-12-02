package panic

import (
	"fmt"
	"log"
)

func writeData(data []string, lineLen int) (string, error) {
	if lineLen <= 0 {
		panic("incorrect line length")
	}
	// As before: create file and write
	return "", nil
}

func process() {
	defer func() {
		val := recover()
		if val != nil {
			log.Printf("panicking: %v\n", val)
		}
	}()
	writeData([]string{}, -1)
	fmt.Println("this is not executed")
}

func main() {
	process()
	fmt.Println("normal execution continues")
}
