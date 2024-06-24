// A ~ Z : 65 ~ 90, a ~ z : 97 ~ 122

package main

import (
"fmt"
"strings"
)
func main() {
    var s1 string
    fmt.Scan(&s1)
    
    var result string
    
    slice := strings.Split(s1,"")
    
    for _, v := range slice {
        compare := []rune(v)
        if compare >= 'A' && compare <= 'Z' {
            compare = string(compare)
            compare = strings.ToUpper(compare)
            result = result + compare
        }else {
            compare = string(compare)
            compare = strings.ToLower(compare)
            result = result + compare
        }
    }
    fmt.Println(result)
}