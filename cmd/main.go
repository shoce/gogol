/*
main.go
gorun
*/
package main
import (
	"fmt"
	"os"
	golr "github.com/shoce/gogol/render"
)
const (
	NL = "\n"
)
var (
	F = fmt.Sprintf
	pout = fmt.Print
)
func main() {
	r := golr.Make(111, 111, 1)
	/*
	if r==nil { 
		perr("render.Make nil")
		os.Exit(1)
	}
	*/
	r.Render()
}
func perr(msg string) {
	fmt.Fprint(os.Stderr, msg+NL)
}

