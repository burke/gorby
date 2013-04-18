package main

/*
#cgo CFLAGS: -DHAVE_STRUCT_TIMESPEC -I/Users/burke/src/g/ruby/include
#cgo LDFLAGS: -I. -L. -L/Users/burke/src/g/ruby -lgorby -lruby-static -lpthread -ldl -lobjc -Wl,-u,_objc_msgSend
#include "ruby.h"
#include "gorby.h"
*/
import "C"

import "os"

func main() {
	C.boot_vm()
	C.rb_eval_string(C.CString(os.Args[1]))
}
