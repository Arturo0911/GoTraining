package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var names string
	var address string
	scanner.Scan()
	names = scanner.Text()
	address = scanner.Text()
	fmt.Println(names, " vive en: ", address)
}
