package demo

import "testing"

func TestMyfun(t *testing.T) {
	r := Myfun(100)
	if r != 55 {
		t.Errorf(" Myfun(10) ")
	}
}