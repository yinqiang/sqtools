package main

import (
	"testing"
)

func Test_Cross(t *testing.T) {
	e := Cross([]string{"", "config.json"})
	if e != nil {
		t.Fatal(e.Error())
	}
}
