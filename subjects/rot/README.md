## rot

### Instructions

By now you will have the knowledge of the so called rotational cipher "ROT13".

A ROT13 on the Latin alphabet would be as follows:

- Plain: abcdefghijklmnopqrstuvwxyz

- Cipher: nopqrstuvwxyzabcdefghijklm

Your purpose in this exercise is to create a similar `rot` function that is a better version of the ROT13 cipher.

Your function will receive a string and a number and it will rotate each letter of that string the number of times settled by the second argument to the right, or to the left if the number are negative.

Your function should only change letters. If the string includes punctuation and numbers
they will remain the same.

### Notions

- https://doc.rust-lang.org/book/ch18-00-patterns.html

### Expected functions

```rust
pub fn rot(input: &str, key: i8) -> String {}
```

### Usage

Here is a program to test your function.

```rust
use rot::rot;

fn main() {

    println!("The letter \"a\" becomes: {}", rot("a", 26));
    println!("The letter \"m\" becomes: {}", rot("m", 0));
    println!("The letter \"m\" becomes: {}", rot("m", 13));
    println!("The letter \"a\" becomes: {}", rot("a", 15));
    println!("The word \"MISS\" becomes: {}", rot("MISS", 5));
    println!(
        "The decoded message is: {}",
        rot("Gur svir obkvat jvmneqf whzc dhvpxyl.", 13)
    );
    println!(
        "The decoded message is: {}",
        rot("Mtb vznhpqd ifky ozrunsl ejgwfx ajc", 5)
    );
    println!(
        "Your cypher wil be: {}",
        rot("Testing with numbers 1 2 3", 4)
    );
    println!("Your cypher wil be: {}", rot("Testing", -14));
    println!("The letter \"a\" becomes: {}", rot("a", -1));
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
The letter "a" becomes: a
The letter "m" becomes: m
The letter "m" becomes: z
The letter "a" becomes: p
The word "MISS" becomes: RNXX
The decoded message is: The five boxing wizards jump quickly.
The decoded message is: Ryg aesmuvi nkpd tewzsxq jolbkc foh
Your cypher wil be: Xiwxmrk amxl ryqfivw 1 2 3
Your cypher wil be: Fqefuzs
The letter "a" becomes: z

student@ubuntu:~/[[ROOT]]/test$
```
