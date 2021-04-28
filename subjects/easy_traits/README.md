## easy_traits

### Instructions

Your task is to implement the trait `AppendStr` for the type String, and `AppendVec` for a vector of strings.

The trait `AppendStr` has only one function, which is appending `"world"` to any object implementing this trait.
And `AppendVec` will append `"three"` to a vector of strings.

### Expected Function

```rust
trait AppendStr {
    fn append_str(self) -> Self;
}

trait AppendVec {
    fn append_vec(&mut self) -> Self;
}

impl AppendStr for String {}

impl AppendVec for Vec<String> {}
```

### Usage

Here is a program to test your function.

```rust
use easy_traits::easy_traits;

fn main() {
    let s = String::from("hello ").append_str();
    println!("{}", s);

    let l = vec![String::from("one"), String::from("two")].append_vec();
    println!("{:?}", l);
}
```

And its output

```console
$ cargo run
hello world
["one", "two", "three"]
$
```
