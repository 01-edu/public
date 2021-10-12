## middle_day

### Instructions

Use the [`chrono crate`](https://docs.rs/chrono/0.4.19/chrono/index.html) to create a **function** called `middle_day`, which returns, wrapped in an Option, the Weekday of the middle day of the year passed as an argument.
`chrono::Weekday` has to be refered as `wd`.

As an even number of days doesn't have a middle day, when the year passed as an argument has an even number of days, the program should return `None`.

### Expected Function

#### For this exercise the signature of the function has to be found out.

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
