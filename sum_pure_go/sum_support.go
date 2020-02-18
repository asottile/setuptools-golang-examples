package main

// #include <Python.h>
//
// PyObject* sum(PyObject* , PyObject*);
//
// int PyArg_ParseTuple_ll(PyObject* args, long* a, long* b) {
//     return PyArg_ParseTuple(args, "ll", a, b);
// }
//
// static struct PyMethodDef methods[] = {
//     {"sum", (PyCFunction)sum, METH_VARARGS},
//     {NULL, NULL}
// };
//
// static struct PyModuleDef module = {
//     PyModuleDef_HEAD_INIT,
//     "sum_pure_go",
//     NULL,
//     -1,
//     methods
// };
//
// PyMODINIT_FUNC PyInit_sum_pure_go(void) {
//     return PyModule_Create(&module);
// }
import "C"
