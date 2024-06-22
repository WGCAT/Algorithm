package main

import "fmt"

func main() {
    var s1, s2 string
    fmt.Scan(&s1, &s2)
    
    str := s1 + s2
    fmt.Println(str)
}