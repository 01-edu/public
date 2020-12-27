## simple_hash

### Instructions

- Create the function `contain` that checks a `HashMap` to see if it contains the given key.

- Create the function `remove` that removes a given key from the `HashMap`.

### Expected Functions

```rust
fn contain(h: HashMap<&str, i32>, s: &str) -> bool {
}

fn remove(mut h: HashMap<&str, i32>, s: &str) {
}
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
        contain(hash.clone(), "Roman")
    );
    println!(
        "Does the HashMap contains the name Katie? => {}",
        contain(hash.clone(), "Katie")
    );
    println!("Removing Robert {:?}", remove(hash.clone(), "Robert"));
    println!("Hash {:?}", hash);
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
Does the HashMap contains the name Roman? => false
Does the HashMap contains the name Katie? => true
Removing Robert ()
Hash {"Daniel": 122, "Ashley": 333, "Robert": 14, "Katie": 334}
student@ubuntu:~/[[ROOT]]/test$
```
