package main

import "fmt"
import "strings"

func main() {
    var secret interface{}

    secret = 2
    var number = secret.(int) * 10
    fmt.Println(secret, "multiplied by 10 is :", number)

    secret = []string{"apple", "manggo", "banana"}
    var gruits = strings.Join(secret.([]string), ", ")
    // Join disini berfungsi untuk menyatukan sebuah array dah menggunakan ', ' sebagai pemisahnya/separatornya
    fmt.Println(gruits, "is my favorite fruits")
}
