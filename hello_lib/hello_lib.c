#include <Python.h>

#if PY_MAJOR_VERSION >= 3
#define PyHelloLib_Bytes_AS_STRING PyBytes_AS_STRING
#else
#define PyHelloLib_Bytes_AS_STRING PyString_AS_STRING
#endif

/* Will come from go */
PyObject* ohai(PyObject*);

/* To shim go's missing variadic function support */
int PyArg_ParseTuple_U(PyObject* args, PyObject** obj) {
    return PyArg_ParseTuple(args, "U", obj);
}

/* To shim go's lack of C macro support */
void PyHelloLib_DECREF(PyObject* obj) {
    Py_DECREF(obj);
}

const char* PyHelloLib_Bytes_AsString(PyObject* s) {
    return PyHelloLib_Bytes_AS_STRING(s);
}

static struct PyMethodDef methods[] = {
    {"ohai", (PyCFunction)ohai, METH_VARARGS},
    {NULL, NULL}
};

#if PY_MAJOR_VERSION >= 3
static struct PyModuleDef module = {
    PyModuleDef_HEAD_INIT,
    "hello_lib",
    NULL,
    -1,
    methods
};

PyMODINIT_FUNC PyInit_hello_lib(void) {
    return PyModule_Create(&module);
}
#else
PyMODINIT_FUNC inithello_lib(void) {
    Py_InitModule3("hello_lib", methods, NULL);
}
#endif
