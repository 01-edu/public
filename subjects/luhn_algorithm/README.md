## luhn_algorithm

### Instructions

Create a function which checks if a number is valid per the Luhn formula.

The function will receive a string and return a boolean.

**Invalid inputs:** An empty string or a number with only one digit are considered invalid.
**Handling spaces:** Spaces are accepted but have to be ignored during the calculation (in other words they won't affect the result).

The Luhn formula is used to check if a number is a valid credit card number and in some other scenarios where you need to check a number fast and without accessing a database.

We can summarize the formula as follow:

- We want to check the number `4539 3195 0343 6467`
- We take every second digit starting by the right
- We multiply those digits by 2
- If the result is more than 9 we subtract 9 from it
- We sum all the digits
- If sum is evenly divisible by 10 then this number is valid

So we will get:

- `4539 3195 0343 6467`
- `4_3_ 3_9_ 0_4_ 6_6_`: numbers to modify
- `8_6_ 6_9_ 0_8_ 3_3_`: modified numbers (for numbers over 9 we subtracted 9 already)
- `8569 6195 0383 3437`: the new sequence of digits
- `80`: the sum of all digits
- `80` is evenly divisible by 10 so the result is `true`

### Expected Function

```rust
pub fn is_luhn_formula(code: &str) -> bool {
    todo!()
}
```

### Usage

Here is a possible program to test your function,

```rust
use luhn_algorithm::is_luhn_formula;

fn main() {
    println!("{}", is_luhn_formula(""));
    println!("{}", is_luhn_formula("1"));
    println!("{}", is_luhn_formula("79927398713"));
    println!("{}", is_luhn_formula("7992 7398 713"));
    println!("{}", is_luhn_formula("1234567890123456"));
}
```

And its output:

```console
$ cargo run
false
false
true
true
false
$
```
