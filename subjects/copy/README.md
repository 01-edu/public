## copy

### Instructions

Your objective is to fix the **functions** work so that the program works.

- `nbr_function` returns a tuple:
  - with the `original` value
  - the `exponential function` of the value
  - and the `natural logarithm` of this `absolute` value
- `str_function` returns a tuple:
  - with the `original` value
  - and the `exponential function` each value as a string (see the example)
- `vec_function` returns a tuple:
  - with the `original` value
  - and the `natural logarithm` of each `absolute` value

The objective is to know how ownership works with different types.

### Notions

- [scope](https://doc.rust-lang.org/rust-by-example/scope/move.html)
- [Primitive Type f64](https://doc.rust-lang.org/std/primitive.f64.html)

### Expected Functions

```rust
pub fn nbr_function(c: u32) -> (u32, f64, f64) {
}

pub fn str_function(a: String) -> (String, String) {
}

pub fn vec_function(b: Vec<u32>) -> (Vec<u32>, Vec<f64>) {
}
```

### Usage

Here is a possible program to test your function :

```rust
use copy::*;

fn main() {
    let a: u32 = 0;
    let b = String::from("1 2 4 5 6");
    let c = vec![1, 2, 4, 5];

    println!("{:?}", nbr_function(a));
    // output : (0, 1.0, inf)
    println!("{:?}", str_function(b));
    // output : ("1 2 4 5 6", "2.718281828459045 7.38905609893065 54.598150033144236 148.4131591025766 403.4287934927351")
    println!("{:?}", vec_function(c));
    //  output : ([1, 2, 4, 5], [0.0, 0.6931471805599453, 1.3862943611198906, 1.6094379124341003])
}

```

And its output:

```console
$ cargo run
(0, 1.0, inf)
("1 2 4 5 6", "2.718281828459045 7.38905609893065 54.598150033144236 148.4131591025766 403.4287934927351")
([1, 2, 4, 5], [0.0, 0.6931471805599453, 1.3862943611198906, 1.6094379124341003])
$
```
