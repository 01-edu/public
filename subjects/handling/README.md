## handling

### Instructions

Create a **function** `open_or_create` with two arguments:

- `file`, which represents a file path.
- `content`, which represents the content to be written to the file.

This function should try to open the file. If it does not exist, the file should be created.
The content should be appended to the file. This means it shouldn't replace whatever content is inside, but append to it.
Should anything go wrong, the function should panic.

### Expected Function

```rust
use std::path::Path;

pub fn open_or_create<P: AsRef<Path>>(path: &P, content: &str) {
    todo!()
}
```

### Usage

Here is a program to test your function

```rust
use std::fs;

fn main() {
    let path = "a.txt";

    handling::open_or_create(&path, "content to be written");

    let contents = fs::read_to_string(path).unwrap();
    println!("{}", contents);
}
```

And its output:

```console
$ cargo run
content to be written
$
```

### Notions

- [ErrorKind](https://doc.rust-lang.org/std/io/enum.ErrorKind.html)
- [File](https://doc.rust-lang.org/std/fs/struct.File.html)
- [OpenOptions](https://doc.rust-lang.org/std/fs/struct.OpenOptions.html)
