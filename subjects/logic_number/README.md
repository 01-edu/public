## logic_number

### Instructions

In this exercise, the logic for a sequence of numbers will be tested.
You will have to create a function which will return `true` if the number is the sum of its own digits, where each digit is first raised to the power of the number of digits.

Examples:

- 9 returns true, because 9 = 9^1 = 9
- 10 returns false, because 10 != 1^2 + 0^2 = 1
- 153 returns true, because: 153 = 1^3 + 5^3 + 3^3 = 1 + 125 + 27 = 153
- 154 returns false, because: 154 != 1^3 + 5^3 + 4^3 = 1 + 125 + 64 = 190

### Expected function

```rust
pub fn number_logic(num: u32) -> bool {

}
```

### Usage

Here is a program to test your function.

```rust
use logic_number::*;

fn main() {
    let array = [9, 10, 153, 154];
    for pat in &array {
        if number_logic(*pat) == true {
            println!(
                "this number returns {} because the number {} obey the rules of the sequence",
                number_logic(*pat),
                pat
            )
        }
        if number_logic(*pat) == false {
            println!("this number returns {} because the number {} does not obey the rules of the sequence", number_logic(*pat),pat )
        }
    }
}
```

And its output:

```console
$ cargo run
this number returns true because the number 9 obey the rules of the sequence
this number returns false because the number 10 does not obey the rules of the sequence
this number returns true because the number 153 obey the rules of the sequence
this number returns false because the number 154 does not obey the rules of the sequence
$
```

### Notions

- [patterns](https://doc.rust-lang.org/book/ch18-00-patterns.html)
