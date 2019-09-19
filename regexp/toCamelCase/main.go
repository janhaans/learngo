package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ToCamelCase(s string) string {
	ss := regexp.MustCompile(`[-_]`).Split(s, -1)
	for i := range ss {
		if i == 0 {
			continue
		}
		ss[i] = strings.Title(ss[i])
	}
	return strings.Join(ss, "")
}

func main() {
	s1 := "Hello_Big_world"
	s2 := "hello-big-world"
	s3 := "hello_big-world"
	fmt.Println(ToCamelCase(s1))
	fmt.Println(ToCamelCase(s2))
	fmt.Println(ToCamelCase(s3))
}
