## reverse_it

### Instructions

Create a function called `reverse_it` that takes a number and returns a string with the number backwards followed by the original number. If the number is negative a char `-` has to be added to the beginning of the string.

### Expected Functions

```rust
pub fn reverse_it(v: i32) -> String {

}
```

### Usage

Here is a program to test your function,

```rust
fn main() {
    println!("{}", reverse_it(123));
    println!("{}", reverse_it(-123));
}
```

And its output:

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
321123
-321123
student@ubuntu:~/[[ROOT]]/test$
```
