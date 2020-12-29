## get_products

### Instructions

Create a function `get_products` that takes a vector of integers, and returns a vector of the products
of each index. For this exercise to be correct you will have to return the product of every index
except the current one.

### Notions

- https://doc.rust-lang.org/std/iter/trait.Iterator.html

### Expected functions

```rust
fn get_products(arr: Vec<usize>) -> Vec<usize> {}
```

### Usage

Here is a program to test your function.

```rust
fn main() {
    let arr: Vec<usize> = vec![1, 7, 3, 4];
    let output = get_products(arr);
    println!("{:?}", output);
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
[84, 12, 28, 21]
student@ubuntu:~/[[ROOT]]/test$
```
