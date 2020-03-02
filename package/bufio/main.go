package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if Question("今日は元気ですか？[y/n] ") {
		fmt.Println("そうですか、頑張ってください！")
	} else {
		fmt.Println("元気出してくださいね！")
	}
}

func Question(q string) bool {
	result := true
	fmt.Print(q)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i := scanner.Text()

		if i == "Y" || i == "y" {
			break
		} else if i == "N" || i == "n" {
			result = false
			break
		} else {
			fmt.Println("yかnで答えてください。")
			fmt.Print(q)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result
}
