## stars

### Instructions

Write a function named `stars` that takes a number as
parameter and returns a string of stars (asterisks) 2^n long (2 to the nth power).

### Expected functions

```rust
fn stars(n: u32) -> String {}
```

### Usage

Here is a program to test your function.

```rust
fn main() {
    println!("{}", stars(1));
    println!("{}", stars(4));
    println!("{}", stars(5));
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
**
****************
********************************
student@ubuntu:~/[[ROOT]]/test$
```
