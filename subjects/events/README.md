## events

### Instructions

You have to design a notification system for a platform.

Depending on the type of event, your event handler will control the size, color and position of the notification.

Create a method named `notify` which returns a `Notification` with the following characteristics for each of:

- `Remainder(&str)`:
  - `size`: `50`
  - `color`: `(50, 50, 50)`
  - `position`: `Bottom`
  - `content`: the `&str` associated to the enum variant
- `Registration(std::time::Duration)`:
  - `size`: `30`
  - `color`: `(255, 2, 22)`
  - `position`: `Top`
  - `content`: `"You have {duration} left before the registration ends"`
- `Appointment(&str)`:
  - `size`: `100`
  - `color`: `(200, 200, 3)`
  - `position`: `Center`
  - `content`: the `&str` associated to the enum variant
- `Holiday`:
  - `size`: `25`
  - `color`: `(0, 255, 0)`
  - `position`: `Top`
  - `content`: `"Enjoy your holiday"`

`duration` must be displayed in the form of `{hours}H:{minutes}M:{seconds}S`. The time will represent the remaining time before the event starts. For example, if there are 13 hours, 38 minutes and 14 seconds left, then the content will be `"You have 13H:38M:14S left before the registration ends"`.

Implement the `std::fmt::Display` trait for `Notification` so that notifications are printed in the correct color and formatted according to the examples given on the Usage section.

### Dependencies

colored = "3.0.0"

### Expected Functions and Data Structures

```rust
use colored::*;
use std::{fmt, time::Duration};

#[derive(Debug, PartialEq, Clone, Copy)]
pub enum Position {
    Top,
    Bottom,
    Center,
}

#[derive(Debug, PartialEq, Clone)]
pub struct Notification {
    pub size: u32,
    pub color: (u8, u8, u8),
    pub position: Position,
    pub content: String,
}

#[derive(Clone, Copy)]
pub enum Event<'a> {
    Remainder(&'a str),
    Registration(Duration),
    Appointment(&'a str),
    Holiday,
}

impl fmt::Display for Notification {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        todo!()
    }
}

impl Event<'_> {
    pub fn notify(self) -> Notification {
        todo!()
    }
}
```

### Usage

Here is a program to test your function.

```rust
use events::*;
use std::time::Duration;

fn main() {
    println!("{}", Event::Remainder("Go to the doctor").notify());
    println!(
        "{}",
        Event::Registration(Duration::from_secs(49094)).notify()
    );
    println!("{}", Event::Appointment("Go to the doctor").notify());
    println!("{}", Event::Holiday.notify());
}
```

And its output

```console
$ cargo run
(Bottom, 50, [38;2;50;50;50mGo to the doctor[0m)
(Top, 30, [38;2;255;2;22mYou have 13H:38M:14S left before the registration ends[0m)
(Center, 100, [38;2;200;200;3mGo to the doctor[0m)
(Top, 25, [38;2;0;255;0mEnjoy your holiday[0m)
$
```

### Notions

- [colored crate](https://docs.rs/colored/3.0.0/colored/)
