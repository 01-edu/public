## panic

### Instructions

Write a function that tries to open a file and panics if the file
doesn't exist

### Expected Function

```rust
pub fn open_file(s: &str) -> File {}
```

### Usage

Here is a program to test your function

```rust
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
student@ubuntu:~/[[ROOT]]/test$ cargo run
File { fd: 3, path: "[[ROOT]]/a.txt", read: true, write: false }
student@ubuntu:~/[[ROOT]]/test$
```
