## reversestrcap

### Instructions

Write a program that takes one or more `string` as arguments and that, **for each argument**:
-puts the last character of each word (if it is a letter) in uppercase and the rest
in lowercase
-then it displays the result followed by a newline (`'\n'`).

If there are no parameter, the program displays a newline.

### Usage

```console
student@ubuntu:~/[[ROOT]]/reversestrcap$ go build
student@ubuntu:~/[[ROOT]]/reversestrcap$ ./reversestrcap "First SMALL TesT" | cat -e
firsT smalL tesT$
student@ubuntu:~/[[ROOT]]/reversestrcap$ ./reversestrcap go run reversestrcap.go "SEconD Test IS a LItTLE EasIEr" "bEwaRe IT'S NoT HARd WhEN " " Go a dernier 0123456789 for the road e" | cat -e
seconD tesT iS A littlE easieR$
bewarE it'S noT harD wheN $
 gO A dernieR 0123456789 foR thE roaD E$
student@ubuntu:~/[[ROOT]]/reversestrcap$ ./reversestrcap | cat -e
$
student@ubuntu:~/[[ROOT]]/reversestrcap$
```
