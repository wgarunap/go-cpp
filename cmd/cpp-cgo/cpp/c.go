package cpp

/*
#include <stdio.h>
#include "c.h"
*/
import "C"

import (
	"fmt"
)

func FuncCpp(x, y int) int64 {
	res := C.testFuncCpp(C.int64(x), C.int64(y))
	fmt.Println(x)
	return int64(res)
}

func FuncC(x, y int) int64 {
	res := C.testFunc(C.int64(x), C.int64(y))
	fmt.Println(x)
	return int64(res)
}
