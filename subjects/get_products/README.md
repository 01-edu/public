## get_products

### Instructions

Create a function named `get_products` that takes a vector of integers, and returns a vector of the products of each index.

You'll need to return the product of every index
except the current one.

### Example:
For `[1,2,3,4]`, we get:

- for the number `1` we get `2*3*4 = 24`.
- for the number `3` we get `1*2*4 = 8`.

### Expected functions

```rust
pub fn get_products(arr: Vec<usize>) -> Vec<usize> {

}
```

### Usage

Here is a program to test your function.

```rust
use get_products::get_products;

fn main() {
    let arr: Vec<usize> = vec![1, 7, 3, 4];
    let output = get_products(arr);
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
