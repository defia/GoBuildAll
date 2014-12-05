package main

import "testing"

func Test_LoadConfigr(t *testing.T) {
	c, err := Loadconfig("test.ini")
	if err != nil {
		t.Fatal(err)
	}
	for i := range c {
		t.Log(c[i])
	}
}
