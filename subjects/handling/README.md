## handling

### Instructions

Write a **function**, called `open_or_create` which has two arguments:

- `file : &str` which is the name of the file
- `content: &str` which will be the content to be written into the file

This function should try to open a file, if it does not exist, it should create it.
In case something goes wrong, it should panic, with the error.

### Notions

- [Error kind](https://doc.rust-lang.org/std/io/enum.ErrorKind.html)
- [struct file](https://doc.rust-lang.org/std/fs/struct.File.html)
- [OPenOptions](https://doc.rust-lang.org/std/fs/struct.OpenOptions.html)

### Expected Function

```rust
pub fn open_or_create(file: &str, content: &str) {

}
```

### Usage

Here is a program to test your function

```rust
use std::fs::File;
use std::io::Read;
use handling::*;

fn main() {
    let path = "a.txt";
    File::create(path).unwrap();
    open_or_create(path, "content to be written");

    let mut file = File::open(path).unwrap();

    let mut s = String::new();
    file.read_to_string(&mut s).unwrap();
    println!("{}", s);
}
```

And its output:

```console
$ cargo run
content to be written
$
```
