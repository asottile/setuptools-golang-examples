package main

// #include <stdlib.h>
// #include <Python.h>
// int PyArg_ParseTuple_U(PyObject*, PyObject**);
// const char* PyRed_Bytes_AsString(PyObject*);
// void PyRed_DECREF(PyObject*);
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
	cstr := C.PyRed_Bytes_AsString(bytes)
	red := ansi.Color(C.GoString(cstr), "red")
	cstr = C.CString(red)
	ret := C.PyUnicode_FromString(cstr)

	C.free(unsafe.Pointer(cstr))
	C.PyRed_DECREF(bytes)

	return ret
}

func main() {}
