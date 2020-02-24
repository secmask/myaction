package main

import (
	"testing"
)

func TestAA(t *testing.T) {
	if Add(1,2)!=3{
		t.Fatal("not equal")
	}
}