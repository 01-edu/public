## displayfile

### Instructions

Write a program that displays, on the standard output, the content of a file given as argument.

-   Create a file `quest8.txt` and write into it the sentence `Almost there!!`

-   The argument of the program in this case should be, `quest8.txt`.

-   In case of error the program should print one the below messages accordingly:
    -   `File name missing`.
    -   `Too many arguments`.

### Usage :

```console
student@ubuntu:~/piscine-go/displayfile$ go build
student@ubuntu:~/piscine-go/displayfile$ ./displayfile
File name missing
student@ubuntu:~/piscine-go/displayfile$ ./displayfile quest8.txt main.go
Too many arguments
student@ubuntu:~/piscine-go/displayfile$ ./displayfile quest8.txt
Almost there!!
```
