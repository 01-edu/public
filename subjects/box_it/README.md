## box_it

### Instructions

Create the following functions:

- `transform_and_save_on_heap`, that receives a string with a number (that can or not have a k (kilo) suffix)
  and transforms those numbers into `u32` and inserts it into a vector that must be saved in the heap using **Box**.

- `take_value_ownership`, that takes the value (unboxed value) form the box and returns it

### Expected Function

```rust
pub fn transform_and_save_on_heap(s: String) -> Box<Vec<u32>> {}
pub fn take_value_ownership(a: Box<Vec<u32>>) -> Vec<u32> {}
```

### Usage

Here is a program to test your function

```rust
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
student@ubuntu:~/[[ROOT]]/test$ cargo run
Box value : [6800, 13500]
size occupied in the stack : 8 bytes
value : [6800, 13500]
size occupied in the stack : 24 bytes
student@ubuntu:~/[[ROOT]]/test$
```

### Notions

- https://doc.rust-lang.org/book/ch15-00-smart-pointers.html
- https://doc.rust-lang.org/book/ch15-01-box.html
