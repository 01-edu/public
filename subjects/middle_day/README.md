## middle_day

### Instructions

Use the [`chrono`](https://docs.rs/chrono/0.4.19/chrono/index.html) crate to create a function called `middle_day`, that returns the Weekday of the middle day of the year passed as an argument, wrapped in an Option.
You also should refer to chrono::Weekday as `wd`.

```rs
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
