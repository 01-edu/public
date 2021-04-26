## unwrap_or_expect

### Instructions

A **function** called **odd_to_even** will be given, which returns a `Result`. If an error occurs the function will
return a tuple with a string, stating the error, and a vector with the elements which causing the error.

The objective is to execute the `odd_to_even` function and handle the error returned by it.

Create the following functions which receives a vector :

- `expect` which returns the error adding the string "ERROR "
- `unwrap_or` which in case of error returns an empty vector
- `unwrap_err` which returns error if its `Ok` and returns the
  string containing the error in case of `Err`
- `unwrap` which unwraps the `Result`
- `unwrap_or_else` which in case of error returns the vector of elements which causes the error

### Notions

- [Error Handling](https://doc.rust-lang.org/book/ch09-00-error-handling.html)
- [Unwrap keywords](https://doc.rust-lang.org/std/?search=unwrap)

### Expected Functions

```rust
pub fn odd_to_even(data: Vec<u32>) -> Result<Vec<u32>, (String, Vec<u32>)> {
    let mut a = Vec::new();
    a.extend(data.iter().filter(|&value| value % 2 == 0));
    if a.len() != 0 {
        return Err(("There is a even value in the vector!".to_string(), a));
    }
    a.extend(data.iter().map(|&value| {
        value + 1
    }));
    Ok(a)
}
pub fn expect(v: Vec<u32>) -> Vec<u32> {

}
pub fn unwrap_or(v: Vec<u32>) -> Vec<u32> {

}
pub fn unwrap_err(v: Vec<u32>) -> (String, Vec<u32>) {

}
pub fn unwrap(v: Vec<u32>) -> Vec<u32> {

}
pub fn unwrap_or_else(v: Vec<u32>) -> Vec<u32> {

}
```

### Usage

Here is a program to test your function:

```rust
use unwrap_or_expect::*;

fn main() {
    // // if uncommented, the below line will give an expect "ERROR "
    // println!("{:?}", expect(vec![1, 3, 2, 5]));

    println!("{:?}", unwrap_or(vec![1, 3, 2, 5]));
    println!("{:?}", unwrap_or(vec![1, 3, 5]));

    println!("{:?}", unwrap_err(vec![1, 3, 2, 5]));

    // // if uncommented, the below line will give an unwraped error
    // println!("{:?}", unwrap_err(vec![1, 3, 5]));

    println!("{:?}", unwrap(vec![1, 3, 5]));

    //// if uncommented, the below line will give an error
    // println!("{:?}", unwrap(vec![1, 3, 2, 5]));

    println!("{:?}", unwrap_or_else(vec![1, 3, 5]));
    println!("{:?}", unwrap_or_else(vec![3, 2, 6, 5]));
}
```

And its output:

```console
[]
[2, 4, 6]
("There is a even value in the vector!", [2])
[2, 4, 6]
Ok([2, 4, 6])
[2, 4, 6]
Err(("There is a even value in the vector!", [2, 6]))
[2, 6]
student@ubuntu:~/unwrap_or_expect/test$
```
