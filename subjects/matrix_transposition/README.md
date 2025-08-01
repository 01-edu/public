## matrix_transposition

### Instructions

- Define a `struct` named `Matrix` as a tuple of 2 tuples. The nested tuple will contain 2 `i32`s.

- Create a **function** named `transpose` that calculates the transposition of a 2x2 matrix.

```rust
pub fn transpose(m: Matrix) -> Matrix {
}
```

The transposition of a matrix, switches the columns to rows, and the rows to columns. For example:

```
( a b )   __ transposition __>   ( a c )
( c d )                          ( b d )
```

`Matrix` must implement `Debug`, `PartialEq` and `Eq`. You can use `derive`.

> Remember that you are defining a library, so any element that can be called from an external crate must be made public.

### Usage

Here is a possible program to test your function

```rust
use matrix_transposition::*;

fn main() {
    let matrix = Matrix((1, 3), (4, 5));
    println!("Original matrix {:?}", matrix);
    println!("Transpose matrix {:?}", transpose(matrix));
}
```

And its output:

```console
$ cargo run
Original matrix Matrix((1, 3), (4, 5))
Transpose matrix Matrix((1, 4), (3, 5))
$
```

### Notions

- [Defining a struct](https://doc.rust-lang.org/stable/book/ch05-01-defining-structs.html)

- [The Tuple Type](https://doc.rust-lang.org/stable/book/ch03-02-data-types.html#the-tuple-type)

- [Tuples](https://doc.rust-lang.org/rust-by-example/primitives/tuples.html)

- [Tuple Structs without Named Fields](https://doc.rust-lang.org/stable/book/ch05-01-defining-structs.html#using-tuple-structs-without-named-fields-to-create-different-types)

- [Adding Useful Functionality with Derived Traits](https://doc.rust-lang.org/stable/book/ch05-02-example-structs.html#adding-useful-functionality-with-derived-traits)

- [Referring to an Item in the Module Tree](https://doc.rust-lang.org/stable/book/ch07-03-paths-for-referring-to-an-item-in-the-module-tree.html)
