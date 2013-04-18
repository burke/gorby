package main

/*
#cgo CFLAGS: -DHAVE_STRUCT_TIMESPEC -I/Users/burke/src/g/ruby/include
#cgo LDFLAGS: -I. -L. -L/Users/burke/src/g/ruby -lgorby -lruby-static -lpthread -ldl -lobjc -Wl,-u,_objc_msgSend
#include "gorby.h"

static char *rstring_ptr(VALUE rstring) {
  return RSTRING_PTR(rstring);
}

static int rstring_len(VALUE rstring) {
  return RSTRING_LEN(rstring);
}

*/
import "C"

import (
	"bufio"
	"fmt"
	"os"
)

func rstringToGostring(rstring C.VALUE) string {
	return C.GoStringN(C.rstring_ptr(rstring), C.rstring_len(rstring))
}

func evalString(code string) C.VALUE {
	val := C.gorby_eval(C.CString(code))
	return C.rb_funcall(val, C.rb_intern(C.CString("inspect")), 0)
}

func main() {
	C.boot_vm()
	r := bufio.NewReader(os.Stdout)
	for {
		fmt.Printf("gorby> ")
		line, _, _ := r.ReadLine()
		result := evalString(string(line))
		ln := rstringToGostring(result)
		fmt.Printf("=> %s\n", ln)
	}
}
