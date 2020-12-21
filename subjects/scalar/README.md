## Scaler

### Instructions
Create the following functions, that receives two parameters:
- sum, that returns the sum between two values from 0 to 255
- diff, that returns the difference between two values from -32768 to 32767
- pro, that returns the product of the multiplication between two values from -128 to 127
- quo, that returns the quotient of the division between two values
- rem, that returns the remainder of the division between two values
### Notions
- https://doc.rust-lang.org/book/ch03-02-data-types.html


### Expected functions

```rust
pub fn sum(a:, b: ) ->  {
	
}

pub fn diff(a: , b: ) -> {

}

pub fn pro(a: , b: ) -> {

}

pub fn quo(a: , b: ) -> {

}

pub fn rem(a: , b: ) -> {

}
```

### Usage:

```rust
fn main() {
    // sum
    println!("sum : {}", sum(234, 2));
    println!("sum : {}", sum(1, 255)); // 'ERROR: attempt to add with overflow'
    // diff
    println!("diff : {}", diff(234, 2));
    println!("diff : {}", diff(-32768, 32766)); // 'ERROR: attempt to subtract with overflow'
    // product
    println!("pro : {}", pro(23, 2));
    println!("pro : {}", pro(-128, 2)); // 'ERROR: attempt to multiply with overflow'
    // quotient
    println!("quo : {}", quo(22.0, 2.0));
    println!("quo : {}", quo(-128.23, 2.0));
    // remainder
    println!("rem : {}", rem(22.0, 2.0));
    println!("rem : {}", rem(-128.23, 2.0));
}
```
