package main

import (
	"fmt"
	"github.com/wgarunap/go-cpp/cmd/cpp-cgo/cpp"
)

func main() {
	x := cpp.FuncCpp(2, 4)
	fmt.Println(x)

	x2 := cpp.FuncC(2, 4)
	fmt.Println(x2)

}
