package main

import (
	"fmt"
	"go2021/model"
	"strconv"
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
	l_, _ := strconv.Atoi(l)
	return model.NewResult(t, l_, strings.Join(v, ""))
}

func main() {
	s := divideString("XX04JKML")
	result := asign(s)
	fmt.Println(result)
}
