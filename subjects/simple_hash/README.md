## simple_hash

### Instructions

Create a **function** named `word_frequency_counter` which will receive a vector of strings (each string being a single word) and return an `HashMap` with the word as the key and the number of repetitions as the value.

Also create a function named `nb_distinct_words` which will take a reference to the `HashMap` and return the number of distinct words present in it.

> all the words tested will be lowercase

### Expected functions

```rust
pub fn word_frequency_counter(words: Vec<&str>) -> HashMap<&str, usize> {}

pub fn nb_distinct_words(frequency_count: &HashMap<&str, usize>) -> usize {}
```

### Usage

Here is a program to test your function.

```rust
use simple_hash::*;

const SENTENCE: &str = "this is a very basic sentence with only a few repetitions. once again this is very basic but it should be enough for basic tests";

fn main() {
    let words = SENTENCE.split_ascii_whitespace().collect::<Vec<_>>();
    let frequency_count = word_frequency_counter(&words);

    println!("{:?}", frequency_count);
    println!("{}", nb_distinct_words(&frequency_count));
}
```

And its output

```console
$ cargo run
{"tests": 1, "with": 1, "this": 2, "it": 1, "enough": 1, "is": 2, "but": 1, "sentence": 1, "only": 1, "basic": 3, "again": 1, "for": 1, "be": 1, "once": 1, "very": 2, "should": 1, "few": 1, "a": 2, "repetitions.": 1}
20
$
```

### Notions

- [HashMap](https://doc.rust-lang.org/rust-by-example/std/hash.html)
