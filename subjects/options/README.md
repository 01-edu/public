# options

## Instructions

Write a program that takes an undefined number of arguments which could be considered as `options` and writes on the standard output a representation of those `options` as groups of `bytes` followed by a newline (`'\n'`).

- An `option` is an argument that begins with a `-` and that can have multiple characters which could be :
  -abcdefghijklmnopqrstuvwxyz

- All `options` are stocked in a single `int` and each `options` represents a bit of that `int`, and should be stocked like this :

            - 00000000 00000000 00000000 00000000
            - ******zy xwvutsrq ponmlkji hgfedcba

- Launching the program without arguments or with the `-h` flag activated must print all the valid `options` on the standard output, as shown on one of the following examples.

- Please note the `-h` flag has priority over the others flags when it is called first in one of the arguments. (See the examples)

- A wrong `option` must print `Invalid Option` followed by a newline.

### Usage

```console
$ go run . | cat -e
options: abcdefghijklmnopqrstuvwxyz$
$ go run . -abc -ijk | cat -e
00000000 00000000 00000111 00000111$
$ go run . -z | cat -e
00000010 00000000 00000000 00000000$
$ go run . -abc -hijk | cat -e
options: abcdefghijklmnopqrstuvwxyz$
$ go run . -h | cat -e
options: abcdefghijklmnopqrstuvwxyz$
$ go run . -zh | cat -e
00000010 00000000 00000000 10000000$
$ go run . -z -h | cat -e
options: abcdefghijklmnopqrstuvwxyz$
$ go run . -hhhhhh | cat -e
options: abcdefghijklmnopqrstuvwxyz$
$ go run . -eeeeee | cat -e
00000000 00000000 00000000 00010000$
$ go run . -% | cat -e
Invalid Option$
$ go run . - | cat -e
Invalid Option$
$
```
