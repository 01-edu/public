## middle_day

### Instructions

Use the [`chrono` crate](https://docs.rs/chrono/0.4.19/chrono/index.html) to create a **function** named `middle_day`. It accepts a year, and returns the weekday of the middle day of that year, wrapped in an `Option`. `chrono::Weekday` has to be referred to as `wd`.

Years with an even number of days do not have a middle day, and should return `None`.

### Expected Function

> You'll need to work out the function signature for yourself.

### Usage

Here is a program to test your function:

```rs
use middle_day::*;

fn main() {
    let date = Utc.ymd(2011, 12, 2).and_hms(21, 12, 09);

    println!("{:?}", middle_day(1022).unwrap());
}
```

And its output:

```sh
$ cargo run
Tue
$
```
