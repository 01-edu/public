## uniqueoccurences

### Instructions

Write a program that outputs `true` if the number of occurrences of each character is unique, otherwise `false`. `\n` should be at the of line.

If number of arguments is not 1 output `\n`.

Only lower case characters will be given.

### Usage

```console
$> go build
$> ./main "abbaac"
true
$> ./main "ab"
false
$> ./main "abcacccazb"
true
```

In first example, character 'a' has 3 occurrences, 'b' has 2 and 'c' has 1. No two characters have the same number of occurrences.

