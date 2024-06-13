package main

// #include <stdlib.h>
// #include <Python.h>
// int PyArg_ParseTuple_s(PyObject*, char**);
import "C"
import "unsafe"
import "github.com/mgutz/ansi"

//export red
func red(self *C.PyObject, args *C.PyObject) *C.PyObject {
	var cstr *C.char
	if C.PyArg_ParseTuple_s(args, &cstr) == 0 {
		return nil
	}
	red := ansi.Color(C.GoString(cstr), "red")

	gocstr := C.CString(red)
	ret := C.PyUnicode_FromString(gocstr)
	C.free(unsafe.Pointer(gocstr))

	return ret
}

func main() {}
