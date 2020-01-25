package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    answer := rand.Intn(10) + 1
    count := 0

    for {
        var guess int

        fmt.Print("Your guess? ")
        fmt.Scanf("%v", &guess)
        // count = count + 1
        // count += 1
        count++
    
        if answer == guess {
            fmt.Printf("Bingo! It took %v guessed!\n", count)
            break
        } else if answer > guess {
            fmt.Println("The answer is higher!")
        } else {
            fmt.Println("The answer is lower!")
        }
    }
}
