from setuptools import Extension
from setuptools import setup


setup(
    name='setuptools-golang-examples',
    description='Examples for https://github.com/asottile/setuptools-golang',
    url='https://github.com/asottile/setuptools-golang-examples',
    version='0.1.1',
    author='Anthony Sottile',
    author_email='asottile@umich.edu',
    classifiers=[
        'License :: OSI Approved :: MIT License',
        'Programming Language :: Python :: 2',
        'Programming Language :: Python :: 2.7',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.4',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: Implementation :: CPython',
        'Programming Language :: Python :: Implementation :: PyPy',
    ],
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
