## find_factorial

### Instruccions

Complete the function `factorial` to return the factorial of a given number

### Expected Function

```rust
pub fn factorial(num: u64) -> u64 {
}
```

### Usage

Here is a possible program to test your function :

```rust
use find_factorial::factorial;

fn main() {
    println!("The factorial of 0 = {}", factorial(0));
    println!("The factorial of 1 = {}", factorial(1));
    println!("The factorial of 5 = {}", factorial(5));
    println!("The factorial of 10 = {}", factorial(10));
    println!("The factorial of 19 = {}", factorial(19));
}
```

And its output:

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
The factorial of 0 = 1
The factorial of 1 = 1
The factorial of 5 = 120
The factorial of 10 = 3628800
The factorial of 19 = 121645100408832000
student@ubuntu:~/[[ROOT]]/test$
```
