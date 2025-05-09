## collect

### Instructions

Implement the **function** `bubble_sort`, which receives a list of integers and sorts it in increasing order using the bubble sort algorithm.

### Expected Function

```rust
pub fn bubble_sort(arr: &mut [i32]) {
}
```

### Usage

Here is a program to test your function.

```rust
use collect::*;

fn main() {
    let mut v = [3, 2, 4, 5, 1, 7];
    let mut v_clone = v;

    bubble_sort(&mut v);
    println!("{:?}", v);

    v_clone.sort_unstable();
    println!("{:?}", v_clone);
}
```

And its output:

```console
$ cargo run
[1, 2, 3, 4, 5, 7]
[1, 2, 3, 4, 5, 7]
$
```
