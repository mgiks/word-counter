package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    fmt.Println("Print the text:")

    reader := bufio.NewReader(os.Stdin)

    text, _ := reader.ReadString('\n')

    wordCount := countWords(text)

    fmt.Printf("Text contains %d words \n", wordCount)
}
