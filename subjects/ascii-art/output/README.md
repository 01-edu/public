## ascii-art-output

### Objectives

- You must follow the same [instructions](../README.md) as in the first subject **while** writing the result into a file. Yes, you will read from one file and write to another.

The file must be named by using the flag `--output=<fileName.txt>`, in which `--output` is the flag and `<fileName.txt>` is the file name which will contain the output.

### Instructions

- Your project must be written in **Go**.
- The code must respect the [**good practices**](../../good-practices/README.md).
- It is recommended that the code presents a **test file**.

### Usage

```console
$ go run . "hello" standard --output=banner.txt
$ cat banner.txt
 _              _   _
| |            | | | |
| |__     ___  | | | |   ___
|  _ \   / _ \ | | | |  / _ \
| | | | |  __/ | | | | | (_) |
|_| |_|  \___| |_| |_|  \___/



$ go run . "Hello There!" shadow --output=banner.txt
$ cat banner.txt

_|    _|          _| _|                _|_|_|_|_| _|                                  _|
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _|
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _|
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _|



$
```

### Allowed packages

- Only the [standard Go](https://golang.org/pkg/) packages are allowed

This project will help you learn about :

- The Go file system(**fs**) API
- Data manipulation
