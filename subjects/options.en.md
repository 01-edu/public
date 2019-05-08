# options

## Instructions

Write a program that takes an undefined number of arguments which could be considered as options and writes on standard output a representation of those options as groups of bytes followed by a newline.

- An option is an argument that begins by a '-' and have multiple characters which could be :
	- abcdefghijklmnopqrstuvwxyz

- All options are stocked in a single int and each options represents a bit of that int, and should be stocked like this :

	- 00000000 00000000 00000000 00000000
	- ******zy xwvutsrq ponmlkji hgfedcba

- Launch the program without arguments or with the '-h' flag activated must print an usage on the standard output, as shown in the following examples.

- A wrong option must print "Invalid Option" followd by a newline.

## Expected output

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test | cat -e
options: abcdefghijklmnopqrstuvwxyz$
student@ubuntu:~/piscine/test$ ./test -abc -ijk | cat -e
00000000 00000000 00000111 00000111$
student@ubuntu:~/piscine/test$ ./test -z | cat -e
00000010 00000000 00000000 00000000$
student@ubuntu:~/piscine/test$ ./test -abc -hijk | cat -e
options: abcdefghijklmnopqrstuvwxyz$
student@ubuntu:~/piscine/test$ ./test -% | cat -e
Invalid Option$
student@ubuntu:~/piscine/test$
```