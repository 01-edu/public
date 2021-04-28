## ascii-art-output

### Objectives

- You must follow the same [instructions](https://public.01-edu.org/subjects/ascii-art/) as in the first subject **while** writing the result into a file.

- The file must be named by using the flag `--output=<fileName.txt>`, in which `--output` is the flag and `<fileName.txt>` is the file name.

This project will help you learn about :

- The Go file system(**fs**) API.
- Ways to receive data.
- Ways to output data.
- Manipulation of strings.
- Choices of outputs.

### Instructions

- Your project must be written in **Go**.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices/).
- It is recommended that the code presents a **test file**.

### Allowed packages

- Only the [standard go](https://golang.org/pkg/) packages are allowed

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
