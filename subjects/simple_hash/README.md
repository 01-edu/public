## simple_hash

### Instructions

Create the function `contain` that checks a `HashMap` to see if it contains the given key.

Create the function `remove` that removes a given key from the `HashMap`.

### Notions

- https://doc.rust-lang.org/rust-by-example/std/hash.html

### Expected functions

```rust
pub fn contain(h: &HashMap<&str, i32>, s: &str) -> bool {}

pub fn remove(h: &mut HashMap<&str, i32>, s: &str) {}
```

### Usage

Here is a program to test your function.

```rust
use std::collections::HashMap;

fn main() {
    let mut hash: HashMap<&str, i32> = HashMap::new();
    hash.insert("Daniel", 122);
    hash.insert("Ashley", 333);
    hash.insert("Katie", 334);
    hash.insert("Robert", 14);

    println!(
        "Does the HashMap contains the name Roman? => {}",
        contain(&hash, "Roman")
    );
    println!(
        "Does the HashMap contains the name Katie? => {}",
        contain(&hash, "Katie")
    );
    println!("Removing Robert {:?}", remove(&mut hash, "Robert"));
    println!(
        "Does the HashMap contains the name Robert? => {}",
        contain(&hash, "Robert")
    );
    println!("Hash {:?}", hash);
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
Does the HashMap contains the name Roman? => false
Does the HashMap contains the name Katie? => true
Removing Robert ()
Does the HashMap contains the name Robert? => false
Hash {"Katie": 334, "Daniel": 122, "Ashley": 333}
student@ubuntu:~/[[ROOT]]/test$
```
