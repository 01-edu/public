## ascii-art-fs

### Objectives

You must follow the same [instructions](https://public.01-edu.org/subjects/ascii-art/) as in the first subject but the second argument must be the name of the template.

This project will help you learn about :

- The Go file system(**fs**) API.
- Ways to receive data.
- Ways to output data.
- Manipulation of strings.
- Manipulation of structures.

### Instructions

- Your project must be written in **Go**.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices/).
- It is recommended that the code should present a **test file**.
- You can see all about the **banners** [here](https://github.com/01-edu/public/tree/master/subjects/ascii-art).

### Allowed packages

- Only the [standard go](https://golang.org/pkg/) packages are allowed

### Usage

```console
student@ubuntu:~/ascii-art$ go build
student@ubuntu:~/ascii-art$ ./ascii-art "hello" standard
  _                _    _
 | |              | |  | |
 | |__      ___   | |  | |    ___
 |  _ \    / _ \  | |  | |   / _ \
 | | | |  |  __/  | |  | |  | (_) |
 |_| |_|   \___|  |_|  |_|   \___/



student@ubuntu:~/ascii-art$ ./ascii-art "Hello There!" shadow

_|    _|          _| _|                _|_|_|_|_| _|                                  _|
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _|
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _|
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _|



student@ubuntu:~/ascii-art$ ./ascii-art "Hello There!" thinkertoy

o  o     o o           o-O-o o
|  |     | |             |   |                o
O--O o-o | | o-o         |   O--o o-o o-o o-o |
|  | |-' | | | |         |   |  | |-' |   |-' o
o  o o-o o o o-o         o   o  o o-o o   o-o
                                              O


student@ubuntu:~/ascii-art$
```
