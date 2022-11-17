## romannumbers

### Instructions

Write a program called `rn`. The objective is to convert a number, given as an argument, into a roman number and print it with its roman number calculation.

The program should have a limit of `4000`. In case of an invalid number, for example `"hello"` or `0` the program should print `ERROR: cannot convert to roman digit`.

Roman Numerals reminder:

|  I  |  1   |
| :-: | :--: |
|  V  |  5   |
|  X  |  10  |
|  L  |  50  |
|  C  | 100  |
|  D  | 500  |
|  M  | 1000 |

For example, the number 1732 would be denoted MDCCXXXII in Roman numerals. However, Roman numerals are not a purely additive number system. In particular, instead of using four symbols to represent a 4, 40, 9, 90, etc. (i.e., IIII, XXXX, VIIII, LXXXX, etc.), such numbers are instead denoted by preceding the symbol for 5, 50, 10, 100, etc., with a symbol indicating subtraction. For example, 4 is denoted IV, 9 as IX, 40 as XL, etc.

The following table gives the Roman numerals for the first few positive integers.

|  1  |  I   | 11  |  XI   | 21  |  XXI   |
| :-: | :--: | :-: | :---: | :-: | :----: |
|  2  |  II  | 12  |  XII  | 22  |  XXII  |
|  3  | III  | 13  | XIII  | 23  | XXIII  |
|  4  |  IV  | 14  |  XIV  | 24  |  XXIV  |
|  5  |  V   | 15  |  XV   | 25  |  XXV   |
|  6  |  VI  | 16  |  XVI  | 26  |  XXVI  |
|  7  | VII  | 17  | XVII  | 27  | XXVII  |
|  8  | VIII | 18  | XVIII | 28  | XXVIII |
|  9  |  IX  | 19  |  XIX  | 29  |  XXIX  |
| 10  |  X   | 20  |  XX   | 30  |  XXX   |

## Usage

```console
$ go run . hello
ERROR: cannot convert to roman digit
$ go run . 123
C+X+X+I+I+I
CXXIII
$ go run . 999
(M-C)+(C-X)+(X-I)
CMXCIX
$ go run . 3999
M+M+M+(M-C)+(C-X)+(X-I)
MMMCMXCIX
$ go run . 4000
ERROR: cannot convert to roman digit
$
```
