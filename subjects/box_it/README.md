## box_it

### Instructions

Create the following **functions**:

- `transform_and_save_on_heap`, which receives a string with a number (that can or not have a k (kilo) suffix)
  and transforms those numbers into `u32` and inserts them into a vector which must be saved in the heap using **Box**.

- `take_value_ownership`, which takes the value (unboxed value) from the box and returns it

### Notions

- [smart pointers](https://doc.rust-lang.org/book/ch15-00-smart-pointers.html)
- [using box](https://doc.rust-lang.org/book/ch15-01-box.html)

### Expected Functions

```rust
pub fn transform_and_save_on_heap(s: String) -> Box<Vec<u32>> {

}
pub fn take_value_ownership(a: Box<Vec<u32>>) -> Vec<u32> {

}
```

### Usage

Here is a program to test your functions

```rust
use box_it::*;

fn main() {
    let new_str = String::from("5.5k 8.9k 32");

    // creating a variable and we save it in the Heap
    let a_h = transform_and_save_on_heap(new_str);
    println!("Box value : {:?}", &a_h);
    println!("size occupied in the stack : {:?} bytes", (std::mem::size_of_val(&a_h)));

    let a_b_v = take_value_ownership(a_h);
    println!("value : {:?}", &a_b_v);
    println!("size occupied in the stack : {:?} bytes", (std::mem::size_of_val(&a_b_v)));
    // whenever the box, in this case "a_h", goes out of scope it will be deallocated, freed
}
```

And its output:

```console
$ cargo run
Box value : [5500, 8900, 32]
size occupied in the stack : 8 bytes
value : [5500, 8900, 32]
size occupied in the stack : 24 bytes
$
```
