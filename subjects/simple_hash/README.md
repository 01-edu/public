## simple_hash

### Instructions

Create a **function** `contain` that checks a `HashMap` to see if it contains a given key.

Create a **function** `remove` that removes a given key from the `HashMap`.

- Note: pay attention to the comment in the [usage](#usage)

### Notions

- [HashMap](https://doc.rust-lang.org/rust-by-example/std/hash.html)

### Expected functions

```rust
pub fn contain(h: &HashMap<&str, i32>, s: &str) -> bool {}

pub fn remove(h: &mut HashMap<&str, i32>, s: &str) {}
```

### Usage

Here is a program to test your function.

```rust
use simple_hash::*;
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
        //----------^^^^^^^^
        // this is not correct, fix it to match the solution the expected function
    );
    println!(
        "Does the HashMap contains the name Katie? => {}",
        contain(hash.clone(), "Katie")
        //----------^^^^^^^^
        // this is not correct, fix it to match the solution the expected function
    );
    println!("Removing Robert {:?}", remove(hash.clone(), "Robert"));
    println!(
        "Does the HashMap contains the name Robert? => {}",
        contain(hash.clone(), "Robert")
        //----------^^^^^^^^
        // this is not correct, fix it to match the solution the expected function
    );
    println!("Hash {:?}", &hash);
}
```

And its output

```console
$ cargo run
Does the HashMap contains the name Roman? => false
Does the HashMap contains the name Katie? => true
Removing Robert ()
Does the HashMap contains the name Robert? => false
Hash {"Katie": 334, "Daniel": 122, "Ashley": 333}
$
```
