package interfaces

import (
	"fmt"
	"io"
)

//Greet someone
func Greet(print io.Writer, name string) {
	//print.Write([]Byte)
	fmt.Fprintf(print, "Hello, %s", name)
}
