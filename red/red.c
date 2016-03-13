#include <Python.h>

#if PY_MAJOR_VERSION >= 3
#define PyRed_Bytes_AS_STRING PyBytes_AS_STRING
#else
#define PyRed_Bytes_AS_STRING PyString_AS_STRING
#endif

/* Will come from go */
PyObject* red(PyObject*);

/* To shim go's missing variadic function support */
int PyArg_ParseTuple_U(PyObject* args, PyObject** obj) {
    return PyArg_ParseTuple(args, "U", obj);
}

/* To shim go's lack of C macro support */
void PyRed_DECREF(PyObject* obj) {
    Py_DECREF(obj);
}

const char* PyRed_Bytes_AsString(PyObject* s) {
    return PyRed_Bytes_AS_STRING(s);
}

static struct PyMethodDef methods[] = {
    {"red", (PyCFunction)red, METH_VARARGS},
    {NULL, NULL}
};

#if PY_MAJOR_VERSION >= 3
static struct PyModuleDef module = {
    PyModuleDef_HEAD_INIT,
    "red",
    NULL,
    -1,
    methods
};

PyMODINIT_FUNC PyInit_red(void) {
    return PyModule_Create(&module);
}
#else
PyMODINIT_FUNC initred(void) {
    Py_InitModule3("red", methods, NULL);
}
#endif
