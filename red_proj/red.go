package main

// #include <stdlib.h>
// #include <Python.h>
// int PyArg_ParseTuple_U(PyObject*, PyObject**);
import "C"
import "unsafe"
import "github.com/mgutz/ansi"

//export red
func red(self *C.PyObject, args *C.PyObject) *C.PyObject {
	var obj *C.PyObject
	if C.PyArg_ParseTuple_U(args, &obj) == 0 {
		return nil
	}
	bytes := C.PyUnicode_AsUTF8String(obj)
	cstr := C.PyBytes_AsString(bytes)
	red := ansi.Color(C.GoString(cstr), "red")
	cstr = C.CString(red)
	ret := C.PyUnicode_FromString(cstr)

	C.free(unsafe.Pointer(cstr))
	C.Py_DecRef(bytes)

	return ret
}

func main() {}
