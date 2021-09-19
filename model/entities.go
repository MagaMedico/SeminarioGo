package model

type Result struct {
	Type   string
	Length string
	Value  string
}

func NewResult(Type string, Length string, Value string) Result {
	return Result{Type, Length, Value}
}
