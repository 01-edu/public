/*
## arrange_it

### Instructions

Create a function called `arrange_phrase` that takes a string literal as a phrase and returns it organized
Each word will have a number that indicates the position of that word

### Example

```rust
```

> This exercise will test the **heap allocation** of your function!
> So try your best to allocate the minimum data on the heap!

### Notions

- https://doc.rust-lang.org/1.22.0/book/first-edition/the-stack-and-the-heap.html
- https://doc.rust-lang.org/std/primitive.str.html#method.split

*/

#[allow(unused_imports)]
#[global_allocator]
static ALLOC: jemallocator::Jemalloc = jemallocator::Jemalloc;
#[allow(unused_imports)]
use arrange_it::*;

#[allow(unused_imports)]
use jemalloc_ctl::{epoch, stats};

#[allow(dead_code)]
fn main() {
	println!("{:?}", arrange_phrase("is2 Thi1s T4est 3a"));
}

// example of function that works but does not pass the heap test
// fn arrange_phrase(phrase: &str) -> String {
//     let words_nbr = phrase.matches(" ").count() + 1;
//     let mut result_vec:Vec<String> = vec!["".to_string();words_nbr];
//     for word in phrase.split_whitespace().into_iter() {
//         for i in 1..words_nbr+1 {
//             if word.contains(&i.to_string()){
//                 result_vec[i-1] = word.split(&i.to_string()).collect::<String>();
//             }
//         }
//     }
//     result_vec.join(" ")
// }

#[allow(dead_code)]
fn arrange_phrase_sol(phrase: &str) -> String {
	let nbrs: Vec<&str> = phrase.matches(char::is_numeric).collect();
	let a = &phrase.replace(char::is_numeric, "");
	let mut m: Vec<&str> = a.split_whitespace().collect();

	for (i, ele) in nbrs.iter().enumerate() {
		let strs: Vec<&str> = a.split_whitespace().collect();
		m[ele.parse::<usize>().unwrap() - 1] = strs[i];
	}
	m.join(" ")
}

#[test]
fn test_heap_memory_allocation() {
	// the statistics tracked by jemalloc are cached
	// The epoch controls when they are refreshed
	let e = epoch::mib().unwrap();
	// allocated: number of bytes allocated by the application
	let allocated = stats::allocated::mib().unwrap();
	let test_value = "4of Fo1r pe6ople g3ood th5e the2";

	arrange_phrase_sol(test_value);
	// this will advance with the epoch giving the its old value
	// where we read the updated heap allocation using the `allocated.read()`
	e.advance().unwrap();
	let solution = allocated.read().unwrap();

	arrange_phrase(test_value);
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
fn test_function() {
	let cases = vec![
		"4of Fo1r pe6ople g3ood th5e the2",
		"is2 Thi1s T4est 3a",
		"w7ork t3he a4rt o5f Per1formance is2 a6voiding",
	];
	for v in cases {
		assert_eq!(arrange_phrase(v), arrange_phrase_sol(v));
	}
}
