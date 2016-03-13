from setuptools import Extension
from setuptools import setup


setup(
    name='setuptools-golang-examples',
    version='0.1.0',
    ext_modules=[
        Extension('c_module', ['c_module/c_module.c']),
        Extension('hello_lib', ['hello_lib/hello_lib.go']),
        Extension('red', ['red/red.go']),
        Extension('sum_go', ['sum_go/sum_go.go']),
        Extension('sum_pure_go', ['sum_pure_go/sum_pure_go.go']),
    ],
    build_golang={'root': 'github.com/asottile/setuptools-golang-examples'},
    setup_requires=['setuptools-golang'],
)
