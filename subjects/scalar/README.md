## Scalar

### Instructions

Create the following **functions**, which each accept two parameters:

- `sum`, which returns the sum between two values from 0 to 255
- `diff`, which returns the difference between two values from -32768 to 32767
- `pro`, which returns the product of the multiplication between two values from -128 to 127
- `quo`, which returns the quotient of the division between two 32bit values
- `rem`, which returns the remainder of the division between two 32bit values

> You will need to figure out the exact data types for the parameters **and** the return values.
> There are some hints in the example.

```rust
pub fn sum(a: X, b: X) -> X {

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

### Example

```rust
use scalar::*;

fn main() {
    // sum
    println!("sum: {}", sum(234, 2)); // 'sum: 236'
    println!("sum: {}", sum(1, 255)); // 'ERROR: attempt to add with overflow'

    // diff
    println!("diff: {}", diff(234, 2)); // 'diff: 232'
    println!("diff: {}", diff(-32768, 32766)); // 'ERROR: attempt to subtract with overflow'

    // product
    println!("pro: {}", pro(23, 2)); // 'pro: 46'
    println!("pro: {}", pro(-128, 2)); // 'ERROR: attempt to multiply with overflow'

    // quotient
    println!("quo: {}", quo(22.0, 2.0)); // 'quo: 11'
    println!("quo: {}", quo(-128.23, 2.0)); // 'quo: -64.115'

    // remainder
    println!("rem: {}", rem(22.0, 2.0)); // 'rem: 0'
    println!("rem: {}", rem(-128.23, 2.0)); // 'rem: -0.22999573'
}
```

### Notions

- [Data Types](https://doc.rust-lang.org/book/ch03-02-data-types.html)
