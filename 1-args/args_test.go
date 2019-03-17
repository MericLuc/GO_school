package main

import (
	"fmt"
	"testing"
)

/*
go test -bench=.
*/

var (
	args = []string{"exec arg0 arg1 arg2 arg3"}
)

func Testargs1(t *testing.T) {
	fmt.Println("args:", args)
	if len(args) > 0 {
		args1(args)
	} else {
		fmt.Println("no args")
	}
}

func Benchmarkargs1(b *testing.B) {
	fmt.Println("args:", args)
	if len(args) > 0 {
		for i := 0; i < b.N; i++ {
			args1(args)
		}
	} else {
		fmt.Println("no args")
	}
}

func Benchmarkargs2(b *testing.B) {
	fmt.Println("args:", args)
	if len(args) > 0 {
		for i := 0; i < b.N; i++ {
			args2(args)
		}
	} else {
		fmt.Println("no args")
	}
}

func Benchmarkargs3(b *testing.B) {
	fmt.Println("args:", args)
	if len(args) > 0 {
		for i := 0; i < b.N; i++ {
			args3(args)
		}
	} else {
		fmt.Println("no args")
	}
}

func Benchmarkargs4(b *testing.B) {
	fmt.Println("args:", args)
	if len(args) > 0 {
		for i := 0; i < b.N; i++ {
			args4(args)
		}
	} else {
		fmt.Println("no args")
	}
}