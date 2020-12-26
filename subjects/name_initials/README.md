## name_initials

### Instructions

Create a function called `initials`, this function will receive a vector of string literals
with names and return a vector of Strings with the initials of each name.

> This exercise will test the **heap allocation** of your function!
> So try your best to allocate the minimum data on the heap!

### Notions

- https://doc.rust-lang.org/1.22.0/book/first-edition/the-stack-and-the-heap.html

### Expected Function

```rust
fn initials(names: &mut Vec<&str>) -> Vec<String> {
}
```

### Usage

Here is a program to test your function

```rust
use name_initials::initials;

fn main() {
    let names = vec!["Harry Potter", "Someone Else", "J. L.", "Barack Obama"]
    println!("{:?}", initials(names));
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
["H. P.", "S. E.", "J. L.", "B. O."]
student@ubuntu:~/[[ROOT]]/test$
```
