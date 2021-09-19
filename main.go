package main

import (
	"fmt"
	"go2021/model"
	"strings"
)

func divideString(s string) []string {
	split := strings.Split(s, "")
	return split
}

func asign(values []string) model.Result {
	t := values[0] + values[1]
	l := values[2] + values[3]
	v := values[4:]
	return model.NewResult(t, l, strings.Join(v, ""))
}

func main() {
	s := divideString("XX04JKML")
	result := asign(s)
	fmt.Println(result)
}
