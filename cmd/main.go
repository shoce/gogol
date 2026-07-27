/*
main.go
gorun
*/
package main
import (
	"fmt"
	"os"
	gol "github.com/shoce/gogol"
)
const (
	NL = "\n"
)
var (
	F = fmt.Sprintf
	pout = fmt.Print
)
func main() {
	r := gol.Make(111, 111, 1)
	r.Render()
}


