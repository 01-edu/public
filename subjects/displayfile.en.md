## displayfile

### Instructions

Write a program that displays, on the standard output, only the content of the file given as argument.

- Create a file `quest8.txt` and write into the file `Almost there!!`

- The argument of the program should be the name of the file, in this case, `quest8.txt`.

- In case of error it should print:
  - `File name missing`.
  - `Too many arguments`.

### Output:

```console
student@ubuntu:~/student/displayfile$ go build
student@ubuntu:~/student/displayfile$ ./displayfile
File name missing
student@ubuntu:~/student/displayfile$ ./displayfile quest8.txt main.go
Too many arguments
student@ubuntu:~/student/displayfile$ ./displayfile quest8.txt
Almost there!!
```
