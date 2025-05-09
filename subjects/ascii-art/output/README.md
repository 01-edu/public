## ascii-art-output

### Objectives

- You must follow the same [instructions](../README.md) as in the first subject **while** writing the result into a file.

The file must be named by using the flag `--output=<fileName.txt>`, in which `--output` is the flag and `<fileName.txt>` is the file name which will contain the output.

- The flag must have exactly the same format as above, any other formats must return the following usage message:

```console
Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard
```

If there are other `ascii-art` optional projects implemented, the program should accept other correctly formatted `[OPTION]` and/or `[BANNER]`.
Additionally, the program must still be able to run with a single `[STRING]` argument.

### Instructions

- Your project must be written in **Go**.
- The code must respect the [**good practices**](../../good-practices/README.md).
- It is recommended to have **test files** for [unit testing](https://go.dev/doc/tutorial/add-a-test).

### Usage

```console
student$ go run . --output=banner.txt "hello" standard
student$ cat -e banner.txt
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
student$ go run . --output=banner.txt 'Hello There!' shadow
student$ cat -e banner.txt
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
student$
```

### Allowed packages

- Only the [standard Go](https://pkg.go.dev/std) packages are allowed

This project will help you learn about :

- The Go file system(**fs**) API
- Data manipulation
