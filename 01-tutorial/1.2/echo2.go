// Echo2 prints its command-line arguements.
package main
import (
	"fmt"
	"os"
)

func main() {
	// this
	s, sep := "", ""
	// or this
	// var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	// range returns index, value from an iterable data structure
	fmt.Println(s);
}
