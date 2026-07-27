/*
main.go
gorun
*/
package main
import (
	"fmt"
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
	r.Render()
}


