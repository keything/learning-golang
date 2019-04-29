package main

import (
	"fmt"
)

type IToken interface {
	Type() string
	Lexeme() string
}

type Match struct {
	value string
}

func (m Match) Type() string {
	return "string"
}

func (m Match) Lexeme() string {
	return "match string" + m.value
}

type integerConstantV1 struct {
	IToken
	name string
}
type integerConstantV2 struct {
	name string
	Match
}
type integerConstantV3 struct {
	name string
	*Match
}

func main() {
	icv1 := integerConstantV1{
		name:   "icv1",
		IToken: Match{value: "matchv1"},
	}
	fmt.Println("icv1=", icv1.Type())
	icv2 := integerConstantV2{
		Match: Match{value: "matchv2"},
		name:  "icv2",
	}
	icv3 := integerConstantV3{
		Match: &Match{value: "matchv3"},
		name:  "icv3",
	}
	fmt.Println("icv1=", icv1.Lexeme())
	fmt.Println("icv2=", icv2.Lexeme())
	fmt.Println("icv3=", icv3.Lexeme())
}
