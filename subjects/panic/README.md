## panic

### Instructions

Create a **function** that tries to open a file and panics if the file does not exist.

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

    //It must panic
    let b = open_file(filename);
}
```

And its output:

```console
$ cargo run
File { fd: 3, path: "panic/a.txt", read: true, write: false }
thread 'main' panicked at 'File not found'
$
```
