## edit_distance

### Instructions

Create a **function** named `edit_distance`, which calculates the minimum number of changes (insertions, deletions and/or substitutions) which are needed to transform the `source` string to the `target` string.

### Expected Function

```rust
pub fn edit_distance(source: &str, target: &str) -> usize {
}
```

### Usage

Here is a program to test your function.

```rust
use edit_distance::*;

fn main() {
    let source = "alignment";
    let target = "assignment";

    println!(
        "It's necessary to make {} change(s) to {:?} to get {:?}",
        edit_distance(source, target),
        source,
        target
    );
}
```

And its output:

```console
$ cargo run
It's necessary to make 2 change(s) to "alignment" to get "assignment"
$
```

### Notions

[Edit Distance (Wikipedia)](https://en.wikipedia.org/wiki/Edit_distance)
