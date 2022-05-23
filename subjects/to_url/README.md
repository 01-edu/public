## to_url

### Instructions

Create a **function** named `to_url` which takes a string and substitutes every white-space with `"%20"`.

### Expected functions
```rust
pub fn to_url(s: &str) -> String {
}
```

### Usage

Here is a program to test your function.

```rust
use to_url::*;

fn main() {
	let s = "Hello, world!";
	println!("{} to be use as an url is {}", s, to_url(s));
}
```

And its output

```console
$ cargo run
Hello, world! to be use as an url is Hello,%20world!
$
```
