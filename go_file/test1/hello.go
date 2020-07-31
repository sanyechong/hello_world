package main
import (
    "fmt"
    "strings"
    "strconv"
)

func main () {
    fmt.Printf("hello world.\n");
    length := len("yyyaaa")
    slength := strconv.Itoa(length)
    sfloat := "3.141592653"
    float, _ := strconv.ParseFloat(sfloat, 64)   
    fmt.Printf("%f", float)
    fmt.Printf("%s", slength+"\n")

    str := "The quick brown fox jumps over the lazy dog"
    sl := strings.Fields(str)
    fmt.Printf("Splitted in slice: %v\n", sl)
    for _, val := range sl {
        fmt.Printf("%s - ", val)
    }
    fmt.Println()
    str2 := "GO1|The ABC of Go|25"
    sl2 := strings.Split(str2, "|")
    fmt.Printf("Splitted in slice: %v\n", sl2)
    for _, val := range sl2 {
        fmt.Printf("%s - ", val)
    }
    fmt.Println()
    str3 := strings.Join(sl2,";")
    fmt.Printf("sl2 joined by ;: %s\n", str3)
}
