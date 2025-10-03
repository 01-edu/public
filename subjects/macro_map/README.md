## macro_map

### Instructions

Create a macro rule called `hash_map!` to initialize and optionally populate `std::collections::HashMap`, similarly to what `vec!` does for `Vec`.

> Your macro should work with or without a leading comma as do the language's list standard macros work.

> Your macro should also be able to work without the need to explicitly import `std::collections::HashMap` outside of its definition.

### Expected Macro

```rust
macro_rules! hash_map {
}
```

### Usage

Here is a possible program to test your function,

```rust
use macro_map::hash_map;
use std::collections::HashMap;

fn main() {
    let empty: HashMap<u32, u32> = hash_map![];
    println!("{:?}", empty);
    println!("{:?}", hash_map!['a' => 22, 'b' => 1, 'c' => 10]);
    println!(
        "{:?}",
        hash_map![
            "first" => hash_map![
                "Rob" => 32.2,
                "Gen" => 44.1,
                "Chris" => 10.,
            ],
            "second" => hash_map![]
        ]
    );
}
```

And its output:

```console
$ cargo run
{}
{'b': 1, 'a': 22, 'c': 10}
{"first": {"Rob": 32.2, "Gen": 44.1, "Chris": 10.0}, "second": {}}
$
```
