package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Butler service initail...")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input service")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	fmt.Println(scanner.Err())
}