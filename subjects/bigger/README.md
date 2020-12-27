## bigger

### Instructions

Create the function `bigger` that gets the biggest positive number in the `HashMap`.

### Expected Function

```rust
fn bigger(h: HashMap<&str, i32>) -> i32 {
}
```

### Usage

Here is a program to test your function.

```rust
use bigger::bigger;
fn main() {
    
    let mut hash = HashMap::new();
    hash.insert("Daniel", 122);
    hash.insert("Ashley", 333);
    hash.insert("Katie", 334);
    hash.insert("Robert", 14);

    println!("The biggest of the elements in the HashMap is {}", bigger(hash));
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
Is `thought` a permutation of `thougth`? = true
student@ubuntu:~/[[ROOT]]/test$
```
