# reversestrcap
## Instructions

Write a program that takes one or more strings and, for each argument, puts 
the last character of each word (if it's a letter) in uppercase and the rest
in lowercase, then displays the result followed by a \n.

A word is a section of string delimited by spaces/tabs or the start/end of the
string. If a word has a single letter, it must be capitalized.

If there are no parameters, display \n.

## Expected output

```console
student@ubuntu:~/student/reversestrcap$ go build
student@ubuntu:~/student/reversestrcap$ ./reversestrcap "First SMALL TesT" | cat -e
firsT smalL tesT$
student@ubuntu:~/student/reversestrcap$ ./reversestrcap go run reversestrcap.go "SEconD Test IS a LItTLE EasIEr" "bEwaRe IT'S NoT HARd WhEN " " Go a dernier 0123456789 for the road e" | cat -e
seconD tesT iS A littlE easieR$
bewarE it'S noT harD wheN $
 gO A dernieR 0123456789 foR thE roaD E$
student@ubuntu:~/student/reversestrcap$ ./reversestrcap | cat -e
$
student@ubuntu:~/student/reversestrcap$
```
