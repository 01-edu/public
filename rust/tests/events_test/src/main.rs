// You're have to design a notification system for a platform
// This events are Remainders, Registrations, Appointments, Holidays
// Create an event handler that depending of the type of event creates
// different notification: different color, different size and
// different position

// The possible positions are Top, Bottom and Center: Create and Enum
// `Position` with those values

// Create a struct called `Notification` with the fields
// size: u32,
// color: (u8, u8, u8),
// position: Position,
// content: String,

// The event that you have to handle are
// enum Event {
// 	Remainder(&str),
// 	Registration(Duration),
// 	Appointment(&str),
// 	Holiday,
// }

// Create a method called `notify`
// fn notify(&self) -> Notification
// That returns a notification with the following caracteristics for
// each
// Remainder:
// size= 50,
// color= (50, 50, 50),
// position= Bottom,
// content= the slice associated to the enum value

// Registration(chrono::Duration),
// size = 30,
// color = (255, 2, 22),
// position = Top,
// content = "You have `duration` left before the registration ends",
// `durations` must be displayed in the form of
// {hours}:{minutes}:{seconds} left for the beginning of the event
// for example if there is two hours 32 minutes and 3 seconds left
// before the registration then the content will be `You have 2:32:2 left before the registration ends`

//		Appointment(text)
// 	size: 100
// 	color: (200, 200, 3)
// 	position: Center
// 	content: text associated to the value

// Holiday
// size: 25
// color: (0, 255, 0)
// position: Top
// content: "Enjoy your holiday"

use chrono::Duration;
use events::Event::*;
#[allow(unused_imports)]
use events::{Notification, Position};

#[allow(dead_code)]
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

#[cfg(test)]
mod tests {

	use super::*;

	#[test]
	fn remainder_notification() {
		let remainder = Remainder("Go to the doctor");
		let notification = remainder.notify();
		println!("{}", &notification);
		assert_eq!(
			notification,
			Notification {
				size: 50,
				color: (50, 50, 50),
				position: Position::Bottom,
				content: "Go to the doctor".to_string(),
			}
		);
	}

	#[test]
	fn registration_notification() {
		let registration = Registration(Duration::seconds(49094));
		let notification = registration.notify();
		println!("{}", registration.notify());
		assert_eq!(
			notification,
			Notification {
				size: 30,
				color: (255, 2, 22),
				position: Position::Top,
				content: "You have 13H:38M:14S left before the registration ends".to_string(),
			}
		);
	}

	#[test]
	fn appointment_notification() {
		let appointment = Appointment("Go to the doctor");
		let notification = appointment.notify();
		println!("{}", &notification);
		assert_eq!(
			notification,
			Notification {
				size: 100,
				color: (200, 200, 3),
				position: Position::Center,
				content: "Go to the doctor".to_string(),
			}
		);
	}

	#[test]
	fn holiday_notification() {
		let holiday = Holiday;
		let notification = Holiday.notify();
		println!("{}", holiday.notify());
		assert_eq!(
			notification,
			Notification {
				size: 25,
				color: (0, 255, 0),
				position: Position::Top,
				content: String::from("Enjoy your holiday"),
			}
		);
	}
}
