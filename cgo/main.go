package main
// #include <stdio.h>
// #include <stdlib.h>
// #include "foo.h"
import "C"

func main() {
	C.foo()
}
