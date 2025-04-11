## middle_day

### Instructions

Use the [`chrono`](https://docs.rs/chrono/0.4.40/chrono) crate to create a **function** named `middle_day`. It accepts a year, and returns the weekday of the middle day of that year, wrapped in an `Option`.

Years with an even number of days do not have an exact single middle day, and thus should return `None`.

### Expected Function

```rs
pub fn middle_day(year: u32) -> Option<chrono::Weekday> {
    todo!()
}
```

### Usage

Here is a program to test your function:

```rs
use middle_day::*;

fn main() {
    println!("{:?}", middle_day(1022));
}
```

And its output:

```sh
$ cargo run
Some(Tue)
$
```

### Notions:

- [chrono-0.4.40](https://docs.rs/chrono/0.4.40/chrono)
