## handling

### Instructions

Write a function, called `open_or_create` that as two arguments:

- `file : &str` which is the name of the files
- `content: &str` being the content to be written into the file

This functions should try to open a file, if it does not exist creates it.
You should panic, with the error, in case something goes wrong.

### Expected Function

```rust
pub fn open_or_create(s: &str, content: &str) {}
```

### Usage

Here is a program to test your function

```rust
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
student@ubuntu:~/[[ROOT]]/test$ cargo run
content to be written
student@ubuntu:~/[[ROOT]]/test$
```

### Notions

- https://doc.rust-lang.org/std/io/enum.ErrorKind.html
- https://doc.rust-lang.org/std/fs/struct.File.html
