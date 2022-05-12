package main

import (
    "fmt"
    "io/ioutil"
    "log"
	"strings"
)

func main() {
	// reading file
    content, err := ioutil.ReadFile("go.txt")

    if err != nil {
        log.Fatal(err)
    }

    res := strings.Split(string(content), " ")
	fmt.Println(len(res))


}