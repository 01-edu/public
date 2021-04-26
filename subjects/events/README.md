## events

### Instructions

You have to design a notification system for a platform.

These events can be: Remainders, Registrations, Appointments or Holidays.

- Create an event handler that, depending on the type of event, creates different notifications with different colors, different sizes and different positions

- The possible positions are Top, Bottom and Center: Create and Enum `Position` with those values

- Create a method called `notify` which returns a notification with the following caracteristics for each

  - Remainder:
    size= 50,
    color= (50, 50, 50),
    position= Bottom,
    content= the slice associated to the enum value

  - Registration(chrono::Duration),
    size = 30,
    color = (255, 2, 22),
    position = Top,
    content = "You have `duration` left before the registration ends",

  `durations` must be displayed in the form of {hours}:{minutes}:{seconds} left for the beginning of the event for example if there is two hours 32 minutes and 3 seconds left before the registration then the content will be `You have 2:32:2 left before the registration ends`

  - Appointment(text)
    size: 100
    color: (200, 200, 3)
    position: Center
    content: text associated to the value

  - Holiday
    size: 25
    color: (0, 255, 0)
    position: Top
    content: "Enjoy your holiday"

- Implement the std::fmt::Display trait so the text of the notification is printed in the right color in the command line

### Notions

- [colored crate](https://docs.rs/colored/2.0.0/colored/)
- [chrono crate](https://crates.io/crates/chrono)

### Expected Functions and Data Structures

```rust
use chrono::Duration;
use colored::*;

#[derive(Debug, Eq, PartialEq)]
enum Position {
	Top,
	Bottom,
	Center,
}

#[derive(Debug, Eq, PartialEq)]
struct Notification {
	size: u32,
	color: (u8, u8, u8),
	position: Position,
	content: String,
}

#[derive(Debug)]
enum Event<'a> {
	Remainder(&'a str),
	Registration(Duration),
	Appointment(&'a str),
	Holiday,
}

use std::fmt;

impl fmt::Display for Notification {
}

use Event::*;

impl Event {
	fn notify(&self) -> Notification {
	}
}
```

### Usage

Here is a program to test your function.

```rust
use events::Event::*;
use chrono::Duration;

fn main() {
	let remainder = Remainder("Go to the doctor");
	println!("{}", remainder.notify());
	let registration = Registration(Duration::seconds(49094));
	println!("{}", registration.notify());
	let appointment = Appointment("Go to the doctor");
	println!("{}", appointment.notify());
	let holiday = Holiday;
	println!("{}", holiday.notify());
}
```

And its output

```console
student@ubuntu:~/events/test$ cargo run
(Bottom, 50, [38;2;50;50;50mGo to the doctor[0m)
(Top, 30, [38;2;255;2;22mYou have 13H:38M:14S left before the registration ends[0m)
(Center, 100, [38;2;200;200;3mGo to the doctor[0m)
(Top, 25, [38;2;0;255;0mEnjoy your holiday[0m)
student@ubuntu:~/events/test$
```
