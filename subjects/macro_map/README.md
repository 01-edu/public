## macro_map

### Instructions

Create a macro rule called `hash_map` to initialize and declare a `HashMap` at the same time, very similar to what `vec!` macro does for `Vector`.

> Your macro should accept both leading and non leading commas syntax to be more flexible in terms of coding style and reflect the language general style.

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
    let empty: HashMap<u32, u32> = hash_map!();
    let new = hash_map!('a' => 22, 'b' => 1, 'c' => 10);
    let nested = hash_map!(
        "first" => hash_map!(
            "Rob" => 32.2,
            "Gen" => 44.1,
            "Chris" => 10.,
        ),
        "second" => hash_map!()
    );
    println!("{:?}", empty);
    println!("{:?}", new);
    println!("{:?}", nested);
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

