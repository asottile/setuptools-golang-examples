# DEPRECATED

it turns out multiple go shared objects in a single process is not supported

it likely broke in [go 1.21] and there is no intention to fix it :(

[go 1.21]: https://github.com/golang/go/issues/65050#issue-2074509727

___

[![build status](https://github.com/asottile/setuptools-golang-examples/actions/workflows/main.yml/badge.svg)](https://github.com/asottile/setuptools-golang-examples/actions/workflows/main.yml)
[![pre-commit.ci status](https://results.pre-commit.ci/badge/github/asottile/setuptools-golang-examples/main.svg)](https://results.pre-commit.ci/latest/github/asottile/setuptools-golang-examples/main)

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
  but it's much more difficult to support multiple versions of python.
