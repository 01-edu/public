## panic

### Instructions

Write a **function** that tries to open a file and panics if the file
does not exist.

### Notions

### Expected Function

```rust
pub fn open_file(s: &str) -> File {

}
```

### Usage

Here is a program to test your function:

```rust
use std::fs::File;
use std::fs;
use panic::*;

fn main() {
    let filename = "created.txt";
    File::create(filename).unwrap();
    let a = open_file(filename);
    println!("{:?}", a);
    fs::remove_file(filename).unwrap();
}
```

And its output:

```console
$ cargo run
File { fd: 3, path: "panic/a.txt", read: true, write: false }
$
```
