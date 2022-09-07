## easy_traits

### Instructions

Your task is to implement the trait `AppendStr` for the type `StringValue`.

The trait `AppendStr` has the following functions:

- `append_str`: that appends the string to the `value`.
- `append_number`: that appends the number to the `value`.
- `remove_punctuation_marks`: that removes punctuation from the `value` (`.`, `,`, `?` and `!`).


### Expected Function

```rust
#[derive(Clone)]
struct StringValue {
    value: String,
}

trait AppendStr {
    fn append_str(self, new_str: String) -> Self;

    fn append_number(self, new_number: f64) -> Self;

    fn remove_punctuation_marks(self) -> Self;
}

impl AppendStr for StringValue {
}
```

### Usage

Here is a program to test your function.

```rust
use easy_traits::*;

fn main() {
    let mut str_aux = StringValue {
        value: String::from("hello"),
    };

    println!("Before append: {}", str_aux.value);

    str_aux.append_str(String::from(" there!"));
    println!("After append: {}", str_aux.value);

    str_aux.remove_punctuation_marks();
    println!("After removing punctuation: {}", str_aux.value);
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
