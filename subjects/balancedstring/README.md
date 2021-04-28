## balancedstring

### Instructions

Balanced string is a string that has equal quantity of 'C' and 'D' characters.

Write a program that takes a string and outputs maximum amount of balanced strings without ignoring any letters.
Display output with `\n` at the end of line.

If the number of arguments is not 1, display nothing.

It will only be tested strings containing the characters 'C' and 'D'.

### Usage

```console
$ go run . "CDCCDDCDCD"
4
$ go run . "CDDDDCCCDC"
3
$ go run . "DDDDCCCC"
1
$ go run . "CDCCCDDCDD"
2
$
```

In first example "CDCCDDCDCD" can be split into "CD", "CCDD", "CD", "CD", each substring contains same number of 'C' and 'D'.

Second, "CDDDDCCCDC" can be split into: "CD", "DDDCCC", "DC".

"DDDDCCCC" can be split into "DDDDCCCC".

"CDCCCDDCDD" can be split into: "CD", "CCCDDCDD", since each substring contains an equal number of 'C' and 'D'
