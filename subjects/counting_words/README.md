## counting_words

### Instructions

Create a function named `counting_words`, that receives a `&str`. It should return each word and the number of times it appears on the string.

Each of the following will count as **1** word:

- A number like "0" or "1234".
- A word or letter like "a" or "they".
- Two words joined by a single apostrophe like "it's" or "they're").

The function must respect the following rules:

- The count is case insensitive, so that "HELLO", "Hello", and "hello" are 3 uses of the same word.
- All forms of punctuation are to be ignored, except for the apostrophe if used like the example above.
- The words can be separated by any form of whitespace (ie "\t", "\n", " ").

### Expected Function

```rust
fn counting_words(words: &str) -> HashMap<String, u32> {
}
```

### Usage

Here is a possible program to test your function :

```rust
use counting_words::counting_words;
use std::collections::HashMap;

fn main() {
    println!("{:?}", counting_words("Hello, world!"));
    println!("{:?}", counting_words("“Two things are infinite: the universe and human stupidity; and I'm not sure about the universe.”
    ― Albert Einstein "));
    println!("{:?}", counting_words("Batman, BATMAN, batman, Stop stop"));
}
```

And its output:

```console
$ cargo run
{"hello": 1, "world": 1}
{"and": 2, "human": 1, "universe": 2, "the": 2, "i\'m": 1, "about": 1, "einstein": 1, "are": 1, "infinite": 1, "sure": 1, "albert": 1, "two": 1, "things": 1, "not": 1, "stupidity": 1}
{"batman": 3, "stop": 2}
$
```
