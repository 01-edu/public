## bigger

### Instructions
Create the function `bigger` that gets the biggest positive number in the `HashMap`

### Notions

- https://doc.rust-lang.org/stable/std/collections/struct.HashMap.html

### Expected functions

```rust
fn bigger(h: HashMap<&str, i32>) -> i32{}
```

### Usage

Here is a program to test your function.

```rust
fn main() {
    let mut hash = HashMap::new();
    hash.insert("Daniel", 122);
    hash.insert("Ashley", 333);
    hash.insert("Katie", 334);
    hash.insert("Robert", 14);

    println!(
        "The biggest of the elements in the HashMap is {}",
        bigger(hash)
    );
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
The biggest of the elements in the HashMap is 334
student@ubuntu:~/[[ROOT]]/test$
```
