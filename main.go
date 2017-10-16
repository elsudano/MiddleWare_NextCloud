package main

import "fmt"
import "strconv"

type Rectangle struct {
    length, width int
}

func (r Rectangle) GetLengthValue() string {
    return strconv.Itoa(r.length)
}

func (r Rectangle) GetWidthValue() string {
    return strconv.Itoa(r.width)
}

func main() {
    r := Rectangle{}
    r.length = 1
    r.width = 1
    fmt.Println("Default rectangle is: ", r.GetLengthValue)
}
