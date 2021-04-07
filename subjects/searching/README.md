## searching

### Instructions

In this exercise you will have to complete the function `search`.
This **function** receives an array and a key of `i32`, then it will return the position
of the given key in the array.
Only arrays with uniques keys will be tested.

### Notions

- [patterns](https://doc.rust-lang.org/book/ch18-00-patterns.html)

### Expected functions

```rust
pub fn search(array: &[i32], key: i32) -> Option<usize> {

}
```

### Usage

Here is a program to test your function.

```rust
use searching::*;

fn main() {
    let ar = [1, 3, 4, 6, 8, 9, 11];
    let f = search(&ar, 6);
    println!(
        "the element 6 is in the position {:?} in the array {:?}",
        f, ar
    );
}
```

And its output:

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
the element 6 is in the position Some(3) in the array [1, 3, 4, 6, 8, 9, 11]
student@ubuntu:~/[[ROOT]]/test$
```
