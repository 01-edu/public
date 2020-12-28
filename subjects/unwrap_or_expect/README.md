## unwrap

### Instructions

It will be given a function called **odd_to_even**, that returns an `Result`. If its an error it will
return a tuple with a string, indicating the error, and a vector with the elements that justifies the error

The objective is to execute the `odd_to_even` function and handle the error given by it

Create the following functions that receives a vector :

- `expect` that returns the error adding the sting "ERROR "
- `unwrap_or` that in case of error returns an empty vector
- `unwrap_err` that returns error if its `Ok` and returns the
   string containing the error in case of `Err`
- `unwrap` that unwraps the `Result`
- `unwrap_or_else` that in case of error returns a the vector that justifies the error

### Expected Function

```rust
pub fn odd_to_even(data: Vec<u32>) -> Result<Vec<u32>, (String, Vec<u32>)> {}
pub fn expect(v: Vec<u32>) -> Vec<u32> {}
pub fn unwrap_or(v: Vec<u32>) -> Vec<u32> {}
pub fn unwrap_err(v: Vec<u32>) -> (String, Vec<u32>) {}
pub fn unwrap(v: Vec<u32>) -> Vec<u32> {}
pub fn unwrap_or_else(v: Vec<u32>) -> Vec<u32> {}
```

### Usage

Here is a program to test your function

```rust
fn main() {
    // this will give an expect error
    // println!("{:?}", expect(vec![1, 3, 2, 5]));
    println!("{:?}", unwrap_or(vec![1, 3, 2, 5]));
    println!("{:?}", unwrap_or(vec![1, 3, 5]));
    println!("{:?}", unwrap_err(vec![1, 3, 2, 5]));
    // this will give an error that is unwraped
    // println!("{:?}", unwrap_err(vec![1, 3, 5]));
    println!("{:?}", unwrap(vec![1, 3, 5]));
    // this will give an error
    // println!("{:?}", unwrap(vec![1, 3, 2, 5]));
    println!("{:?}", unwrap_or_else(vec![1, 3, 5]));
    println!("{:?}", unwrap_or_else(vec![3, 2, 6, 5]));
}
```

And its output:

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
[]
[2, 4, 6]
("There is a even value in the vector!", [2])
[2, 4, 6]
[2, 4, 6]
[2, 6]
student@ubuntu:~/[[ROOT]]/test$
```

### Notions

- https://doc.rust-lang.org/std/?search=unwrap
