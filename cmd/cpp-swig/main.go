package main

import (
	"fmt"
	"github.com/wgarunap/go-cpp/cmd/cpp-swig/cpp"
)

func main() {
	x := cpp.TestFuncCpp(2, 4)
	fmt.Println(x)

	x = cpp.TestFunc(444, 4)
	fmt.Println(x)

}
