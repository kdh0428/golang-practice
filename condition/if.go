package main 

import "fmt"

func main() {
    a := 7
    if a%2 == 0 {
        fmt.Println(a, "is even")
    } else {
        fmt.Println(a, "is odd")
    }

    if 8%4 == 0{
        fmt.Println("8 is divisible by 4")
    }

    if num := 9; num < 0{
        fmt.Println(num, "is negative")
    } else if num < 10{
    } else {
        fmt.Println(num, "has multiple digits")
    }
}
