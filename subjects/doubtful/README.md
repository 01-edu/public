## doubtful

### Instructions

Write a function called `doubtful` that adds to every string passed to it a question mark (?).

You have to complete the signature the functions to make it fit the usage and you also have to modify the usage to be able to do what is expected.

- Restrictions:

  - `doubtful` cannot return any value.

  - And in the example of the usage you can only modify the argument of the function

### Expected

```rust
pub fn doubtful(s: /*give the correct type*/ ) {
}
```

### Usage

Here is a program to test your function

```rust
fn main() {
	let mut s = String::from("Hello");

	println!("Before changing the string: {}", s);

	doubtful(/*add your code here*/);

	println!("After changing the string: {}", s);
}
```

And its output

```console
student@ubuntu:~/doubtful/test$ cargo run
Before changing the string: Hello
After changing the string: Hello?
student@ubuntu:~/doubtful/test$
```
