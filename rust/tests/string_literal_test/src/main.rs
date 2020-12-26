/*
## string literal

### Instructions

Create the following functions:

- `is_empty`, that returns true if a string is empty
- `is_ascii`, that returns true if all characters of a given string is in ASCII range
- `contains`, that returns true if the string contains a pattern given
- `split_at`, that divides a string in two returning a tuple
- `find', that returns the index if the first character of a given string that matches the pattern

> This exercise will test the **heap allocation** of your function!
> So try your best to allocate the minimum data on the heap! (hit: &str)

### Notions

- https://doc.rust-lang.org/1.22.0/book/first-edition/the-stack-and-the-heap.html
- https://doc.rust-lang.org/rust-by-example/primitives/literals.html

*/

#[global_allocator]
static ALLOC: jemallocator::Jemalloc = jemallocator::Jemalloc;

#[allow(unused_imports)]
use jemalloc_ctl::{epoch, stats};
#[allow(unused_imports)]
use string_literal::*;

#[allow(dead_code)]
fn is_empty_sol(v: &str) -> bool {
	v.is_empty()
}

#[allow(dead_code)]
fn is_ascii_sol(v: &str) -> bool {
	v.is_ascii()
}

#[allow(dead_code)]
fn contains_sol(v: &str, pat: &str) -> bool {
	v.contains(pat)
}

#[allow(dead_code)]
fn split_at_sol(v: &str, index: usize) -> (&str, &str) {
	v.split_at(index)
}

#[allow(dead_code)]
fn find_sol(v: &str, pat: char) -> usize {
	v.find(pat).unwrap()
}

#[test]
fn test_memory() {
	// the statistics tracked by jemalloc are cached
	// The epoch controls when they are refreshed
	let e = epoch::mib().unwrap();
	// allocated: number of bytes allocated by the application
	let allocated = stats::allocated::mib().unwrap();

	// sense we only use string literals (&str)
	// the heap will not be allocated, because
	// &str are preallocated text (saved on the binary file)
	is_empty_sol("");
	is_ascii_sol("rust");
	contains_sol("rust", "ru");
	split_at_sol("rust", 2);
	find_sol("rust", 'u');
	// this will advance with the epoch giving the its old value
	// where we read the updated heap allocation using the `allocated.read()`
	e.advance().unwrap();
	let solution = allocated.read().unwrap();

	is_empty("");
	is_ascii_sol("rust");
	contains("rust", "er");
	split_at("rust", 2);
	find("rust", 'u');

	e.advance().unwrap();
	let student = allocated.read().unwrap();

	assert!(
		student <= solution,
		format!(
			"your heap allocation is {}, and it must be less or equal to {}",
			student, solution
		)
	);
}

#[test]
fn test_functions() {
	assert!(is_empty(""));
	assert!(!is_empty("something"));
	assert!(is_ascii("rust"));
	assert!(!is_ascii("rust¬"));
	assert!(contains("rust", "ru"));
	assert!(!contains("something", "mer"));
	assert_eq!(split_at("rust", 2), ("ru", "st"));
	assert_eq!(find("ru-st-e", '-'), 2);
	assert_eq!(find("ru-st-e", 'e'), 6);
}
