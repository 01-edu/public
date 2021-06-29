## ascii-reverse

### Objectives

You must follow the same [instructions](https://public.01-edu.org/subjects/ascii-art/) as in the first subject but this time the process will be reversed. desrever fo dnik siht toN.

Ascii-reverse consists on reversing the process, converting the graphic representation into a text. You will have to create a text file containing a graphic representation of a random `string` given as an argument.

The argument will be a **flag**, `--reverse=<fileName>`, in which `--reverse` is the flag and `<fileName>` is the file name. The program must then print this `string` in **normal text**.

### Instructions

- Your project must be written in **Go**.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices/).
- It is recommended that the code should present a **test file**.

### Allowed packages

- Only the [standard go](https://golang.org/pkg/) packages are allowed

### Usage

```console
$ cat file.txt
 _              _   _
| |            | | | |
| |__     ___  | | | |   ___
|  _ \   / _ \ | | | |  / _ \
| | | | |  __/ | | | | | (_) |
|_| |_|  \___| |_| |_|  \___/



$ go run . --reverse=file.txt
hello
$
```

This project will help you learn about :

- The Go file system(**fs**) API
- Data manipulation
