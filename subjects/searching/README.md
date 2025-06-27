## searching

### Instructions

Complete the function `search`. It should return the index of the **last** element which matches `key` in the array.

### Expected functions

```rust
pub fn search(array: &[i32], key: i32) -> Option<usize> {
    todo!()
}
```

### Usage

Here is a program to test your function.

```rust
use searching::*;

fn main() {
    let ar = [1, 3, 4, 6, 8, 9, 11, 8];
    let f = search(&ar, 8);
    println!(
        "the element 8 is last in the position {:?} in the array {:?}",
        f, ar
    );
}
```

And its output:

```console
$ cargo run
the element 8 is last in the position Some(7) in the array [1, 3, 4, 6, 8, 9, 11, 8]
$
```

### Notions

- [patterns](https://doc.rust-lang.org/book/ch18-00-patterns.html)
