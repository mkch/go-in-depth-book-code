package main

import (
	"flag"
	"os"
	"testing"
)

var arg1 int
var v bool

func TestMain(m *testing.M) {
	flag.IntVar(&arg1, "arg1", 0, "")
	flag.BoolVar(&v, "v", false, "")
	flag.Parse()
	m.Run()
}

func TestFlags(t *testing.T) {
	t.Log("os.Args:", os.Args)
	t.Log("arg1:", arg1)
	t.Log("v:", v)
}
