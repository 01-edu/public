## Scalar

### Instructions
Create the following **functions**, which each receives two parameters:
- `sum`, which returns the sum between two values from 0 to 255
- `diff`, which returns the difference between two values from -32768 to 32767
- `pro`, which returns the product of the multiplication between two values from -128 to 127
- `quo`, which returns the quotient of the division between two values (32bit and you have to figure out the second part)
- `rem`, which returns the remainder of the division between two values (32bit and you have to figure out the second part)

You **must** complete the Expected functions parameters data type accordingly (Replace the Xs)! 

### Notions
- https://doc.rust-lang.org/book/ch03-02-data-types.html


### Expected functions (Incomplete, you must precise the Data Types)

```rust
pub fn sum(a: X , b: X) -> X {
	
}

pub fn diff(a: X, b: X) -> X {

}

pub fn pro(a: X, b: X) -> X {

}

pub fn quo(a: X, b: X) -> X {

}

pub fn rem(a: X, b: X) -> X {

}
```

### Usage (do not forget to comment the ERROR test if you want to use all the tests):

```rust
use scalar::sum;
use scalar::diff;
use scalar::pro;
use scalar::quo;
use scalar::rem;

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
