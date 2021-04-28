## nenokku

### Instructions

You are given unknown amount of operations. There are 3 types of operations, you should handle each appropriately:

1. "A word" - add the "word" into your "dictionary" of words.
2. "? word" - check if the "word" is in your "dictionary".
3. "x word" - end of operations.

Intersection of words is counted as if it is in your dictionary (see example).
Write a program that takes as arguments operations.

```console
$ go run . "? love" "? is" "A loveis" "? love" "? Who" "A Whoareyou" "? is"
NO
NO
YES
NO
YES
$
```
