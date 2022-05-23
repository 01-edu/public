## name_initials

### Instructions

Create a **function** called `initials`, this function will receive a vector of string literals
with names and return a vector of Strings with the initials of each name.

```rust
pub fn initials(names: Vec<&str>) -> Vec<String> {
}
```

> This exercise will test how many times the **heap is going to be allocated**!\
> So try your best to allocate the minimum data on the heap!


### Usage

Here is a program to test your function:

```rust
use name_initials::initials;

fn main() {
    let names = vec!["Harry Potter", "Someone Else", "J. L.", "Barack Obama"];
    println!("{:?}", initials(names));
}
```

And its output

```console
$ cargo run
["H. P.", "S. E.", "J. L.", "B. O."]
$
```

### Notions

- [stack and heap](https://doc.rust-lang.org/1.22.0/book/first-edition/the-stack-and-the-heap.html)
