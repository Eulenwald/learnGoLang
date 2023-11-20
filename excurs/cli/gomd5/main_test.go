package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestPrintMD5(t *testing.T) {
	in := strings.NewReader("golang")
	out := &bytes.Buffer{}

	printMD5(in, out)

	var want string = "21cc28409729565fc1a4d2dd92db269f"

	var got string = out.String()

	if got != want {
		fmt.Printf(" in ist vom Typ %T\nout ist vom Typ %T\n", in, out)
		fmt.Printf("%d %s\n%d %s\n", len(want), want, len(got), got)
		t.Error("So nich!")
	}
}
