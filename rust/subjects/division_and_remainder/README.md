## division_and_remainder

### Instructions

Complete the function `divide` to return the value of the integer division of the x and y and the remainder of that division.
You're only allowed to change the body of the function

### Notions

- https://doc.rust-lang.org/rust-by-example/primitives/tuples.html

### Expected functions

```rust
fn divide(x: i32, y: i32) -> (i32, i32) {}
```

### Usage

Here is a program to test your function.

```rust

fn main() {
    let x = 9;
    let y = 4;
    let (division, remainder) = divide(x, y);
    println!(
        "{}/{}: division = {}, remainder = {}",
        x, y, division, remainder
    );
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
9/4: division = 2, remainder = 1
student@ubuntu:~/[[ROOT]]/test$
```
