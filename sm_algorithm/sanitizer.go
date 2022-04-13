package smalgorithm

import (
    "regexp"
    "fmt"
)

func main() {
    r, a := regexp.Compile("[ACGT]+")
    fmt.Println(r, a)
}
