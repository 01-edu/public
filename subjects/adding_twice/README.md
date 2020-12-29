## adding_twice

### Instructions

In this exercise you will have to reuse your `add_curry` function
Then you have to complete the function `twice` using closures, this function will
take a function f(x) as parameter and return a function f(f(x))
So, the purpose of this function is to add two times the value in `add_curry` to the original value.

### Notions

- https://doc.rust-lang.org/rust-by-example/fn/hof.html#higher-order-functions

### Expected functions

The type of the arguments are missing use the example `main` function to determine the correct type.

```rust
fn twice<T>(F: _) -> _{}
```

### Usage

Here is a program to test your function.

```rust
fn main() {
    let add10 = add_curry(10);
    let value = twice(add10);
    println!("The value is {}", value(7));

    let add20 = add_curry(20);
    let value = twice(add20);
    println!("The value is {}", value(7));

    let add30 = add_curry(30);
    let value = twice(add30);
    println!("The value is {}", value(7));

    let neg = add_curry(-32);
    let value = twice(neg);
    println!("The value is {}", value(7));
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
The value is 27
The value is 47
The value is 67
The value is -57
student@ubuntu:~/[[ROOT]]/test$
```
