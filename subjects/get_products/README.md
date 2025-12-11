## get_products

### Instructions

Create a function named `get_products` that takes an array of integers, and returns an array of the products of each other index.

For each element, you'll need to return the product of every other index.

If an array of one element is given, the same element should always become `1`.

### Example:

For `[1,2,3,4]`, we get:

- The element `1` becomes `2*3*4 = 24`.
- The element `3` becomes `1*2*4 = 8`.

### Expected functions

```rust
pub fn get_products(arr: [usize; _]) -> [usize; _] {
    todo!()
}
```

### Usage

Here is a program to test your function.

```rust
use get_products::get_products;

fn main() {
    let output = get_products([1, 7, 3, 4]);
    println!("{:?}", output);
}
```

And its output:

```console
$ cargo run
[84, 12, 28, 21]
$
```

### Notions

- [Trait iterator](https://doc.rust-lang.org/std/iter/trait.Iterator.html)
