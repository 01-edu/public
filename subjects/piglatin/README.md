## piglatin

### Instructions

Write a **program** that transforms a string passed as argument in its `Pig Latin` version.

The rules used by Pig Latin are as follows:

- If a word begins with a vowel, just add "ay" to the end.
- If it begins with a consonant, then we take all consonants before the first vowel and we put them on the end of the word and add "ay" at the end.
- Only the latin vowels will be considered as vowel(aeiou).

If the word has no vowels, the program should print "No vowels".

If the number of arguments is different from one, the program prints nothing.

### Usage

```console
$ go run .
$ go run . pig | cat -e
igpay$
$ go run . Is | cat -e
Isay$
$ go run . crunch | cat -e
unchcray$
$ go run . crnch | cat -e
No vowels$
$ go run . something else | cat -e
$
```
