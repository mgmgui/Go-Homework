package hello

import (
	"github.com/mgmgui/go-homework/hello/calculator"
	"rsc.io/quote/v3"
)

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}

func Add() int {
	return calculator.Add(1, 2)
}
