// Echo3 prints its command-line arguements.
package main
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " ")) // joins string elements of the slice, add
	// joins 'em using a space
	fmt.Println(os.Args[1:]) // prints the slice itself
}
