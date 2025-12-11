## slices_to_map

### Instructions:

Create a function that borrows two slices and returns a hashmap where the first slice represents the keys and the second represents the values.

- If the slices have different sizes, the function should return the hashmap with the size of the smallest list.

### Expected Function

```rust
pub fn slices_to_map(keys: &[T], values: &[U]) -> HashMap<T, U> {
    todo!()
}
```

### Usage

Here is a program to test your function.

```rust
use slices_to_map::*;

fn main() {
    let keys = ["Olivia", "Liam", "Emma", "Noah", "James"];
    let values = [1, 3, 23, 5, 2];
    println!("{:?}", slices_to_map(&keys, &values));
}
```

And its output

```console
$ cargo run
{"James": 2, "Liam": 3, "Emma": 23, "Noah": 5, "Olivia": 1}
$
```
