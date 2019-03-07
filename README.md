[![Build Status](https://asottile.visualstudio.com/asottile/_apis/build/status/asottile.setuptools-golang-examples?branchName=master)](https://asottile.visualstudio.com/asottile/_build/latest?definitionId=12&branchName=master)
[![Azure DevOps coverage](https://img.shields.io/azure-devops/coverage/asottile/asottile/12/master.svg)](https://dev.azure.com/asottile/asottile/_build/latest?definitionId=12&branchName=master)

setuptools-golang-examples
==========================

A few examples utilizing [setuptools-golang](https://github.com/asottile/setuptools-golang).

## `c_module`

- Demonstrates that you can mix go extensions with c extensions seamlessly.

## `go_sum`

- A very basic hello-world-y demo
- This example is roughly lifted from @[FiloSottile](https://github.com/FiloSottile)'s [blog post](https://blog.filippo.io/building-python-modules-with-go-1-5/)

## `hello_lib`

- This module demonstrates importing go code within the project.

## `red`

- This module demonstrates importing external code (in this case [ansi](https://github.com/mgutz/ansi))

## `sum_pure_go`

- This module demonstrates it is possible to write an extension using only go
  files.
- It's slightly cheaty in that one of the go files is entirely a C header.
- You could instead do something similar to [this example](https://blog.filippo.io/building-python-modules-with-go-1-5/#bonustheneedlesslyhardway)
  but it's much more difficult (impossible?) to support py2 + py3.
