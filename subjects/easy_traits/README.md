## easy_traits

### Instructions

Your task is to implement the trait `AppendStrExt` for the type `String`.

The trait `AppendStrExt` has the following functions:

- `append_str`: that appends `str_to_append` to `self` and returns it.
- `append_number`: that appends `nb_to_append` to `self` and returns it.
- `remove_punctuation_marks`: that removes all punctuation from `self`.

> **Note: For the sake of this exercise, we only consider `.`, `,`, `?` and `!` as punctuation.**

### Expected Function

```rust
pub trait AppendStrExt {
    fn append_str(&mut self, str_to_append: &str) -> &mut Self;

    fn append_number(&mut self, nb_to_append: f64) -> &mut Self;

    fn remove_punctuation_marks(&mut self) -> &mut Self;
}

impl AppendStrExt for String {
    fn append_str(&mut self, str_to_append: &str) -> &mut Self {
        todo!()
    }

    fn append_number(&mut self, nb_to_append: f64) -> &mut Self {
        todo!()
    }

    fn remove_punctuation_marks(&mut self) -> &mut Self {
        todo!()
    }
}
```

### Usage

Here is a program to test your function.

```rust
use easy_traits::*;

fn main() {
    let mut s = "hello".to_owned();

    println!("Before append: {}", s);

    s.append_str(" there!");
    println!("After append: {}", s);

    s.remove_punctuation_marks();
    println!("After removing punctuation: {}", s);
}
```

And its output

```console
$ cargo run
Before append: hello
After append: hello there!
After removing punctuation: hello there
$
```
