package hermes

import "fmt"

type Printer interface {
	Print(a []byte)
}

type StdoutPrinter struct{}

func (s StdoutPrinter) Print(a []byte) {
	fmt.Println(string(a))
}
