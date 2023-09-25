## is_anagram

### Instructions

Write a function called `is_anagram` that checks if one string is an anagram of another string. An anagram is a word or phrase formed by rearranging the letters of another, such as "listen" and "silent."

```rust
pub fn is_anagram(s1: &str, s2: &str) -> bool {
	// Your code goes here
}
```

- `s1: &str`: The first input string.
- `s2: &str`: The second input string.

The function should return `true` if `s1` is an anagram of `s2`, and `false` otherwise.
Your task is to implement the `is_anagram` function to determine whether the two input strings are anagrams of each other.

### Usage

Here is a possible runner to test your function:

```rust
use is_anagram::is_anagram;

fn main() {
    let s1 = "listen";
    let s2 = "silent";
    
    if is_anagram(s1, s2) {
        println!("{} and {} are anagrams!", s1, s2);
    } else {
        println!("{} and {} are not anagrams.", s1, s2);
    }
}
```

And its output:

```console
$ cargo run
listen and silent are anagrams!
```
