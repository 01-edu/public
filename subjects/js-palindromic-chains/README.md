## Palindromic Chains

### Instructions

Write a function `palindromicChain(numbers)` that takes an array of integers, checks if each number is a palindrome, repeatedly adds it to its reverse until it becomes a palindrome or reaches 100 attempts (returning 0 if unsuccessful), and returns an array of the resulting palindromes.

### Expected function

```js
function palindromicChain(numbers) {}
```

### Usage

Here is a possible program to test your function:

```js
console.log(palindromicChain([87, 33, 123, 196]));
```

And its output:

```console
$ node main.js
[4884, 33, 444, 0]
$
```

### Explanation:

For 87: 87 + 78 = 165; 165 + 561 = 726; 726 + 627 = 1353; 1353 + 3531 = 4884.
For 12: 12 + 21 = 33.
For 123: 123 + 321 = 444.
